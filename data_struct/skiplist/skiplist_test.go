package skiplist

import (
	"fmt"
	"testing"
)

func TestNowFile(t *testing.T) {
	sl := NewSkipList()

	sl.Put("1", 1)
	fmt.Println(sl.PrintSkipList())
	sl.Put("2", 2)
	sl.Put("3", 3)

	t.Log(sl.PrintSkipList())
}
