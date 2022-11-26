package skiplist

import "fmt"

func (s *SkipList) search(key string) (*Node , error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	nowPtr := s.Head
	nowLevel := s.MaxLevel

	for nowLevel >= 0 {
		if nowPtr == s.Tail {
			break
		}

		if len(nowPtr.nextPtr) <= nowLevel {
			nowLevel--
			continue
		}

		nextPtr := nowPtr.nextPtr[nowLevel]
		if nextPtr == s.Tail || nextPtr.key > key{
			nowLevel--
			continue
		}
		
		if nextPtr.key == key{
			return nextPtr,nil
		}

		nowPtr = nextPtr
	}

	return nil, fmt.Errorf("no such key")
}
