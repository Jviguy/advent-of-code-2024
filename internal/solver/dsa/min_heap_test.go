package dsa

import (
	"slices"
	"testing"
)

func TestMinHeap_ExtractMin(t *testing.T) {
	h := NewMinHeap()
	// empty heap
	actual := h.ExtractMin()
	expected := -1
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(1)
	h.Insert(5)

	actual = h.ExtractMin()
	expected = 1
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	actual = h.ExtractMin()
	expected = 2
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	actual = h.ExtractMin()
	expected = 3
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	actual = h.ExtractMin()
	expected = 4
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	actual = h.ExtractMin()
	expected = 5
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	actual = h.ExtractMin()
	expected = -1
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestMinHeap_Insert(t *testing.T) {
	h := NewMinHeap()
	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(1)
	h.Insert(5)

	actual := h.ExtractMin()
	expected := 1
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestNewMinHeap(t *testing.T) {
	h := NewMinHeap()
	if h == nil {
		t.Errorf("actual %v, expected %v", h, nil)
	}
}

func TestMinHeap_Peek(t *testing.T) {
	h := NewMinHeap()
	// empty heap
	actual := h.Peek()
	expected := -1
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(1)
	h.Insert(5)

	actual = h.Peek()
	expected = 1
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestMinHeap_Size(t *testing.T) {
	h := NewMinHeap()
	// empty heap
	actual := h.Size()
	expected := 0
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(1)
	h.Insert(5)

	actual = h.Size()
	expected = 5
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestMinHeap_IsEmpty(t *testing.T) {
	h := NewMinHeap()
	// empty heap
	actual := h.IsEmpty()
	expected := true
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}

	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(1)
	h.Insert(5)

	actual = h.IsEmpty()
	expected = false
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestMinHeap_Clear(t *testing.T) {
	h := NewMinHeap()
	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(1)
	h.Insert(5)

	h.Clear()
	actual := h.IsEmpty()
	expected := true
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestMinHeap_Values(t *testing.T) {
	h := NewMinHeap()
	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(1)
	h.Insert(5)

	actual := h.Values()
	expected := []int{1, 2, 3, 4, 5}
	if !slices.Equal(actual, expected) {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

func TestMinHeap_String(t *testing.T) {
	h := NewMinHeap()
	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(1)
	h.Insert(5)

	actual := h.String()
	expected := "1 2 3 4 5"
	if actual != expected {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}
