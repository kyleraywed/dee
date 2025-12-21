package promise

type Promise[T any] struct {
	promise *T
}

func (p Promise[T]) Get() (T, bool) {
	if p.promise != nil {
		return *p.promise, true
	} else {
		var zero T
		return zero, false
	}
}

func (p *Promise[T]) Set(value T) {
	if p.promise == nil {
		p.promise = new(T)
	}
	*p.promise = value
}
