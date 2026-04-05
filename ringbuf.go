package ringbuf

import (
	"iter"
)

type T[E any] struct {
	len int
	cap int
	buf []E
}

func New[E any](cap int) *T[E] {
	x := T[E]{}
	x.Init(cap)
	return &x
}

func (r *T[E]) Init(cap int) {
	r.len = 0
	r.cap = cap
	r.buf = make([]E, cap)
}

func (r *T[E]) Add(e E) {
	r.buf[r.len%r.cap] = e
	r.len++
}

// Add an element and return a pointer to it, that is valid till the next
// operation on this ringbuffer.
func (r *T[E]) AddAndRef(e E) *E {
	idx := r.len % r.cap
	r.buf[idx] = e
	r.len++
	return &r.buf[idx]
}

// Reference index in the past.
// 0 current element
// 1 last element
func (r *T[E]) BackRef(revidx int) *E {
	currentLen := r.Len()

	if revidx < 0 || revidx >= currentLen {
		return nil
	}

	idx := ((r.len-1-revidx)%r.cap + r.cap) % r.cap

	return &r.buf[idx]
}

func (r *T[E]) Len() int {
	if r.len > r.cap {
		return r.cap
	}
	return r.len
}

func (r *T[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		if r.len <= r.cap {
			for _, e := range r.buf[:r.len] {
				if !yield(e) {
					return
				}
			}
		} else {
			start := r.len % r.cap
			for _, e := range r.buf[start:] {
				if !yield(e) {
					return
				}
			}
			for _, e := range r.buf[:start] {
				if !yield(e) {
					return
				}
			}
		}
	}
}

func (r *T[E]) Reverse() iter.Seq[E] {
	return func(yield func(E) bool) {
		length := r.Len()
		for i := 0; i < length; i++ {
			if !yield(*r.BackRef(i)) {
				return
			}
		}
	}
}

func (r *T[E]) UnsortedSlice() []E {
	return r.buf
}
