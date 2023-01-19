package main

type MonkeyBuilder struct {
	monkey *Monkey
	cb     *ConditionalBuilder
}

func (mb *MonkeyBuilder) Id(id string) *MonkeyBuilder {
	mb.monkey.id = id
	return mb
}

func (mb *MonkeyBuilder) Op(op Operation) *MonkeyBuilder {
	mb.monkey.op = op
	return mb
}

func (mb *MonkeyBuilder) Action(action Action) *MonkeyBuilder {
	mb.monkey.action = action
	return mb
}

func (mb *MonkeyBuilder) Items(items ...int) *MonkeyBuilder {
	mb.monkey.items = items
	return mb
}

func (mb *MonkeyBuilder) Test(test Test) *MonkeyBuilder {
	mb.cb.Test(test)
	return mb
}

func (mb *MonkeyBuilder) WhenTrue(a Action) *MonkeyBuilder {
	mb.cb.WhenTrue(a)
	return mb
}

func (mb *MonkeyBuilder) WhenFalse(a Action) *MonkeyBuilder {
	mb.cb.WhenFalse(a)
	return mb
}

func (mb *MonkeyBuilder) Build() *Monkey {
	mb.monkey.action = mb.cb.Build()
	return mb.monkey
}

type ConditionalBuilder struct {
	test                Test
	whenTrue, whenFalse Action
}

func (cb *ConditionalBuilder) Test(test Test) *ConditionalBuilder {
	cb.test = test
	return cb
}

func (cb *ConditionalBuilder) WhenTrue(action Action) *ConditionalBuilder {
	cb.whenTrue = action
	return cb
}

func (cb *ConditionalBuilder) WhenFalse(action Action) *ConditionalBuilder {
	cb.whenFalse = action
	return cb
}

func (cb *ConditionalBuilder) Build() Action {
	return Conditional(cb.test, cb.whenTrue, cb.whenFalse)
}
