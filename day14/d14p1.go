package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <in file>")
	}
	input, err := readLines(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Read %d input lines", len(input))

	template := input[0]

	insertions := map[string]string{}
	for i := 2; i < len(input); i++ {
		bits :=	strings.Split(input[i], " -> ")
		insertions[bits[0]] = bits[1]
	}

	fmt.Println(insertions)
	fmt.Println("Template:     "+template)

	type op struct {
		pos int
		orig string
		sub bool
		result string
	}
	var operations []op

	for j := 0; j < 10; j++ {
		for i := 0; i < len(template)-1; i++ {
			pair := template[i:i+2]
			insert, ok := insertions[pair]
			operations = append(operations, op{
				pos: i,
				orig: pair,
				sub: ok,
				result: pair[0:1]+insert+pair[1:2],
			})
		}
		//fmt.Printf("Operations: %v\n", operations)
		var nextTmpl strings.Builder
		for i, op := range operations {
			count := len(op.result)
			if i < len(operations)-1 {
				count--
			}
			nextTmpl.WriteString(op.result[:count])
		}
		operations = []op{}
		template = nextTmpl.String()
		//fmt.Printf("After step %d: %v\n", j+1, template)
	}

	counts := map[string]int {}
	for i := 0; i < len(template); i++ {
		counts[template[i:i+1]]++
	}
	//fmt.Println(counts)
	min, max := int(99999999999), int(0)
	minCh, maxCh := "", ""
	for ch, cnt := range counts {
		if cnt > max { max = cnt ; maxCh = ch }
		if cnt < min { min = cnt ; minCh = ch }
	}
	fmt.Printf("Max count %s(%d), max count %s(%d), diff %d\n", maxCh, max, minCh, min, max-min)
}

