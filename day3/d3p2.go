package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type candidates struct {
	kind string
	isCandidate map[int]struct{}
	rating uint64
}

type keepFn func(bitAtPos byte, more0, more1 bool) bool

type counts struct{ n0, n1 int }

func newCandidates(kind string, size int) *candidates {
	c := &candidates{
		kind: kind,
		isCandidate: make(map[int]struct{}, size),
	}
	for idx := 0; idx < size; idx++ {
		c.isCandidate[idx] = struct{}{}
	}
	return c
}

func (c *candidates) countBits(lines []string) []counts {
	var bitCounts []counts
	for idx, bin := range lines {
		_, isCand := c.isCandidate[idx]
		if bin == "" || !isCand {
			continue
                }

		sz := len(bin)
		for len(bitCounts) < sz {
			bitCounts = append(bitCounts, counts{})
                }
                for pos := sz-1; pos >= 0; pos-- {
                        if bin[pos] == '0' {
                                bitCounts[pos].n0++;
                        } else if bin[pos] == '1' {
                                bitCounts[pos].n1++;
                        }
                }
	}
	return bitCounts
}

func (c *candidates) filter(lines []string, pos int, k keepFn) bool {
	bitCounts := c.countBits(lines)

	for idx, bin := range lines {
		_, isCand := c.isCandidate[idx]
                if !isCand {
                        continue
                }

                bitAtPos := bin[pos]
                cnt := bitCounts[pos]
		more0, more1 := cnt.n0>cnt.n1, cnt.n1>cnt.n0

		log.Printf(" --- line %02d %q pos %d: bit=%q, counts.n0=%d,.n1=%d\n", idx, bin, pos, bitAtPos, cnt.n0, cnt.n1)

		if (!k(bitAtPos, more0, more1)) {
			equal := cnt.n0==cnt.n1
			log.Printf("  +- more0=%t more1=%t equal=%t bit=%q - removing %s cand %q\n", more0, more1, equal, bitAtPos, c.kind, bin)
                        delete(c.isCandidate, idx)
		}

		if len(c.isCandidate) == 1 {
			ratingBin := ""
			for idx := range c.isCandidate {
				ratingBin = lines[idx]
			}
			rating, err := strconv.ParseUint(ratingBin, 2, 32)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%s rating: %s (%d)\n", c.kind, ratingBin, rating)
			c.rating = rating
			return false
		}
        }
	return true
}

func (c *candidates) dump(pos int, lines []string) string {
	var b strings.Builder
	fmt.Fprintf(&b, "After considering bit position %d, remaining %s cands are:\n", pos, c.kind)
	for idx := range c.isCandidate {
		fmt.Fprintf(&b, " %s", lines[idx])
	}
	b.WriteByte('\n')
	return b.String()
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

	oxyCands := newCandidates("oxy", len(lines))
	co2Cands := newCandidates("co2", len(lines))

	pos := 0
	for {
		more := oxyCands.filter(lines, pos, func(bitAtPos byte, more0, more1 bool) bool {
			return (bitAtPos == '0' && more0) || (bitAtPos == '1' && !more0)
		})
		log.Println(oxyCands.dump(pos, lines))
		if !more {
			break
		}
		pos++
	}

	pos = 0
	for {
		more := co2Cands.filter(lines, pos, func(bitAtPos byte, more0, more1 bool) bool {
                        return (bitAtPos == '0' && !more1) || (bitAtPos == '1' && more0)
                })
		log.Println(co2Cands.dump(pos, lines))
		if !more {
			break
		}
		pos++
	}
	fmt.Printf("Rating product = %v\n", oxyCands.rating*co2Cands.rating)
}

