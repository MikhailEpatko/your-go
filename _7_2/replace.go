package _7_2

import "fmt"

func PrintReplaced(s string) {
	fmt.Println(replacedStr(s))
}

func replacedStr(s string) string {
	res := []rune{}
	for _, r := range s {
		if r == 'у' {
			res = append(res, 'а')
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}
