package _7_4

import (
	"fmt"
	"strings"
)

const (
	TEMPLATE         = "%d x %d = %d"
	COLUMN_DELIMETER = "\t"
	ROW_DELIMETER    = "\n"
)

func printTable(num int) {
	fmt.Print(tableStr(num))
}

func tableStr(num int) string {
	b := strings.Builder{}
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			_, _ = b.WriteString(fmt.Sprintf(TEMPLATE, i, j, i*j))
			if j != num {
				_, _ = b.WriteString(COLUMN_DELIMETER)
			}
		}
		_, _ = b.WriteString(ROW_DELIMETER)
	}
	return b.String()
}
