package cache_replacement_policies

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type lifoCachePolicy struct {
	keys []string
}

func NewLIFOCachePolicy() CachePolicy {
	return &lifoCachePolicy{}
}

func (cp lifoCachePolicy) PickKeyToInvalidate() string {
	return cp.keys[len(cp.keys)-1]
}

func (cp *lifoCachePolicy) OnKeySet(key string) {
	cp.keys = append(cp.keys, key)
}

func (cp lifoCachePolicy) OnKeyGet(key string) {}

func (cp *lifoCachePolicy) OnKeyInvalidate(key string) error {
	indexToInvalidate := slices.Index(cp.keys, key)
	if indexToInvalidate == -1 {
		return fmt.Errorf("can't invalidate key %s not present in cache", key)
	}
	cp.keys = slices.Delete(cp.keys, indexToInvalidate, indexToInvalidate+1)
	return nil
}
