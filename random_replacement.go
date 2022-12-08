package cache_replacement_policies

import (
	"math/rand"
)

type rrCachePolicy struct{}

func NewRRCachePolicy() CachePolicy {
	return rrCachePolicy{}
}

func (rrCachePolicy) PickIndexToDiscard(items []cacheItem) int {
	return rand.Intn(len(items) - 1)
}
