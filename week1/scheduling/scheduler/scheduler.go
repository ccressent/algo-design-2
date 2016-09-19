package scheduler

import "container/heap"

type GreedyScheduler struct {
	queue PriorityQueue
}

func (s *GreedyScheduler) Add(j *Job) {
	heap.Push(&s.queue, j)
}

func (s *GreedyScheduler) Schedule() []*Job {
	pqSize := s.queue.Len()
	schedule := make([]*Job, 0, s.queue.Len())

	for i := 0; i < pqSize; i++ {
		schedule = append(schedule, heap.Pop(&s.queue).(*Job))
	}

	return schedule
}
