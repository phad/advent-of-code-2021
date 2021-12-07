package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type board struct {
	cells [][]int
	marks [][]bool
}

type point struct { x, y int }
type line struct { from, to point }

func makePoint(input string) (point, error) {
	vals := strings.Split(input, ",")
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		return point{}, err
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		return point{}, err
	}
	return point{x: x, y: y}, nil
}

func parseInput(input []string) ([]*line, error) {
	var lines []*line
	for _, in := range input {
//		log.Printf("next line: %q", in)
		coords := strings.Split(in, " -> ")
		from, err := makePoint(coords[0])
		if err != nil {
			return nil, err
		}
		to, err := makePoint(coords[1])
                if err != nil {
                        return nil, err
                }
		lines = append(lines, &line{from: from, to: to})
	}
	return lines, nil
}

func boardDimensions(lines []*line) (int, int) {
	var maxX, maxY int
	for _, l := range lines {
		if l.from.x > maxX { maxX = l.from.x } 
		if l.to.x > maxX { maxX = l.to.x }
		if l.from.y > maxY { maxY = l.from.y }
		if l.to.y > maxY { maxY = l.to.y }
	}
	return maxX+1, maxY+1
}

func newBoard(lines []*line) *board {
	w, h := boardDimensions(lines)
	board := &board{}
	for r := 0; r < h; r++ {
		board.cells = append(board.cells,  make([]int, w))
	}
	return board
}

func (b *board) markLines(lines []*line) {
	for _, l := range lines {
//		fmt.Printf("Marking line #%d : (%d,%d) -> (%d,%d)\n", idx, l.from.x, l.from.y, l.to.x, l.to.y)
		b.markLine(l)
//		fmt.Println(b.String())
	}
}

func (b *board) markLine(l *line) {
	if l.from.x == l.to.x {
		yRng := []int{l.from.y, l.to.y}
		sort.Ints(yRng)
		for y := yRng[0]; y <= yRng[1]; y++ {
//			fmt.Printf("y/marking cell (%d,%d)\n", y, l.from.x)
			b.cells[y][l.from.x]++
		}
	} else if l.from.y == l.to.y {
		xRng := []int{l.from.x, l.to.x}
		sort.Ints(xRng)
		for x := xRng[0]; x <= xRng[1]; x++ {
//			fmt.Printf("x/marking cell (%d,%d)\n", x, l.from.y)
			b.cells[l.from.y][x]++
		}
	}
}

func (b board) String() string {
	var buf bytes.Buffer
        for row := range b.cells {
                buf.WriteString("[")
                for col := range b.cells[row] {
                        buf.WriteString(fmt.Sprintf("%01d ", b.cells[row][col]))
                }
                buf.WriteString("]\n")
        }
	return buf.String()
}

func (b board) moreThanOneSum() int {
	s := 0
	for row := range b.cells {
		for _, count := range b.cells[row] {
			if count > 1 {
				s++
			}
		}
	}
	return s
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

	lines, err := parseInput(input)
	if err != nil {
		log.Fatalf("parseInput err: %v", err)
	}

	board := newBoard(lines)
	board.markLines(lines)

	fmt.Printf("Number of cells with >1 overlapping lines: %d\n", board.moreThanOneSum())
}
