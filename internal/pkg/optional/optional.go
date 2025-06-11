package optional

// Java style Optional implementation
type Optional[T any] struct {
	value   T
	present bool
}

func (o *Optional[T]) IsPresent() bool {
	return o.present
}

func (o *Optional[T]) Get() T {
	if !o.present {
		panic("No value present")
	}
	return o.value
}

func (o *Optional[T]) OrElse(other T) T {
	if o.present {
		return o.value
	}
	return other
}

func (o *Optional[T]) OrElseGet(supplier func() T) T {
	if o.present {
		return o.value
	}
	return supplier()
}

func (o *Optional[T]) OrElseThrow(err error) T {
	if o.present {
		return o.value
	}
	panic(err)
}
