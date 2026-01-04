package ringbuf

import (
	"testing"
)

func TestSmallRing(t *testing.T) {
	r := New[int](3)
	testIts(t, r)
	r.Add(1)
	testIts(t, r, 1)
	r.Add(2)
	r.Add(3)
	testIts(t, r, 1, 2, 3)
	r.Add(4)
	testIts(t, r, 2, 3, 4)
	r.Add(5)
	testIts(t, r, 3, 4, 5)
	r.Add(6)
	testIts(t, r, 4, 5, 6)
	r.Add(7)
	testIts(t, r, 5, 6, 7)
	r.Add(8)
	testIts(t, r, 6, 7, 8)
	r.Add(9)
	testIts(t, r, 7, 8, 9)
	r.Add(10)
	testIts(t, r, 8, 9, 10)
}

func testLenAndIts[E any](t *testing.T, r *T[E], l int) {
	if r.Len() != l {
		t.Fatal("Length should be", l)
	}
	it := 0
	for _ = range r.All() {
		it++
	}
	if it != l {
		t.Fatal("Iterations should be", l)
	}
}

func testIts[E comparable](t *testing.T, r *T[E], es ...E) {
	l := len(es)
	if r.Len() != l {
		t.Fatal("Length should be", l)
	}
	it := 0
	for x := range r.All() {
		if x != es[it] {
			t.Fatal("Expected x==es[i]: ", x, "==", es[it])
		}
		it++
	}
	if it != l {
		t.Fatal("Iterations should be", l)
	}
	for x := range r.Reverse() {
		it--
		if x != es[it] {
			t.Fatal("Expected x==es[i]: ", x, "==", es[it])
		}
	}
	if it != 0 {
		t.Fatal("Reverse Iterations countdown should be 0")
	}

}
