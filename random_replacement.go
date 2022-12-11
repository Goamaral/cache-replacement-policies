package cache_replacement_policies

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/exp/slices"
)

type rrCachePolicy struct {
	rng  *rand.Rand
	keys []string // TODO: Change to binary search tree. Also search other data structures to improve search.
}

func NewRRCachePolicy() CachePolicy {
	return &rrCachePolicy{
		rng: rand.New(
			rand.NewSource(time.Now().UnixNano()),
		),
	}
}

func (cp rrCachePolicy) PickKeyToInvalidate() string {
	return cp.keys[cp.rng.Intn(len(cp.keys)-1)]
}

func (cp *rrCachePolicy) OnKeySet(key string) {
	cp.keys = append(cp.keys, key)
}

func (cp *rrCachePolicy) OnKeyInvalidate(key string) error {
	indexToInvalidate := slices.Index(cp.keys, key)
	if indexToInvalidate == -1 {
		return fmt.Errorf("can't invalidate key %s not present in cache", key)
	}
	cp.keys = slices.Delete(cp.keys, indexToInvalidate, indexToInvalidate+1)
	return nil
}
