package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type scrambling struct {
	// from the puzzle input
	digits []string
	output []string
	//
	candidates      map[string][]int  // TODO: remove this
	candidateDigits map[int][]string
	//
	// maps scrambled segment letter to standard
	mapping    map[byte]byte
	// maps standard segment letter to scrambled
	revMapping map[byte]byte
	//
	correctDigits   map[int]string
}

func (s *scrambling) String() string {
	return fmt.Sprintf("digits: %v output:%v", strings.Join(s.digits, ","), strings.Join(s.output, ","))
}

func (s *scrambling) init() {
	// abcdefg refer to 7 segments.  We know that digits
	// 1: lights two segments
	// 7: lights three segments
	// 4: lights four segments
	// 2, 3, 5: light five segments
	// 0, 6, 9: light six segments
	// 8: lights seven segments
	for _, p := range s.digits {
//		log.Printf("Considering: %s of length %d\n", p, len(p))
		switch len(p) {
		case 0, 1:
			log.Fatalf("Impossible 7-segment combination! Num segs=%d", len(p))
		case 2:
			// 1: lights two segments
			s.candidates[p] = append(s.candidates[p], 1)
			s.candidateDigits[1] = append(s.candidateDigits[1], p)
			s.correctDigits[1] = p
		case 3:
			// 7: lights three segments
			s.candidates[p] = append(s.candidates[p], 7)
			s.candidateDigits[7] = append(s.candidateDigits[7], p)
			s.correctDigits[7] = p
		case 4:
			// 4: lights four segments
			s.candidates[p] = append(s.candidates[p], 4)
			s.candidateDigits[4] = append(s.candidateDigits[4], p)
			s.correctDigits[4] = p
		case 5:
			// 2, 3, 5: light five segments
			s.candidates[p] = append(s.candidates[p], []int{2, 3, 5}...)
			s.candidateDigits[2] = append(s.candidateDigits[2], p)
			s.candidateDigits[3] = append(s.candidateDigits[3], p)
			s.candidateDigits[5] = append(s.candidateDigits[5], p)
		case 6:
			// 0, 6, 9: light six segments
			s.candidates[p] = append(s.candidates[p], []int{0, 6, 9}...)
			s.candidateDigits[0] = append(s.candidateDigits[0], p)
			s.candidateDigits[6] = append(s.candidateDigits[6], p)
			s.candidateDigits[9] = append(s.candidateDigits[9], p)
		case 7:
			// 8: lights seven segments
			s.candidates[p] = append(s.candidates[p], 8)
			s.candidateDigits[8] = append(s.candidateDigits[8], p)
			s.correctDigits[8] = p
		default:
			log.Printf("len segs = %d", len(p))
		}
	}
//	fmt.Printf("candidateDigits: %v\n", s.candidateDigits)
}

func (s *scrambling) Deduce() {
	// Here's the standard segment lettering:
	//    aaa
	//   b   c
	//   b   c
	//    ddd
	//   e   f
	//   e   f
	//    ggg

	// Compare 7 (3 segments) with 1 (2 segments):
	//  Foreach scrambled segment 's':
	//   if s in 7 but not in 1 -> s=aaa
	seven, one := s.correctDigits[7], s.correctDigits[1]
	diffs := charsInFirstNotInSecond(seven, one)
	assert(len(diffs) == 1)
	s.mapping[diffs[0]] = 'a'
	s.revMapping['a'] = diffs[0]
//	fmt.Printf("Segment mapping: %v\nReverse mapping: %v\n", s.mapping, s.revMapping)

	// We also know by analysing the standard digits
	// that whichever segment 's' is only used:
	//   9 times in the 10 digits -> s=fff
	//   6 times in the 10 digits -> s=bbb
	//   4 times in the 10 digits -> s=eee
	segmentCounts := map[byte]int{}
	for _, d := range s.digits {
		for i := 0; i < len(d); i++ {
			segmentCounts[d[i]]++
		}
	}
//	fmt.Printf("segmentCounts=%v\n", segmentCounts)
	for seg, count := range segmentCounts {
		switch count {
		case 4:
			if _, ok := s.mapping[seg]; ok { assert(false) }
			s.mapping[seg] = 'e'
			s.revMapping['e'] = seg
		case 6:
			if _, ok := s.mapping[seg]; ok { assert(false) }
			s.mapping[seg] = 'b'
			s.revMapping['b'] = seg
		case 9:
			if _, ok := s.mapping[seg]; ok { assert(false) }
			s.mapping[seg] = 'f'
			s.revMapping['f'] = seg
		default:
		}
	}
//	fmt.Printf("Segment mapping: %v\nReverse mapping: %v\n", s.mapping, s.revMapping)

	// Now we have f, we can work out c, as only 1 has those two.
	if s.correctDigits[1][0] == s.revMapping['f'] {
		seg := s.correctDigits[1][1]
		s.mapping[seg] = 'c'
		s.revMapping['c'] = seg
	} else if s.correctDigits[1][1] == s.revMapping['f'] {
		seg := s.correctDigits[1][0]
		s.mapping[seg] = 'c'
		s.revMapping['c'] = seg
	}
//	fmt.Printf("Segment mapping: %v\nReverse mapping: %v\n", s.mapping, s.revMapping)

	// Once we have 5 segments a, b, c, e, f the remaining two can be
	// worked out as follows:
	// Digit 0 has 'g' but not 'd', so we're looking for one of the 6-segment
	// candidates where we lack 2 mappings (the other two 6 and 9 both have 'd'
	// so for those we only lack 1 mapping).
	for _, sixCand := range s.candidateDigits[6] {
		diffs := charsInFirstNotInSecond(s.correctDigits[8], sixCand)
		assert(len(diffs) == 1)
		seg := diffs[0]
		if _, ok := s.mapping[seg]; ok {
			continue
		}
		s.mapping[seg] = 'd'
		s.revMapping['d'] = seg
	}
//	fmt.Printf("Segment mapping: %v\nReverse mapping: %v\n", s.mapping, s.revMapping)

	// The final unmapped segment must be 'g'.
	var b strings.Builder
	for ch := range s.mapping {
		b.WriteByte(ch)
	}
	diffs = charsInFirstNotInSecond("abcdefg", b.String())
	assert(len(diffs) == 1)
	seg := diffs[0]
	s.mapping[seg] = 'g'
	s.revMapping['g'] = seg
//	fmt.Printf("Segment mapping: %v\nReverse mapping: %v\n", s.mapping, s.revMapping)
}

func assert(cond bool) {
	if !cond { log.Fatalf("boom") }
}

func charsInFirstNotInSecond(first, second string) []byte {
	m := map[byte]bool{}
	for i := 0; i < len(first); i++ {
		m[first[i]] = true
	}
	for j := 0; j < len(second); j++ {
		delete(m, second[j])
	}
	var diff []byte
	for ch := range m {
		diff = append(diff, ch)
	}
	return diff
}

func (s *scrambling) Candidates(digit string) []int {
	return s.candidates[digit]
}

func (s *scrambling) digitAt(pos int) (int, error) {
	assert(len(s.mapping) > 0)

	digitsBySegs := map[string]int{
		"abcefg":  0,
		"cf":      1,
		"acdeg":   2,
		"acdfg":   3,
		"bcdf":    4,
		"abdfg":   5,
		"abdefg":  6,
		"acf":     7,
		"abcdefg": 8,
		"abcdfg":  9,
	}

	outDigit := s.output[pos]
	var segs []string
	for i := 0; i < len(outDigit); i++ {
		segs = append(segs, string(s.mapping[outDigit[i]]))
	}
	sort.Strings(segs)
	digit, ok := digitsBySegs[strings.Join(segs, "")]
	if !ok {
		return 0, fmt.Errorf("digitsBySegs lookup failed for seg %q", strings.Join(segs, ""))
	}
	return digit, nil
}

func (s *scrambling) OutputValue() (int, error) {
	val := 0
	for i := 0; i < len(s.output); i++ {
		digit, err := s.digitAt(i)
		if err != nil {
			return 0, err
		}
		val += digit
		if i < len(s.output) - 1 {
			val *= 10
		}
	}
	return val, nil
}


func newScrambling(digits, output []string) (*scrambling, error) {
	if len(digits) != 10 {
		return nil, fmt.Errorf("wrong number of digits: %q", digits)
	}
	if len(output) != 4 {
		return nil, fmt.Errorf("wrong number of output: %q", output)
	}
	scr := &scrambling{
		digits:          digits,
		output:          output,
		candidates:      make(map[string][]int),
		candidateDigits: make(map[int][]string),
		correctDigits:   make(map[int]string),
		mapping:         make(map[byte]byte),
		revMapping:      make(map[byte]byte),
	}
	scr.init()
	return scr, nil
}

func parseInput(input []string) []*scrambling {
	// 10* scrambled segment states | 4* output digits
	// scrambling is consistent for the whole line, but
	// differs from one line to another.
	var scramblings []*scrambling
	for _, in := range input {
		parts := strings.Split(in, " | ")
		if len(parts) == 0 {
			continue
		}
		if len(parts) != 2 {
			log.Fatalf("Malformed input: %q", in)
		}
		digits := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")
		scr, err := newScrambling(digits, output)
		if err != nil {
			log.Fatalf("Parse error: %v", err)
		}
//		fmt.Printf("Read scrambling: %v\n", scr)
		scramblings = append(scramblings, scr)
	}
	return scramblings
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

	scramblings := parseInput(input)


	// Part 1: how many times do digits 1, 4, 7, or 8 appear.
	count1478 := 0
	for _, scr := range scramblings {
		for _, outDigit := range scr.output {
			cands := scr.Candidates(outDigit)
//			fmt.Printf("Candidates for %q are %v\n", outDigit, cands)
			if len(cands) == 1 {
				count1478++
			}
		}
	}
	fmt.Printf("Part 1: Digits 1, 4, 7 or 8 appear %d times in input.\n", count1478)

	sumOutputs := 0
	for _, scr := range scramblings {
		scr.Deduce()
		outVal, err := scr.OutputValue()
		if err != nil {
			log.Fatalf("OutputValue err: %v", err)
		}
		sumOutputs += outVal
	}
	fmt.Printf("Part 2: Sum of all the outputs is %d\n", sumOutputs)
}

