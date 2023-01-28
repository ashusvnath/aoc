package parser

import (
	. "day13/models"
	"day13/utility"
	"log"
	"strings"
)

func Parse(input string) List {
	currentList := List{}
	currentInt := Int(0)
	readingInteger := false
	s := utility.NewStack[List]()
	for _, c := range input[1 : len(input)-1] {
		switch {
		case c == '[':
			readingInteger = false
			s.Push(currentList)
			currentList = List{}
		case c == ']':
			if readingInteger {
				currentList.Append(currentInt)
				readingInteger = false
				currentInt = 0
			}
			completedList := currentList
			currentList = s.Pop()
			currentList.Append(completedList)
		case c >= '0' && c <= '9':
			readingInteger = true
			entry := Int(c - '0')
			currentInt *= 10
			currentInt += entry
		case c == ',':
			if readingInteger {
				readingInteger = false
				currentList.Append(currentInt)
				currentInt = 0
			}
		default:
			log.Fatalf("Encountered invalid value: %c", c)
		}
	}
	if readingInteger {
		currentList.Append(currentInt)
	}
	return currentList
}

func ParsePairs(input string) []*Pair {
	pairs := []*Pair{}
	pairStrings := strings.Split(strings.TrimRight(input, "\n"), "\n\n")
	for _, pairString := range pairStrings {
		temp := strings.Split(pairString, "\n")
		pair := NewPair(Parse(temp[0]), Parse(temp[1]))
		pairs = append(pairs, pair)
	}
	return pairs
}
