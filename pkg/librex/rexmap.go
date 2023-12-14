package librex

import "sync"

// Rootexec instance store
type RexMap struct {
	lockPtr *sync.Mutex
	mapPtr  map[string]*Rex
}

func (rm RexMap) Delete(key string) *Rex {
	rm.lockPtr.Lock()
	delete(rm.mapPtr, key)
	rm.lockPtr.Unlock()
	return nil
}

func (rm RexMap) Load(key string) *Rex {
	rm.lockPtr.Lock()
	rp := rm.mapPtr[key]
	rm.lockPtr.Unlock()
	return rp
}

func (rm RexMap) Store(key string, rp *Rex) {
	rm.lockPtr.Lock()
	rm.mapPtr[key] = rp
	rm.lockPtr.Unlock()
}

func (rm RexMap) Available(key string) bool {
	_, avail := rm.mapPtr[key]
	return avail
}

func MakeRexMap() RexMap {
	return RexMap{
		lockPtr: new(sync.Mutex),
		mapPtr:  make(map[string]*Rex),
	}
}
