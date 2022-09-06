package main

import (
	"fmt"

	"github.com/hold7techs/goval"
)

type Favor struct {
	ID    int
	Title string
}

type Address map[string]string

type User struct {
	id    int
	name  string
	Addrs Address
	Favs  []*Favor
}

func main() {
	userA := &User{
		id:   100,
		name: "Terry",
		Addrs: Address{
			"home": "Red Flower 101",
			"com":  "NanShan Street 10000",
		},
		Favs: []*Favor{
			{1, "badminton"},
			{2, "swimming"},
		},
	}
	fmt.Printf("user info[default %%+v] => %+v\n", userA)
	fmt.Printf("user info[goval.ToString(v)] => %s\n", goval.ToString(userA))
	fmt.Printf("user info[goval.ToTypeString(v)] => %s\n\n", goval.ToTypeString(userA))

	// slice
	userB := &User{
		id:    200,
		name:  "Clark",
		Addrs: nil,
		Favs:  nil,
	}
	userList := []*User{userA, userB}
	fmt.Printf("user list[default %%+v] => %+v\n", userList)
	fmt.Printf("user list[goval.ToString(v)] => %s\n", goval.ToString(userList))
	fmt.Printf("user list[goval.ToTypeString(v)] => %s\n\n", goval.ToTypeString(userList))

	// map
	userMap := map[int]*User{
		100: userA,
		200: userB,
	}
	fmt.Printf("user map[default %%+v] => %+v\n", userMap)
	fmt.Printf("user map[goval.ToString(v)] => %s\n", goval.ToString(userMap))
	fmt.Printf("user map[goval.ToTypeString(v)] => %s\n\n", goval.ToTypeString(userMap))

	// chan print pointer address
	ch := make(chan int)
	fmt.Printf("chan => %s\n", goval.ToString(ch))
	fmt.Printf("chan => %s\n", goval.ToTypeString(ch))
}
