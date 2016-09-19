package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./scheduler"
)

func main() {
	sched := new(scheduler.GreedyScheduler)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")

		// Skip lines that don't have a [weight] [length] format
		if len(tokens) != 2 {
			continue
		}

		w, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: ", err)
			continue
		}

		l, err := strconv.ParseFloat(tokens[1], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error: ", err)
			continue
		}

		sched.Add(&scheduler.Job{w, l})
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error: ", err)
		return
	}

	// Compute the schedule
	s := sched.Schedule()

	// Compute the sum of the weighted completion times of the schedule
	r := 0.0
	c := 0.0
	for _, j := range s {
		c += j.Length
		r += j.Weight * c
	}

	fmt.Println(r)
}
