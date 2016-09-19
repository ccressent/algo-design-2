package scheduler

type Job struct {
	Weight float64
	Length float64
}

func (j Job) Score() float64 {
	return WLRatio(j.Weight, j.Length)
}

func WLDifference(w, l float64) float64 {
	return (w - l)
}

func WLRatio(w, l float64) float64 {
	return (w / l)
}
