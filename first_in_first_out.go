package cache_replacement_policies

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type fifoCachePolicy struct {
	keys []string
}

func NewFIFOCachePolicy() CachePolicy {
	return &fifoCachePolicy{}
}

func (cp fifoCachePolicy) PickKeyToInvalidate() string {
	return cp.keys[0]
}

func (cp *fifoCachePolicy) OnKeySet(key string) {
	cp.keys = append(cp.keys, key)
}

func (cp fifoCachePolicy) OnKeyGet(key string) {}

func (cp *fifoCachePolicy) OnKeyInvalidate(key string) error {
	indexToInvalidate := slices.Index(cp.keys, key)
	if indexToInvalidate == -1 {
		return fmt.Errorf("can't invalidate key %s not present in cache", key)
	}
	cp.keys = slices.Delete(cp.keys, indexToInvalidate, indexToInvalidate+1)
	return nil
}
