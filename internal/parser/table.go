package parser

import (
	"fmt"
	"math"
	"strings"
)

type truthTable [][]string

func (t truthTable) Print() {
	if len(t) == 0 {
		return
	}
	numCols := len(t[0])
	columnWidths := make([]int, numCols)
	for _, row := range t {
		for i, val := range row {
			columnWidths[i] = max(columnWidths[i], len(val))
		}
	}
	t.printBorder(columnWidths)
	for i, row := range t {
		for i, val := range row {
			fmt.Printf("| %-*s ", columnWidths[i], val)
		}
		fmt.Println("|")
		if i == 0 {
			t.printBorder(columnWidths)
		}
	}
	t.printBorder(columnWidths)
}

func (t truthTable) printBorder(columnWidths []int) {
	for _, width := range columnWidths {
		fmt.Printf("+%-*s", width+2, strings.Repeat("-", width))
	}
	fmt.Println("+")
}

func emptyTable(n int) ([][]bool, int) {
	rowCount := int(math.Pow(2, float64(n)))
	table := make([][]bool, rowCount)
	for i := 0; i < rowCount; i++ {
		row := make([]bool, n+1)
		for j := 0; j < n; j++ {
			alternationHeight := int(math.Pow(2, float64(n-j-1)))
			if (i/alternationHeight)%2 == 0 {
				row[j] = true
			} else {
				row[j] = false
			}
		}
		table[i] = row
	}
	return table, rowCount
}
