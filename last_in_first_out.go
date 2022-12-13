package cache_replacement_policies

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Cache works like a stack, last item to enter will be the first to be evicted
type lifoCachePolicy struct {
	keys []string
}

func NewLIFOCachePolicy() CachePolicy {
	return &lifoCachePolicy{}
}

func (cp lifoCachePolicy) PickKeyToEvict() string {
	return cp.keys[len(cp.keys)-1]
}

func (cp *lifoCachePolicy) OnKeySet(key string) {
	cp.keys = append(cp.keys, key)
}

func (cp lifoCachePolicy) OnKeyGet(key string) {}

func (cp *lifoCachePolicy) OnKeyEviction(key string) error {
	indexToEvict := slices.Index(cp.keys, key)
	if indexToEvict == -1 {
		return fmt.Errorf("can't evict key %s not present in cache", key)
	}
	cp.keys = slices.Delete(cp.keys, indexToEvict, indexToEvict+1)
	return nil
}
