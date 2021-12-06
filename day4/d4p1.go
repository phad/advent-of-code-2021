package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type board struct {
	cells [][]int
	marks [][]bool
}

func newBoard(lines []string) (*board, error) {
	b := &board{}
	for _, line := range lines {
		log.Printf("next line: %q", line)
		var rowVals []int
		var rowMarks []bool
		for _, n := range strings.Split(line, " ") {
			if n == "" {
				continue
			}
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			rowVals = append(rowVals, num)
			rowMarks = append(rowMarks, false)
		}
		b.cells = append(b.cells, rowVals)
		b.marks = append(b.marks, rowMarks)
	}
	return b, nil
}

func (b *board) mark(drawn int) {
	for r := 0; r < len(b.cells); r++ {
		for c := 0; c < len(b.cells[r]); c++ {
			if b.cells[r][c] == drawn {
				b.marks[r][c] = true
				return
			}
		}
	}
}

func (b board) isRowFullyMarked(r int) bool {
	m := true
	for c := range b.marks[r] {
		m = m && b.marks[r][c]
	}
	return m
}

func (b board) isColFullyMarked(c int) bool {
	m := true
	for r := range b.marks {
		m = m && b.marks[r][c]
	}
	return m
}

func (b board) isBingo() bool {
	// assumes square boards.
	for i := range b.marks {
		if b.isRowFullyMarked(i) || b.isColFullyMarked(i) {
			return true
		}
	}
	return false
}

func (b board) unmarkedSum() int {
	s := 0
	for row := range b.marks {
		for col, mark := range b.marks[row] {
			if !mark {
				s += b.cells[row][col]
			}
		}
	}
	return s
}

func parseNumbersDrawn(nums string) ([]int, error) {
	var drawn []int
	for _, n := range strings.Split(nums, ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		drawn = append(drawn, num)
	}
	return drawn, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <in file>")
	}
	lines, err := readLines(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Read %d input lines", len(lines))

	idx, remain := 0, len(lines)
	numbersDrawn, err := parseNumbersDrawn(lines[idx])
	if err != nil {
		log.Fatalf("parseNumbersDrawn: err=%v", err)
	}
	log.Printf("Numbers drawn: %#v", numbersDrawn)

	idx += 2  // header line, blank line
	remain -= 2

	var boards []*board
	for {
		if remain < 5 {
			break
		}
		b, err := newBoard(lines[idx:idx+5])
		if err != nil {
			log.Fatalf("newBoard: err=%v", err)
		}
		boards = append(boards, b)
		idx += 6  // 5 for board, 1 empty
		remain -= 6
	}

	log.Printf("Read %v boards", len(boards))

	for _, drawn := range numbersDrawn {
		// mark all boards
		for idx, board := range boards {
			board.mark(drawn)
			if board.isBingo() {
				unmarkedSum := board.unmarkedSum()
				fmt.Printf("Board %d, Drawn %d: BINGO! Unmarked sum %d, Product %d\n", idx, drawn, unmarkedSum, drawn*unmarkedSum)
			}
		}
	}
}
