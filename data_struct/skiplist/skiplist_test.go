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

func TestSkipListSearch(t *testing.T) {
	sl := NewSkipList()

	sl.Put("1", 1)
	fmt.Println(sl.PrintSkipList())
	sl.Put("2", 2)
	sl.Put("3", 3)

	u, mm := sl.search("2")
	if mm != nil {
		t.Error(mm)
	}

	if u == nil {
		t.Error("G111")
	}
	if u == nil || u.val == nil {
		t.Error("u val is nil")
	}
	p, s := u.val.(int)
	if !s {
		t.Error("tarns fail")
	}
	if p != 2 {
		t.Error("fewfw")
	}
	t.Log(p)

}
