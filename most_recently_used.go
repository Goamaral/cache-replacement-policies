package cache_replacement_policies

import "time"

// The most recently used item will be evicted
type mruCachePolicy struct {
	lastUsedAt          map[string]int64
	mostRecentlyUsedKey string
}

func NewMRUCachePolicy() CachePolicy {
	return &mruCachePolicy{lastUsedAt: map[string]int64{}}
}

func (cp mruCachePolicy) PickKeyToEvict() string {
	return cp.mostRecentlyUsedKey
}

func (cp *mruCachePolicy) OnKeySet(key string) {
	if cp.mostRecentlyUsedKey == "" {
		cp.mostRecentlyUsedKey = key
	}
	cp.lastUsedAt[key] = time.Now().Unix()
}

func (cp *mruCachePolicy) OnKeyGet(key string) {
	if cp.mostRecentlyUsedKey == "" {
		cp.mostRecentlyUsedKey = key
	}
	cp.lastUsedAt[key] = time.Now().Unix()
}

func (cp *mruCachePolicy) OnKeyEviction(key string) error {
	delete(cp.lastUsedAt, key)
	mostRecentlyUsed := struct {
		key       string
		timestamp int64
	}{timestamp: time.Now().Unix()}
	for key, timestamp := range cp.lastUsedAt {
		if timestamp > mostRecentlyUsed.timestamp {
			mostRecentlyUsed.timestamp = timestamp
			mostRecentlyUsed.key = key
		}
	}
	cp.mostRecentlyUsedKey = mostRecentlyUsed.key

	return nil
}
