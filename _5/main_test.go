package main

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GenerateCompliment("Павел"))
	}
}
