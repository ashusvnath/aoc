package main

import "fmt"

type Recorder interface {
	Record(t CPURecord)
	Report() string
}

type RawCPURecorder struct {
	name           string
	start, incr    int
	recordedValues map[int]int
}

func (r *RawCPURecorder) Record(t CPURecord) {
	tickValue := t.Cycle()
	if (tickValue-r.start)%r.incr == 0 {
		r.recordedValues[tickValue] = t.RegisterX()
	}
}

func (r *RawCPURecorder) Report() string {
	result := 0
	for cycle, register := range r.recordedValues {
		result += cycle * register
	}
	return fmt.Sprintf("%d", result)
}

func NewRecorder(name string, start, incr int) Recorder {
	return &RawCPURecorder{name, start, incr, make(map[int]int)}
}
