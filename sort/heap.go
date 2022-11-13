// package sort heap.go implement the heap sort method in golang with generics
// in this file support slice sort, the true usage is
// HeapSort -> pop/push
// lessF return true is the element a should before b
package sort

// HeapAdjust 堆发生单点变化之后，调整堆
func HeapAdjust[T any](ts []T, pos int, lessF func(a T, b T) bool) {
	l := len(ts)
	if pos >= l {
		return
	}

	lChildPos := 2*pos + 1
	rChildPos := 2*pos + 2
	tmp := pos
	if lChildPos < l && lessF(ts[lChildPos], ts[tmp]) {
		tmp = lChildPos
	}

	if rChildPos < l && lessF(ts[rChildPos], ts[tmp]) {
		tmp = rChildPos
	}

	if tmp == pos {
		return
	}

	tmpD := ts[pos]
	ts[pos] = ts[tmp]
	ts[tmp] = tmpD
	HeapAdjust(ts, tmp, lessF)
}

// MakeHeap 构造堆
func MakeHeap[T any](ts []T, lessF func(a T, b T) bool) {
	l := len(ts)

	for i := 0; i <= l/2; i++ {
		HeapAdjust(ts, i, lessF)
	}
}

// HeapPush 加入一个新的元素
func HeapPush[T any](ts []T, item T, lessF func(a T, b T) bool) []T {
	ts = append([]T{item}, ts...)

	HeapAdjust(ts, 0, lessF)
	return ts
}

// HeapPop 弹出一个元素
func HeapPop[T any](ts []T, lessF func(a T, b T) bool) ([]T, T, bool) {
	l := len(ts)

	var t T
	if l == 0 {
		return []T{}, t, false
	}

	t = ts[0]
	ts[0] = ts[l-1]
	ts = ts[:l-1]
	HeapAdjust(ts, 0, lessF)

	return ts, t, true
}

// HeapSort 直接堆排序
func HeapSort[T any](ts []T, lessF func(a, b T) bool) []T {
	l := len(ts)
	if l <= 0 {
		return ts
	}

	MakeHeap(ts, lessF)
	res := make([]T, 0, l)

	var item T
	var ok bool
	for i := 0; i < l; i++ {
		ts, item, ok = HeapPop(ts, lessF)
		if !ok {
			break
		}
		res = append(res, item)
	}

	return res
}
