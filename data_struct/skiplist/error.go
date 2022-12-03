package skiplist

import "fmt"

var (
	NoFound error
	DupKey  error
)

func init() {
	NoFound = fmt.Errorf("no found")
	DupKey = fmt.Errorf("dup key")
}
