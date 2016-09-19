package scheduler

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	var pq PriorityQueue
	heap.Init(&pq)

	jobs := []*Job{
		&Job{Weight: 5, Length: 1},
		&Job{Weight: 2, Length: 1},
		&Job{Weight: 3, Length: 1},
		&Job{Weight: 4, Length: 1},
		&Job{Weight: 1, Length: 1},
	}

	for _, j := range jobs {
		heap.Push(&pq, j)
	}

	pqSize := pq.Len()
	result := make([]*Job, 0, len(jobs))
	for i := 0; i < pqSize; i++ {
		result = append(result, heap.Pop(&pq).(*Job))
	}

	expected := []*Job{
		jobs[0],
		jobs[3],
		jobs[2],
		jobs[1],
		jobs[4],
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Got %v, expected %v", result, expected)
	}
}
