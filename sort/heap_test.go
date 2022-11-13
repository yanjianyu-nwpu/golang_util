package sort

import "testing"

func TestHeapSort(t *testing.T) {
	nums := []int{3, 1, 10, 20, 15, 6, 17, 28}

	intLessF := func(a, b int) bool { return a < b }
	ts := HeapSort(nums, intLessF)
	t.Log(ts)
	MakeHeap(ts, intLessF)
	t.Log(ts)

	ts, top, _ := HeapPop(ts, intLessF)
	t.Log(ts)
	t.Log(top)

	ts = HeapPush(ts, 14, intLessF)
	t.Log(ts)
}
