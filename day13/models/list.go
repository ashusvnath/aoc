package models

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type List []Entry

var tabLevel int = 0

func (l List) String() string {
	strs := []string{}
	for _, v := range l {
		strs = append(strs, v.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}

func (l List) Compare(other Entry) ComparisonResult {
	if l == nil || other == nil {
		return NotComparable
	}
	prefix := strings.Repeat("  ", tabLevel)
	log.Printf("%s- Compare %v vs %v", prefix, l, other)

	tabLevel++
	defer func() { tabLevel-- }()
	prefix = strings.Repeat("  ", tabLevel)

	otherAsList, ok := other.(List)
	if !ok {
		log.Printf("%s- Mixed types; convert right to [%v] and retry comparison", prefix, other)
		otherAsList = other.AsList()
	}
	lOther := len(otherAsList)
	lSelf := len(l)
	minLen := int(math.Min(float64(lSelf), float64(lOther)))
	for i := 0; i < minLen; i++ {
		res := l[i].Compare(otherAsList[i])
		if res != Equal {
			log.Printf("%s- Left is %v right", prefix, res)
			return res
		}
	}
	if lSelf < lOther {
		log.Printf("%s- Left side ran out of items", prefix)
		return Less
	}
	if lOther < lSelf {
		log.Printf("%s- Right side ran out of items", prefix)
		return More
	}
	return Equal
}

func (l List) AsList() List {
	return l
}

func (l *List) Append(entry Entry) {
	(*l) = append((*l), entry)
}
