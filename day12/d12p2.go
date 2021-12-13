package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type edge struct {
	id       int
	from, to string
}

func (e *edge) String() string {
	return fmt.Sprintf("edge %d: %s-%s", e.id, e.from, e.to)
}

type vertex struct {
	id      string
	isSmall bool
	edges   []int
}

func (v *vertex) String() string {
	return fmt.Sprintf("vert %s (small: %t): edges %v", v.id, v.isSmall, v.edges)
}

type graph struct {
	edges      []*edge
	vertices   map[string]*vertex
	start, end *vertex
}

func (g *graph) String() string {
	var b strings.Builder
	b.WriteString("edges:\n")
	for _, e := range g.edges {
		b.WriteString("  "+e.String()+"\n")
	}
	b.WriteString("vertices:\n")
	for _, v := range g.vertices {
		b.WriteString("  "+v.String()+"\n")
	}
	return b.String()
}

type route []*vertex

type path struct {
	route        route
}

func (p path) canVisitPart1(vID string) bool {
	for _, v := range p.route {
		if vID == v.id {
			return false
		}
	}
	return true
}

func (p path) canVisitPart2(vID string) bool {
	if vID == "start" {
		return false
	}
        smallCounts := map[string]int{}
	haveSmall2Visits := false
        for _, v := range p.route {
		if v.isSmall {
			smallCounts[v.id]++
		}
		if smallCounts[v.id] == 2 {
			haveSmall2Visits = true
		}
	}
	for _, v := range p.route {
                if vID != v.id {
			continue
		}
		if v.isSmall && haveSmall2Visits {
                        return false
                }
        }
        return true
}

func (r route) String() string {
	var ids []string
	for _, v := range r {
		ids = append(ids, v.id)
	}
	return strings.Join(ids, ",")
}

func readGraph(input []string) *graph {
	g := &graph{vertices: make(map[string]*vertex)}
	for idx, in := range input {
		vertIDs := strings.Split(in, "-")
		if len(vertIDs) != 2 {
			log.Fatalf("malformed input: %q", in)
		}
		fwdEdge  := &edge{id: 2*idx, from: vertIDs[0], to: vertIDs[1]}
		backEdge := &edge{id: 2*idx+1, from: vertIDs[1], to: vertIDs[0]}
		for _, vertexID := range []string{vertIDs[0], vertIDs[1]} {
			v, ok := g.vertices[vertexID]
			if !ok {
				small := vertexID==strings.ToLower(vertexID)
				v = &vertex{id: vertexID, isSmall: small, edges: []int{}}
				g.vertices[vertexID] = v
				if vertexID == "start" { g.start = v }
				if vertexID == "end"   { g.end   = v }
			}
			if vertexID == vertIDs[0] {
				v.edges = append(v.edges, fwdEdge.id)
			} else {
				v.edges = append(v.edges, backEdge.id)
			}
		}
		g.edges = append(g.edges, fwdEdge)
		g.edges = append(g.edges, backEdge)
	}
	return g
}

func exploreVertex(g *graph, p *path, maxDepth int, recordFn func(r route)) {
	if len(p.route) == 0 || maxDepth == 0 {
		return
	}

//	fmt.Printf("exploreVertex: currentPath = %v\n", p.route)
	currVert := p.route[len(p.route)-1]
//	fmt.Printf("  currVert = %v\n", currVert)

	for _, eID := range currVert.edges {
		edge := g.edges[eID]
//		fmt.Printf("    next edge %v\n", edge)
		nextVert := g.vertices[edge.to]

		// Don't explore small nodes we've already been to.
		if nextVert.isSmall && !p.canVisitPart2(edge.to) {
//			fmt.Printf("      already visited small vert %v, skip\n", edge.to)
			continue
		}
		if nextVert == currVert {
//			fmt.Printf("      edge takes us back to self %v, skip\n", edge.to)
			continue
		}

		// Push nextVert onto our path before exploring.
		p.route = append(p.route, nextVert)

		// If we're at the end, record it
		if nextVert == g.end {
//			fmt.Printf("      reached end, recording route %v\n", p.route)
			recordFn(p.route)
		} else {
			// recurse into nextVert.
			exploreVertex(g, p, maxDepth-1, recordFn)
		}

		// done exploring through nextVert so pop from the path.
		p.route = p.route[0:len(p.route)-1]
	}
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

	caves := readGraph(input)
	fmt.Printf("Cave System:\n%v\n", caves)

	var routes []route

	p := &path{ route: []*vertex{caves.start} }
	exploreVertex(caves, p, 20, func(r route) {
		var rClone route
		for _, v := range r {
			rClone = append(rClone, v)
		}
		routes = append(routes, rClone)
	})

	fmt.Printf("Found %d routes\n", len(routes))
	for i, r := range routes {
		fmt.Printf("Route %d: %v\n", i, r)
	}
}

