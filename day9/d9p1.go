package main

import (
	"fmt"
	"log"
	"os"
)

type heightMap struct {
	numRows, numCols int
	heights [][]int
}

func (h *heightMap) isLowPoint(row, col int) bool {
	if row < 0 || row > h.numRows {
		log.Fatalf("row out of range: got %d want 0<=row<%d", row, h.numRows)
	}
	if col < 0 || col > h.numCols {
		log.Fatalf("col out of range: got %d want 0<=col<%d", col, h.numCols)
	}
	height := h.heights[row][col]
	// fmt.Printf("Cell %d, %d has height %d\n", row, col, height)
	if row > 0 && h.heights[row-1][col] <= height {
		// fmt.Println("row above check failed")
		return false
	}
	if row < h.numRows-1 && h.heights[row+1][col] <= height {
		// fmt.Println("row below check failed")
		return false
	}
	if col > 0 && h.heights[row][col-1] <= height {
		// fmt.Println("col left check failed")
		return false
	}
	if col < h.numCols-1 && h.heights[row][col+1] <= height {
		// fmt.Println("col right check failed")
		return false
	}
	// fmt.Printf("Cell %d, %d is a low point.\n", row, col)
	return true
}

func (h *heightMap) riskLevel() int {
	risk := 0
	for r, row := range h.heights {
		for c, height := range row {
			if !h.isLowPoint(r, c) {
				continue
			}
			risk += height + 1
		}
	}
	return risk
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <in file>")
	}
	input, err := readLines(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Read %d input lines", len(input))

	numRows := len(input)
	numCols := len(input[0])

	hMap := &heightMap{
		numRows: numRows,
		numCols: numCols,
	}
	for r := 0; r < numRows; r++ {
		//fmt.Printf("input row %d = %q\n", r, input[r]);
		row := []int{}
		if len(input[r]) != numCols {
			log.Fatalf("Input row %d is wrong length: got %d want %d", r, len(input[r]), numCols)
		}
		for c := 0; c < numCols; c++ {
			row = append(row, int(input[r][c] - '0'))
		}
		hMap.heights = append(hMap.heights, row)
	}
	fmt.Printf("Risk level = %d\n", hMap.riskLevel())
}

