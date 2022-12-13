package cache_replacement_policies

import "math"

// The least used item will be evicted
type lfuCachePolicy struct {
	timesUsed    map[string]int
	leastUsedKey string
}

func NewLFUCachePolicy() CachePolicy {
	return &lfuCachePolicy{timesUsed: map[string]int{}}
}

func (cp lfuCachePolicy) PickKeyToEvict() string {
	return cp.leastUsedKey
}

func (cp *lfuCachePolicy) OnKeySet(key string) {
	if cp.leastUsedKey == "" {
		cp.leastUsedKey = key
	}
	cp.timesUsed[key]++
}

func (cp *lfuCachePolicy) OnKeyGet(key string) {
	if cp.leastUsedKey == "" {
		cp.leastUsedKey = key
	}
	cp.timesUsed[key]++
}

func (cp *lfuCachePolicy) OnKeyEviction(key string) error {
	delete(cp.timesUsed, key)
	leastUsed := struct {
		key    string
		nTimes int
	}{nTimes: math.MaxInt}
	for key, nTimes := range cp.timesUsed {
		if nTimes < leastUsed.nTimes {
			leastUsed.nTimes = nTimes
			leastUsed.key = key
		}
	}
	cp.leastUsedKey = leastUsed.key

	return nil
}
