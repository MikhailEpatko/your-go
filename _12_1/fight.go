package _12_1

import "fmt"

type Character interface {
	Attack() string
}

func Fight(characters []Character) {
	for _, ch := range characters {
		fmt.Println(ch.Attack())
	}
}
