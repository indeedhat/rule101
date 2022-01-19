package main

import (
	"bytes"
	"fmt"
)

const GridSize = 100

type GridRow []uint8

func (gr GridRow) String() string {
	var buf bytes.Buffer

	for _, cell := range gr {
		if cell == 1 {
			buf.WriteByte('*')
		} else {
			buf.WriteByte(' ')
		}
	}

	return buf.String()
}

func main() {
	prevRow := make(GridRow, GridSize)
	prevRow[GridSize-2] = 1
	fmt.Println(prevRow)

	for row := 1; row < GridSize; row++ {
		currentRow := GridRow{}

		for col := 0; col < GridSize; col++ {
			mask := prevRow[max(0, col-1)] << 2
			mask |= prevRow[col] << 1
			mask |= prevRow[min(GridSize-1, col+1)] << 0

			if mask == 7 || mask == 4 || mask == 0 {
				currentRow = append(currentRow, 0)
			} else {
				currentRow = append(currentRow, 1)
			}
		}
		fmt.Println(currentRow)
		prevRow = currentRow
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
