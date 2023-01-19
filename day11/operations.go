package main

import (
	"math/big"
)

type Operation func(*big.Int) *big.Int

func Add(n *big.Int) Operation {
	return func(inp *big.Int) *big.Int {
		return inp.Add(inp, n)
	}
}

func Mul(n *big.Int) Operation {
	return func(inp *big.Int) *big.Int {
		return inp.Mul(inp, n)
	}
}

func Square() Operation {
	return func(inp *big.Int) *big.Int {
		return inp.Exp(inp, big.NewInt(2), nil)
	}
}

func Divide(n *big.Int) Operation {
	return func(inp *big.Int) *big.Int {
		return inp.Div(inp, n)
	}
}

func Identity() Operation {
	return func(i *big.Int) *big.Int { return i }
}

type Test func(*big.Int) bool

func DivisibleBy(n *big.Int) Test {
	zero := big.NewInt(0)
	q, m := big.NewInt(0), big.NewInt(0)
	return func(i *big.Int) bool {
		q.DivMod(i, n, m)
		return m.Cmp(zero) == 0
	}
}

type Action func(*big.Int)

func Conditional(test Test, whenTrue, whenFalse Action) Action {
	return func(input *big.Int) {
		if test(input) {
			whenTrue(input)
		} else {
			whenFalse(input)
		}
	}
}
