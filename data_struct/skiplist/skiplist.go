package skiplist

import "sync"

type Node struct {
	key string
	val any

	level int
	//score   float64
	nextPtr []*Node
	prePtr  []*Node
}

type SkipList struct {
	mu       sync.Mutex
	Head     *Node
	Tail     *Node
	MaxLevel int
	Num      int
}

func NewNode(key string, val any, maxLevel int) *Node {
	level := getRandHeight(maxLevel)

	return &Node{
		key:     key,
		val:     val,
		level:   level,
		nextPtr: make([]*Node, level+1),
		prePtr:  make([]*Node, level+1),
	}

}

func NewSkipList() *SkipList {
	r := &SkipList{
		Num:      0,
		MaxLevel: 0,
	}
	tail := &Node{
		level: 0,
	}

	head := &Node{
		key:     "",
		val:     nil,
		level:   0,
		nextPtr: []*Node{tail},
	}
	tail.prePtr = []*Node{head}

	r.Head = head
	r.Tail = tail
	return r
}

func (s *SkipList) PrintSkipList() string {
	res := ""

	for i := s.MaxLevel; i >= 0; i-- {
		res += "|"
		tmp := s.Head.nextPtr[i]
		for tmp != s.Tail {
			res += "\t"
			res += tmp.key
			tmp = tmp.nextPtr[i]
		}
		res += "\t|"
		res += "\n"
	}

	return res
}
