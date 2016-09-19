package scheduler

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	sched := new(GreedyScheduler)

	sched.Add(&Job{})
	sched.Add(&Job{})
	sched.Add(&Job{})

	result := sched.queue.Len()
	if result != 3 {
		t.Errorf("Got %d, expected 3", result)
	}
}

func TestSchedule(t *testing.T) {
	sched := new(GreedyScheduler)

	jobs := []*Job{
		&Job{Weight: 5, Length: 1},
		&Job{Weight: 2, Length: 1},
		&Job{Weight: 3, Length: 1},
		&Job{Weight: 4, Length: 1},
		&Job{Weight: 1, Length: 1},
	}

	for _, j := range jobs {
		sched.Add(j)
	}

	result := sched.Schedule()
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
