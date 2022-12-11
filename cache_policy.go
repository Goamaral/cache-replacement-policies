package cache_replacement_policies

type CachePolicy interface {
	PickIndexToInvalidate(items []cacheItem) int
}
