package cache_replacement_policies

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/exp/slices"
)

// A random item is evicted
type rrCachePolicy struct {
	rng  *rand.Rand
	keys []string
}

func NewRRCachePolicy() CachePolicy {
	return &rrCachePolicy{
		rng: rand.New(
			rand.NewSource(time.Now().UnixNano()),
		),
	}
}

func (cp rrCachePolicy) PickKeyToEvict() string {
	return cp.keys[cp.rng.Intn(len(cp.keys)-1)]
}

func (cp *rrCachePolicy) OnKeySet(key string) {
	cp.keys = append(cp.keys, key)
}

func (cp rrCachePolicy) OnKeyGet(key string) {}

func (cp *rrCachePolicy) OnKeyEviction(key string) error {
	indexToEvict := slices.Index(cp.keys, key)
	if indexToEvict == -1 {
		return fmt.Errorf("can't evict key %s not present in cache", key)
	}
	cp.keys = slices.Delete(cp.keys, indexToEvict, indexToEvict+1)
	return nil
}
