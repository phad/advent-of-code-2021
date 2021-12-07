package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type fish struct {
	spawnRate int
	cycleCount int
	daysToNextSpawn int
}

func (f fish) String() string {
	return fmt.Sprintf("rate: %d days cycles-since-birth: %d days days-to-next-spawn: %d", f.spawnRate, f.cycleCount, f.daysToNextSpawn)
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

	if len(input) > 1 {
		log.Fatalf("too much input: got %d lines, want 1", len(input))
	}

	var shoal []*fish
	fmt.Println("input = "+input[0]);
	valStrs := strings.Split(input[0], ",")
	fmt.Sprintf("%v", valStrs)
	for _, vStr := range valStrs {
		val, err := strconv.Atoi(vStr)
		if err != nil {
			log.Fatalf("parseInput err: %v", err)
		}
		shoal = append(shoal, &fish{
			spawnRate: val,
			cycleCount: 0,
			daysToNextSpawn: val,
		})
	}

	for cycle := 0; cycle < 80; cycle++ {
		fmt.Printf("---> Start cycle %d Shoal size: %d --->\n", cycle, len(shoal))
		sort.Sort(byIncreasingDaysToSpawn(shoal))
		var roes []*fish
		for i, f := range shoal {
			_ = i
//			fmt.Printf("fish #%d: %v\n", i, f.String())
			if f.daysToNextSpawn == 0 {
				parentRate := f.spawnRate
				roe := &fish{
					spawnRate: parentRate,
					cycleCount: 0,
					daysToNextSpawn: 8,  // 2+parentRate,
				}
//				fmt.Printf("new fish! %v\n", roe.String())
				f.daysToNextSpawn = 7
				roes = append(roes, roe)
			}

			f.cycleCount++
			f.daysToNextSpawn--
		}
		for _, r := range roes {
			shoal = append(shoal, r)
		}
//		intervals := []string{}
//		for _, f := range shoal {
//			intervals = append(intervals, fmt.Sprintf("%d", f.daysToNextSpawn))
//		}
//		fmt.Printf("After %2d days: %v\n", cycle+1, strings.Join(intervals, ","))
	}
	fmt.Printf("Final number of lanternfish: %d\n", len(shoal))
}

type byIncreasingDaysToSpawn []*fish
func (d byIncreasingDaysToSpawn) Len() int           { return len(d) }
func (d byIncreasingDaysToSpawn) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d byIncreasingDaysToSpawn) Less(i, j int) bool { return d[i].daysToNextSpawn < d[j].daysToNextSpawn }

