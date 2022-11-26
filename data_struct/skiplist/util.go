package skiplist

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 获取插入元素的高度
func getRandHeight(maxH int) int {
	res := 0

	for res <= maxH {
		tmp := rand.Int()
		if tmp%2 == 1 {
			break
		}

		res++
	}

	return res
}
