package librex

import "sync"

// Rootexec instance store
type RexMap struct {
	LockPtr *sync.Mutex
	MapPtr  map[string]*Rex
}

func (rm RexMap) Delete(key string) *Rex {
	rm.LockPtr.Lock()
	delete(rm.MapPtr, key)
	rm.LockPtr.Unlock()
	return nil
}

func (rm RexMap) Load(key string) *Rex {
	rm.LockPtr.Lock()
	rp := rm.MapPtr[key]
	rm.LockPtr.Unlock()
	return rp
}

func (rm RexMap) Store(key string, rp *Rex) {
	rm.LockPtr.Lock()
	rm.MapPtr[key] = rp
	rm.LockPtr.Unlock()
}

func (rm RexMap) Available(key string) bool {
	_, avail := rm.MapPtr[key]
	return avail
}

// Autoboot

// Autohalt

func MakeRexMap(size int) RexMap {
	return RexMap{
		LockPtr: new(sync.Mutex),
		MapPtr:  make(map[string]*Rex, size),
	}
}
