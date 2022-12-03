package skiplist

func (s *SkipList) search(key string) (*Node, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	nowPtr := s.Head

	for nowLevel := s.MaxLevel; nowLevel >= 0; nowLevel-- {
		for nowPtr != s.Tail && nowPtr != nil && nowPtr.nextPtr[nowLevel] != s.Tail {
			if nowPtr == nil {
				return nil, NoFound
			}

			if nowPtr != s.Head && nowPtr.key == key {
				return nowPtr, nil
			}

			nextPtr := nowPtr.nextPtr[nowLevel]
			if nextPtr == s.Tail || nextPtr.key > key {
				nowLevel--
				break
			}

			nowPtr = nextPtr
		}
	}

	return nil, NoFound
}
