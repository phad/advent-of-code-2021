package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dot struct { x, y int }
type fold struct { x, y int }

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <in file>")
	}
	input, err := readLines(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Read %d input lines", len(input))

	dots := map[dot]bool{}
	var folds []fold

	// Parse dot coords.
	idx := 0
	for ; idx < len(input); idx++ {
		if len(input[idx]) == 0 { break }
		valStrs := strings.Split(input[idx], ",")
		x, err := strconv.Atoi(valStrs[0])
		if err != nil { log.Fatalf("Atoi(%s): err=%v", valStrs[0], err) }
		y, err := strconv.Atoi(valStrs[1])
		if err != nil { log.Fatalf("Atoi(%s): err=%v", valStrs[1], err) }

//		fmt.Printf("Dot: %d,%d\n", x, y)
		dots[dot{x, y}] = true
	}
	idx++ // skip to start of next section

	// Parse fold instructions.
	for ; idx < len(input); idx++ {
		bits := strings.Split(input[idx], "=")
		val, err := strconv.Atoi(bits[1])
		if err != nil { log.Fatalf("Atoi(%s): err=%v", bits[0], err) }

		var f fold
		len := len(bits[0])
		axis := bits[0][len-1:len]
		if axis == "x" { f.x = val }
		if axis == "y" { f.y = val }
		folds = append(folds, f)
//		fmt.Printf("Fold: %v\n", f)
	}

	// Execute fold instructions.
	for i, f := range folds {
		var remove, add []dot
		for d := range dots {
			if f.x > 0 {
				// Folding in x
				if d.x > f.x {
					remove = append(remove, d)
					add = append(add, dot{x: 2*f.x - d.x, y: d.y})
				}
			} else if f.y > 0 {
				// Folding in y
				if d.y > f.y {
					remove = append(remove, d)
					add = append(add, dot{x: d.x, y: 2*f.y - d.y})
				}
			}
		}
		for _, d := range remove { delete(dots, d) }
		for _, d := range add    { dots[d] = true  }
		fmt.Printf("After fold %d we have %d dots.\n", i+1, len(dots))
	}
	// Output result
//	fmt.Printf("Dots: %v\n", dots)
	width, height := 0, 0
	for d := range dots {
		if d.x > width  { width  = d.x }
		if d.y > height { height = d.y }
	}
	banner := [][]byte{}
	for y := 0; y <= height; y++ {
		var line []byte
		for x := 0; x <= width; x++ {
			line = append(line, byte('.'))
		}
		banner = append(banner, line)
	}
	for d := range dots {
		banner[d.y][d.x]=byte('#')
	}
	for _, b := range banner {
		fmt.Println(string(b))
	}
}


