package goval

import (
	"testing"
)

func TestToGoTypeString(t *testing.T) {
	type user struct {
		id   int
		Name string
	}
	userA := user{id: 1, Name: "userA"}
	userB := user{id: 1, Name: "userB"}

	tests := []struct {
		name string
		v    interface{}
		want string
	}{
		{"nil", nil, "<nil>"},
		{"int", 1, "1"},
		{"bool-false", false, "false"},
		{"bool-true", true, "true"},
		{"struct-normal", userA, `goval.user{id:1, Name:"userA"}`},
		{"struct-pointer", &userA, `&goval.user{id:1, Name:"userA"}`},
		{"slice-nil", nil, `<nil>`},
		{"slice-normal", []user{userA, userB},
			`[]goval.user[goval.user{id:1, Name:"userA"}, goval.user{id:1, Name:"userB"}]`},
		{"slice-pointer", []*user{&userA, &userB},
			`[]*goval.user[&goval.user{id:1, Name:"userA"}, &goval.user{id:1, Name:"userB"}]`},
		{"map-nil", nil, `<nil>`},
		{"map-normal", map[string]interface{}{"user1": userA, "user2": userB},
			`map[string]interface {}[{"user1":&goval.user{id:1, Name:"userA"}} {"user2":&goval.user{id:1, Name:"userB"}}]`},
		{"map-pointer", map[string]*user{"user1": &userA},
			`map[string]*goval.user[{"user1":&goval.user{id:1, Name:"userA"}}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToTypeString(tt.v)
			if got != tt.want {
				t.Errorf("ToTypeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type user struct {
		id   int
		Name string
	}
	userA := user{id: 1, Name: "userA"}
	userB := user{id: 1, Name: "userB"}

	tests := []struct {
		name string
		v    interface{}
		want string
	}{
		{"nil", nil, "<nil>"},
		{"int", 1, "1"},
		{"bool-false", false, "false"},
		{"bool-true", true, "true"},
		{"struct-normal", userA, `{id:1, Name:"userA"}`},
		{"struct-pointer", &userA, `&{id:1, Name:"userA"}`},
		{"slice-nil", nil, `<nil>`},
		{"slice-normal", []user{userA, userB}, `[{id:1, Name:"userA"}, {id:1, Name:"userB"}]`},
		{"slice-pointer", []*user{&userA, &userB}, `[&{id:1, Name:"userA"}, &{id:1, Name:"userB"}]`},
		{"map-nil", nil, `<nil>`},
		{"map-normal", map[string]interface{}{"user1": userA, "user2": userB},
			`[{"user1":&{id:1, Name:"userA"}} {"user2":&{id:1, Name:"userB"}}]`},
		{"map-pointer", map[string]*user{"user1": &userA}, `[{"user1":&{id:1, Name:"userA"}}]`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToString(tt.v)
			if got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintNestUserInfo(t *testing.T) {
	type ExtInfo struct {
		From      string
		Intro     string
		SalePrice map[string]float64
	}

	// 结构体指针嵌套
	type Goods struct {
		id       int
		Price    float64
		Name     string
		LabelIDs []int
		ExtInfo  *ExtInfo
	}

	apple := &Goods{
		id:       100,
		Price:    10.25,
		Name:     "apple",
		LabelIDs: []int{4, 5},
		ExtInfo: &ExtInfo{
			From:  "Hk",
			Intro: "red apple...",
			SalePrice: map[string]float64{
				"high": 15.00,
				"low":  3.00,
			},
		},
	}
	banner := &Goods{
		id:       200,
		Price:    6.18,
		Name:     "banner",
		LabelIDs: []int{4, 5},
		ExtInfo:  nil,
	}
	goods := []*Goods{apple, banner}

	// 不带类型
	t.Logf("user1 => %s", ToString(goods))

	// 带Go Type类型
	t.Logf("user1 type val => %s", ToTypeString(goods))
}

// old -> BenchmarkGetGoString-10           350077              3407 ns/op            2473 B/op         97 allocs/op
func BenchmarkGetGoString(b *testing.B) {
	type Extend struct {
		School string
		Home   string
	}

	// 结构体指针嵌套
	type NestUser struct {
		ID     int
		Name   string
		Favs   []int
		Extend *Extend
		Family map[string]string
	}

	u1 := &NestUser{
		ID:   100,
		Name: "UserA",
		Favs: []int{4, 5},
		Extend: &Extend{
			School: "SchoolXiao",
			Home:   "HomeTang",
		},
		Family: map[string]string{
			"mother": "mo",
			"father": "fa",
		},
	}
	for i := 0; i < b.N; i++ {
		ToString(u1)
	}
}
