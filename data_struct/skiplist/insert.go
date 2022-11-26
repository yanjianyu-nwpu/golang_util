package skiplist

import (
	"fmt"
	"log"
)

func (s *SkipList) Put(key string, val any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	node := NewNode(key, val, s.MaxLevel)

	if node.level >= s.MaxLevel+2 {
		panic(fmt.Sprintf("node level error %d, but %d", s.MaxLevel, node.level))
	}

	sLevel := node.level
	if node.level == s.MaxLevel+1 {
		s.Head.nextPtr = append(s.Head.nextPtr, node)
		s.Tail.prePtr = append(s.Tail.prePtr, node)
		node.nextPtr[node.level] = s.Tail
		node.prePtr[node.level] = s.Head
		s.MaxLevel = node.level
		sLevel--
	}

	for j := sLevel; j >= 0; j-- {
		nowPtr := s.Head

		for nowPtr != s.Tail {
			nextPtr := nowPtr.nextPtr[j]

			if nextPtr == s.Tail || nextPtr.key > key {
				nowPtr.nextPtr[j] = node
				node.nextPtr[j] = nextPtr
				nextPtr.prePtr[j] = node
				break
			}

			if nextPtr.key == key {
				log.Fatalf("same key %s", key)
				return fmt.Errorf("dup key")
			}

			nowPtr = nextPtr
		}

	}

	return nil
}
