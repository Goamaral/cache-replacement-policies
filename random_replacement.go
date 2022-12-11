package cache_replacement_policies

import (
	"math/rand"
	"time"
)

type rrCachePolicy struct {
	rng *rand.Rand
}

func NewRRCachePolicy() CachePolicy {
	return rrCachePolicy{
		rng: rand.New(
			rand.NewSource(time.Now().UnixNano()),
		),
	}
}

func (cp rrCachePolicy) PickIndexToInvalidate(items []cacheItem) int {
	return cp.rng.Intn(len(items) - 1)
}
