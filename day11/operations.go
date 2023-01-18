package main

type Operation func(int) int

func Add(n int) Operation {
	return func(inp int) int {
		return inp + n
	}
}

func Mul(n int) Operation {
	return func(inp int) int {
		return inp * n
	}
}

func Square() Operation {
	return func(inp int) int {
		return inp * inp
	}
}

func Divide(n int) Operation {
	return func(inp int) int {
		return inp / n
	}
}

type Test func(int) bool

func DivisibleBy(n int) Test {
	return func(i int) bool {
		return i%n == 0
	}
}

type Action func(int)

func Conditional(test Test, whenTrue, whenFalse Action) Action {
	return func(input int) {
		if test(input) {
			whenTrue(input)
		} else {
			whenFalse(input)
		}
	}
}
