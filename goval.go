package goval

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// goVal GoVal类型
type goVal struct {
	// Value GoVal底层的值类型
	Value reflect.Value

	// Show 是否展示类型前缀
	Show bool
}

func newGoVal(v reflect.Value, show bool) *goVal {
	return &goVal{
		Value: v,
		Show:  show,
	}
}

// ToString 将任意go类型变量，转为字符串
func ToString(v any) string {
	return newGoVal(reflect.ValueOf(v), false).String()
}

// ToTypeString 将任意go类型变量，转为含类型前缀的字符串
func ToTypeString(v any) string {
	return newGoVal(reflect.ValueOf(v), true).String()
}

// 断言v的值，不断迭代返回结果
// 	如果v为普通类型，这直接打印
//	如果v为指针或者接口，获取其类型值，继续迭代
// 	如果v为结构体，for循环每个结构体的filed，迭代结果值
//	如果v为Slice，for循环迭代每个元素的值
func (gv *goVal) String() string {
	buff := &strings.Builder{}
	writeValBuff(gv.Value, buff, gv.Show)
	return buff.String()
}

// 将v的内容写入buff
func writeValBuff(v reflect.Value, buff *strings.Builder, show bool) {
	switch v.Kind() {
	case reflect.Invalid:
		buff.WriteString("<nil>")
	case reflect.Interface, reflect.Ptr:
		t := v.Type()
		if v.IsZero() {
			buff.WriteString(fmt.Sprintf("%s<nil>", getTypeName(t, show)))
		}

		// &elem
		buff.WriteString("&")
		writeValBuff(v.Elem(), buff, show)
	case reflect.Struct:
		t := v.Type()
		buff.WriteString(getTypeName(t, show))

		// {key1:val1, key2:val2, ...}
		buff.WriteString("{")
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buff.WriteString(", ")
			}
			buff.WriteString(t.Field(i).Name)
			buff.WriteString(":")
			writeValBuff(v.Field(i), buff, show)
		}
		buff.WriteString("}")
	case reflect.Slice:
		buff.WriteString(getTypeName(v.Type(), show))
		if v.IsZero() {
			buff.WriteString("<nil>")
			return
		}

		// [item1, item2, ...]
		buff.WriteString("[")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buff.WriteString(", ")
			}
			writeValBuff(v.Index(i), buff, show)
		}
		buff.WriteString("]")
	case reflect.Map:
		buff.WriteString(getTypeName(v.Type(), show))
		if v.IsZero() {
			buff.WriteString("<nil>")
			return
		}

		// [{"key1":"val1", "key2":"val2"} ...]
		buff.WriteString("[")

		// sort map keys
		keys := v.MapKeys()
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})

		// write buff
		for i, key := range keys {
			if i > 0 {
				buff.WriteString(" ")
			}
			buff.WriteString("{")
			buff.WriteString(fmt.Sprintf(`"%v":`, key))
			writeValBuff(v.MapIndex(key), buff, show)
			buff.WriteString("}")
		}

		buff.WriteString("]")
	default:
		buff.WriteString(fmt.Sprintf("%#v", v))
	}
}

type sortMapKeys struct {
	keys []reflect.Value
}

// getTypeName 获取类型名称
func getTypeName(t reflect.Type, showTypePrefix bool) string {
	if !showTypePrefix {
		return ""
	}
	if t.PkgPath() == "main" {
		return t.Name()
	}
	return t.String()
}
