package cache_replacement_policies

type CachePolicy interface {
	PickIndexToDiscard(items []cacheItem) int
}
