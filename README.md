# Package

The `goval` package supports converting complex type variables into human-readable strings, mainly for formatting log
recording.
(Including pointer type structure nesting, pointer slices, pointer Maps, etc.)

# Usage
`go get -u github.com/hold7techs/goval`

# Feature
1. **Easy Using**: `goval.ToString(v)`, `goval.ToTypeString(v)`
2. **Sort Map Key**

# Example

```
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
    // struct
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
```

**Result Print**

```
user info[default %+v] => &{id:100 name:Terry Addrs:map[com:NanShan Street 10000 home:Red Flower 101] Favs:[0x1400009e000 0x1400009e018]}
user info[goval.ToString(v)] => &{id:100, name:"Terry", Addrs:[{"home":"Red Flower 101"} {"com":"NanShan Street 10000"}], Favs:[&{ID:1, Title:"badminton"}, &{ID:2, Title:"swimming"}]}
user info[goval.ToTypeString(v)] => &User{id:100, name:"Terry", Addrs:Address[{"home":"Red Flower 101"} {"com":"NanShan Street 10000"}], Favs:[]*main.Favor[&Favor{ID:1, Title:"badminton"}, &Favor{ID:2, Title:"swimming"}]}

user list[default %+v] => [0x1400009a000 0x1400009a080]
user list[goval.ToString(v)] => [&{id:100, name:"Terry", Addrs:[{"home":"Red Flower 101"} {"com":"NanShan Street 10000"}], Favs:[&{ID:1, Title:"badminton"}, &{ID:2, Title:"swimming"}]}, &{id:200, name:"Clark", Addrs:<nil>, Favs:<nil>}]
user list[goval.ToTypeString(v)] => []*main.User[&User{id:100, name:"Terry", Addrs:Address[{"home":"Red Flower 101"} {"com":"NanShan Street 10000"}], Favs:[]*main.Favor[&Favor{ID:1, Title:"badminton"}, &Favor{ID:2, Title:"swimming"}]}, &User{id:200, name:"Clark", Addrs:Address<nil>, Favs:[]*main.Favor<nil>}]

user map[default %+v] => map[100:0x1400009a000 200:0x1400009a080]
user map[goval.ToString(v)] => [{"100":&{id:100, name:"Terry", Addrs:[{"home":"Red Flower 101"} {"com":"NanShan Street 10000"}], Favs:[&{ID:1, Title:"badminton"}, &{ID:2, Title:"swimming"}]}} {"200":&{id:200, name:"Clark", Addrs:<nil>, Favs:<nil>}}]
user map[goval.ToTypeString(v)] => map[int]*main.User[{"200":&User{id:200, name:"Clark", Addrs:Address<nil>, Favs:[]*main.Favor<nil>}} {"100":&User{id:100, name:"Terry", Addrs:Address[{"home":"Red Flower 101"} {"com":"NanShan Street 10000"}], Favs:[]*main.Favor[&Favor{ID:1, Title:"badminton"}, &Favor{ID:2, Title:"swimming"}]}}]

chan => (chan int)(0x140000aa3c0)
chan => (chan int)(0x140000aa3c0)
```

# Benchmark

The performance is not very good, but it can be used normally, and performance optimization may be done in the future!

## v0.2.0
```
$ go test -benchmem -bench .
goos: darwin
goarch: arm64
pkg: github.com/hold7techs/goval
BenchmarkGetGoString-10           620848              1942 ns/op             912 B/op         44 allocs/op
PASS
ok      github.com/hold7techs/goval     2.294s
```

## v0.1.0
```
$ go test -benchmem -bench .
goos: darwin
goarch: arm64
pkg: github.com/hold7techs/goval
BenchmarkGetGoString-10           304468              3438 ns/op            2465 B/op         96 allocs/op
PASS
ok      github.com/hold7techs/goval     1.167s
```

