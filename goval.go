package goval

import (
	"fmt"
	"reflect"
)

// goVal GoVal类型
type goVal struct {
	// ShowTypePrefix 是否展示类型前缀
	ShowTypePrefix bool

	// Value GoVal底层的值类型
	Value reflect.Value
}

func newGoVal(v reflect.Value, typePrefix bool) *goVal {
	return &goVal{
		ShowTypePrefix: typePrefix,
		Value:          v,
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
// 	如果v为结构体，for循环每个结构体的filed，迭代结果值
// 	如果v为普通类型，这直接打印
//	如果v为指针或者接口，获取其类型值，继续迭代
//	如果v为Slice，for循环迭代每个元素的值
func (gv *goVal) String() string {
	v := gv.Value
	typePrefix := gv.ShowTypePrefix
	switch v.Kind() {
	case reflect.Invalid:
		return "<nil>"
	case reflect.Interface, reflect.Ptr:
		t := v.Type()
		if v.IsZero() {
			return fmt.Sprintf("%s<nil>", gv.getTypeName(t))
		}
		gv := newGoVal(v.Elem(), typePrefix)
		return fmt.Sprintf("&%s", gv.String())
	case reflect.Struct:
		t := v.Type()
		out := gv.getTypeName(t) + "{"
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				out += ", "
			}
			gv := newGoVal(v.Field(i), typePrefix)
			out += fmt.Sprintf(`%s:%s`, t.Field(i).Name, gv.String())
		}
		out += "}"
		return out
	case reflect.Slice:
		out := gv.getTypeName(v.Type())
		if v.IsZero() {
			out += "<nil>"
			return out
		}

		out += "["
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				out += ", "
			}
			gv := newGoVal(v.Index(i), typePrefix)
			out += gv.String()
		}
		out += "]"
		return out
	case reflect.Map:
		out := gv.getTypeName(v.Type())
		if v.IsZero() {
			out += "<nil>"
			return out
		}

		out += "["
		for i, key := range v.MapKeys() {
			if i > 0 {
				out += " "
			}
			gv := newGoVal(v.MapIndex(key), typePrefix)
			out += fmt.Sprintf(`{"%v":%v}`, key, gv.String())
		}
		out += "]"
		return out
	default:
		return fmt.Sprintf("%#v", v)
	}
}

// getTypeName 获取类型名称
func (gv *goVal) getTypeName(t reflect.Type) string {
	if !gv.ShowTypePrefix {
		return ""
	}
	if t.PkgPath() == "main" {
		return t.Name()
	}
	return t.String()
}
