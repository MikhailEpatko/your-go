package _8_4

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	STR_COUNT = 3
)

func CountVowelsInArray(strs [STR_COUNT]string) {
	fmt.Println(countVowelsInArrayStr(&strs))
}

func countVowelsInArrayStr(strs *[STR_COUNT]string) string {
	var counters [STR_COUNT]int
	for i, s := range strs {
		for _, c := range s {
			switch c {
			case 'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я',
				'А', 'Е', 'Ё', 'И', 'О', 'У', 'Ы', 'Э', 'Ю', 'Я':
				counters[i]++
			}
		}
	}
	var stringCounters [STR_COUNT]string
	for i, counter := range counters {
		stringCounters[i] = strconv.Itoa(counter)
	}
	return strings.Join(stringCounters[:], " ")
}
