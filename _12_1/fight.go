package _12_1

import "fmt"

func Fight(characters []Character) {
	for _, ch := range characters {
		fmt.Println(ch.Attack())
	}
}
