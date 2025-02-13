package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	random := 82.0

	fmt.Printf("Исходное число: %f\n", random)
	fmt.Printf("Исходное число, увеличенное на 10%%: %.5f\n", random*1.1)
	fmt.Println("Исходное число является четным:", math.Mod(random, 2) == 0)
	fmt.Println("Предпоследняя цифра целой части исходного числа:", (int64(random)/10)%10)

	var val any
	val = 5
	msg := ""
	switch val.(type) {
	case int:
		msg = "В переменной val находится тип int."
	case string:
		msg = "В переменной val находится тип string."
	case float64:
		msg = "В переменной val находится тип float64."
	case bool:
		msg = "В переменной val находится тип bool."
	default:
		msg = "В переменной val находится неизвестный тип данных."
	}
	fmt.Println(msg)

}

func GenerateCompliment(name string) string {
	v := rand.Intn(3) + 1
	msg := ""
	switch v {
	case 1:
		msg = fmt.Sprintf("Ты великолепен, %s!", name)
	case 2:
		msg = fmt.Sprintf("У тебя потрясающая улыбка, %s!", name)
	case 3:
		msg = fmt.Sprintf("Ты вдохновляешь, %s!", name)
	}
	return msg
}
