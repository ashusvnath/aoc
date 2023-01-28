package main_test

import (
	"day13/assert"
	"day13/parser"
	"testing"
)

func TestE2E(t *testing.T) {
	t.Run("should return false for given input", func(t *testing.T) {
		//t.SkipNow()
		input := `[[10,0,2]]
[[[[4,3],5,[2,2,6,9],[],[8,4,4,7,2]],[3],1],[[6,5,[8],[10,9,9],6],[10],[3,[0,4,6,0],[3,4,8],9],0,[3,1,[1,10]]],[8],[[[1,0,6,1]],[[7,2,1,3,6],[7,8,0],9],[[9,7],7,4]]]`
		ps := parser.ParsePairs(input)

		assert.False(ps[0].IsOrderedCorrectly(), t)
	})
}
