package scheduler

type PriorityQueue []*Job

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	scoreI := pq[i].Score()
	scoreJ := pq[j].Score()

	// If 2 jobs have the same score, pick the one with the biggest weight
	if scoreI == scoreJ {
		return pq[i].Weight > pq[j].Weight
	}

	return scoreI > scoreJ
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(j interface{}) {
	job := j.(*Job)
	*pq = append(*pq, job)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	item := old[n-1]
	*pq = old[:n-1]

	return item
}
