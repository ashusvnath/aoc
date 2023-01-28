package models

type ComparisonResult int

const (
	Less          ComparisonResult = -1
	Equal         ComparisonResult = 0
	More          ComparisonResult = 1
	NotComparable ComparisonResult = 2
)

var strValues = map[ComparisonResult]string{
	Less:          "less than",
	More:          "more than",
	Equal:         "equal",
	NotComparable: "not comparable to",
}

func (cr ComparisonResult) String() string {
	return strValues[cr]
}

type Entry interface {
	Compare(Entry) ComparisonResult
	AsList() List
	String() string
}
