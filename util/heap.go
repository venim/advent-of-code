package util

type MinHeapItem[T any] interface {
	Less(T) bool
}

type MinHeap[T MinHeapItem[T]] []T

func (h MinHeap[T]) Len() int           { return len(h) }
func (h MinHeap[T]) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h MinHeap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

func (h *MinHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = *new(T)
	*h = old[0 : n-1]
	return x
}
