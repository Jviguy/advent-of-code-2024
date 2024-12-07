package dsa

import (
	"strconv"
	"strings"
)

type MinHeap struct {
	heap []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{heap: []int{}}
}

func (h *MinHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.minHeapify()
}

func (h *MinHeap) minHeapify() {
	for i := len(h.heap) - 1; i > 0; i-- {
		parent := (i - 1) / 2
		if h.heap[parent] > h.heap[i] {
			h.heap[parent], h.heap[i] = h.heap[i], h.heap[parent]
		}
	}
}

func (h *MinHeap) ExtractMin() int {
	if h.IsEmpty() {
		return -1
	}
	m := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.minHeapify()
	return m
}

func (h *MinHeap) Peek() int {
	if h.IsEmpty() {
		return -1
	}
	return h.heap[0]
}

func (h *MinHeap) Size() int {
	return len(h.heap)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.heap) == 0
}

func (h *MinHeap) Clear() {
	h.heap = []int{}
}

func (h *MinHeap) Values() []int {
	return h.heap
}

func (h *MinHeap) String() string {
	builder := strings.Builder{}
	for _, v := range h.heap {
		builder.WriteString(strconv.Itoa(v))
		builder.WriteString(" ")
	}
	return strings.TrimSpace(builder.String())
}
