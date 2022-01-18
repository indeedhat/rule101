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
	grid := []GridRow{}

	for row := 0; row < GridSize; row++ {
		grid = append(grid, GridRow{})

		for col := 0; col < GridSize; col++ {
			if row == 0 && col == GridSize-2 {
				grid[row] = append(grid[row], 1)
				continue
			}

			if row == 0 {
				grid[row] = append(grid[row], 0)
				continue
			}

			mask := grid[row-1][max(0, col-1)] << 2
			mask |= grid[row-1][col] << 1
			mask |= grid[row-1][min(GridSize-1, col+1)] << 0

			if mask == 7 || mask == 4 || mask == 0 {
				grid[row] = append(grid[row], 0)
			} else {
				grid[row] = append(grid[row], 1)
			}
		}
		fmt.Println(grid[row])
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
