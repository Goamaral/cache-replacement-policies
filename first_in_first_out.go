package cache_replacement_policies

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// Cache works like a queue, first item to enter will be the first to be evicted
type fifoCachePolicy struct {
	keys []string
}

func NewFIFOCachePolicy() CachePolicy {
	return &fifoCachePolicy{}
}

func (cp fifoCachePolicy) PickKeyToEvict() string {
	return cp.keys[0]
}

func (cp *fifoCachePolicy) OnKeySet(key string) {
	cp.keys = append(cp.keys, key)
}

func (cp fifoCachePolicy) OnKeyGet(key string) {}

func (cp *fifoCachePolicy) OnKeyEviction(key string) error {
	indexToEvict := slices.Index(cp.keys, key)
	if indexToEvict == -1 {
		return fmt.Errorf("can't evict key %s not present in cache", key)
	}
	cp.keys = slices.Delete(cp.keys, indexToEvict, indexToEvict+1)
	return nil
}
