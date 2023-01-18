package main

import "log"

type Recorder struct {
	uniquePos Set[complex128]
}

func (r *Recorder) Record(knot *Knot) {
	log.Printf("Recording %v", knot.position)
	r.uniquePos.Add(knot.position)
}

func (r *Recorder) Count() int {
	return len(r.uniquePos)
}

func NewRecorder() *Recorder {
	positions := make(Set[complex128])
	positions.Add(0)
	return &Recorder{positions}
}
