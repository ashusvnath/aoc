package utility

type Notifiable[T any] func(T)

type Observable[T any] struct {
	observed  T
	observers []Notifiable[T]
}

func (o *Observable[T]) Register(n Notifiable[T]) {
	o.observers = append(o.observers, n)
}

func (o *Observable[T]) Notify() {
	for _, notify := range o.observers {
		notify(o.observed)
	}
}

func (o *Observable[T]) NotifyWith(elem T) {
	for _, notify := range o.observers {
		notify(elem)
	}
}

func NewObservable[T any](target T) *Observable[T] {
	return &Observable[T]{
		observed:  target,
		observers: []Notifiable[T]{},
	}
}
