package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./graph"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Skip the first line
	scanner.Scan()

	G := graph.NewUndirectedGraph()

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")

		// Skip any line that doesn't have the [vertex] [vertex] [cost] format
		if len(tokens) != 3 {
			continue
		}

		v1Name := tokens[0]
		v2Name := tokens[1]

		cost, err := strconv.ParseInt(tokens[2], 10, 0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: ", err)
			fmt.Fprintln(os.Stderr, "Skipping line")
			continue
		}

		if G.Get(v1Name) == nil {
			G.Add(&graph.Vertex{Name: v1Name})
		}

		if G.Get(v2Name) == nil {
			G.Add(&graph.Vertex{Name: v2Name})
		}

		v1 := G.Get(v1Name)
		v2 := G.Get(v2Name)

		G.Connect(v1, v2, int(cost))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		os.Exit(1)
	}

	mstCost := 0
	mstEdges := G.MST_Prim()

	for _, e := range mstEdges {
		mstCost += e.Cost
	}

	fmt.Println("MST cost: ", mstCost)
}
