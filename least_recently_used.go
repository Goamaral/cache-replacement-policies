package cache_replacement_policies

import "time"

type lruCachePolicy struct {
	lastUsedAt           map[string]int64
	leastRecentlyUsedKey string
}

func NewLRUCachePolicy() CachePolicy {
	return &lruCachePolicy{lastUsedAt: map[string]int64{}}
}

func (cp lruCachePolicy) PickKeyToInvalidate() string {
	return cp.leastRecentlyUsedKey
}

func (cp *lruCachePolicy) OnKeySet(key string) {
	if cp.leastRecentlyUsedKey == "" {
		cp.leastRecentlyUsedKey = key
	}
	cp.lastUsedAt[key] = time.Now().Unix()
}

func (cp *lruCachePolicy) OnKeyGet(key string) {
	if cp.leastRecentlyUsedKey == "" {
		cp.leastRecentlyUsedKey = key
	}
	cp.lastUsedAt[key] = time.Now().Unix()
}

func (cp *lruCachePolicy) OnKeyInvalidate(key string) error {
	delete(cp.lastUsedAt, key)
	leastRecentlyUsed := struct {
		key       string
		timestamp int64
	}{timestamp: time.Now().Unix()}
	for key, timestamp := range cp.lastUsedAt {
		if timestamp < leastRecentlyUsed.timestamp {
			leastRecentlyUsed.timestamp = timestamp
			leastRecentlyUsed.key = key
		}
	}
	cp.leastRecentlyUsedKey = leastRecentlyUsed.key

	return nil
}
