package main

import (
	"bytes"
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

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

// TODO: the two halves of this func are nearly identical
// and could be collapsed.
func (b *board) markLine(l *line) {
	xRng := l.to.x - l.from.x
	yRng := l.to.y - l.from.y
	if abs(xRng) > abs(yRng) {
		// Iterate in x
//		fmt.Println("it in x")
		xInc, yInc := 1, 0
		if xRng < 0 { xInc = -1 }
		if yRng > 0 {
			yInc = 1
		} else if yRng < 0 {
			yInc = -1
		}
		x, y := l.from.x, l.from.y
		for {
			if x == l.to.x+xInc {
				break
			}
//			fmt.Printf("x/marking cell (%d,%d)\n", x, y)
			b.cells[y][x]++
			x += xInc
			y += yInc
		}
	} else {
		// Iterate in y
//		fmt.Println("it in y")
		xInc, yInc := 0, 1
		if yRng < 0 { yInc = -1 }
		if xRng > 0 {
			xInc = 1
		} else if xRng < 0 {
			xInc = -1
		}
		x, y := l.from.x, l.from.y
		for {
			if y == l.to.y+yInc {
				break
			}
//			fmt.Printf("y/marking cell (%d,%d)\n", x, y)
			b.cells[y][x]++
                        x += xInc
			y += yInc
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
