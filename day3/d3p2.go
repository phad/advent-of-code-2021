package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <in file>")
	}
	lines, err := readLines(os.Args[1])
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Read %d input lines", len(lines))

	type counts struct{ n0, n1 int }
	var bitCounts []counts

	for _, bin := range lines {
		if bin == "" {
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

	oxyCandidates, co2Candidates := map[int]struct{}{}, map[int]struct{}{}
	for idx := range lines {
		oxyCandidates[idx] = struct{}{}
		co2Candidates[idx] = struct{}{}
	}
outer:
	for pos, counts := range bitCounts {
		for idx, bin := range lines {
			_, isOxyCandidate := oxyCandidates[idx]
			_, isCo2Candidate := co2Candidates[idx]
			if !isOxyCandidate && !isCo2Candidate {
				continue
			}
			bitAtPos := bin[pos]
			more0, more1, equal := counts.n0>counts.n1, counts.n1>counts.n0, counts.n0==counts.n1
			matchesMostCommonOrEqually1 :=
				(bitAtPos == '0' && more0) || (bitAtPos == '1' && !more0)
			matchesLeastCommonOrEqually0 :=
				(bitAtPos == '0' && !more0) || (bitAtPos == '1' && more0)

			//fmt.Printf("considering line %02d %q pos %d: bit=%q, counts.n0=%d,.n1=%d\n", idx, bin, pos, bitAtPos, counts.n0, counts.n1)
			if isOxyCandidate && !matchesMostCommonOrEqually1 && (len(oxyCandidates) > 1) {
				fmt.Printf("   more0=%t more1=%t equal=%t bit=%q - removing Oxy cand %q\n", more0, more1, equal, bitAtPos, bin)
				delete(oxyCandidates, idx)
			}
			if isCo2Candidate && !matchesLeastCommonOrEqually0 && (len(co2Candidates) > 1) {
				fmt.Printf("   more0=%t more1=%t equal=%t bit=%q - removing Co2 cand %q\n", more0, more1, equal, bitAtPos, bin)
				delete(co2Candidates, idx)
			}
			if len(oxyCandidates) == 1 && len(co2Candidates) == 1 {
				fmt.Println("\nAborting search, we're done to one for each")
				break outer
			}
		}
		fmt.Printf("\n\nat end of considering bit position pos, remaining oxy cands are:\n")
		for idx := range oxyCandidates {
			fmt.Printf(" %s", lines[idx])
		}
		fmt.Printf("\n\nat end of considering bit position pos, remaining co2 cands are:\n")
		for idx := range co2Candidates {
                        fmt.Printf(" %s", lines[idx])
                }
		fmt.Println("\n")
	}

	for idx := range oxyCandidates {
		fmt.Printf("oxygen generator rating = %v (idx %d)\n", lines[idx], idx)
	}
	for idx := range co2Candidates {
		fmt.Printf("co2 scrubber rating = %v (idx %d)\n", lines[idx], idx)
	}
}
