package librex

import "sync"

// Rootexec instance store
type RexMap struct {
	Lock *sync.Mutex
	Map  map[string]*Rex
}

// Allocate a RexMap
func MakeRexMap(size int) RexMap {
	return RexMap{
		Lock: new(sync.Mutex),
		Map:  make(map[string]*Rex, size),
	}
}
