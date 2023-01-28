package models

import (
	"fmt"
	"log"
	"strings"
)

type Int int

func (i Int) Compare(other Entry) ComparisonResult {
	prefix := strings.Repeat("  ", tabLevel)
	log.Printf("%s- Compare %v vs %v", prefix, i, other)
	otherAsInt, ok := other.(Int)
	if !ok {
		log.Printf("%s- Mixed types; convert left to [%v] and retry comparison", prefix, i)
		return i.AsList().Compare(other)
	}
	if i-otherAsInt == 0 {
		//log.Printf("%s - Left is equal to right", prefix)
		return Equal
	} else if i-otherAsInt < 0 {
		log.Printf("%s  - Left is smaller", prefix)
		return Less
	}
	log.Printf("%s  - Right is smaller", prefix)
	return More
}

func (i Int) AsList() List {
	return List([]Entry{i})
}

func (i Int) String() string {
	return fmt.Sprintf("%v", int(i))
}
