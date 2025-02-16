package _11_4

import (
	"fmt"
)

func passByValue() {
	s := make([]int, 1)
	m := make(map[int]int)
	u := User{}
	fmt.Println(s, m, u)
	change(s, m, u)
	fmt.Println(s, m, u)
}

func change(s []int, m map[int]int, u User) {
	s[0] = 100
	m[0] = 100
	u.Name = "name"
}
