package cache_replacement_policies

type CachePolicy interface {
	PickKeyToEvict() string
	OnKeySet(key string)
	OnKeyGet(key string)
	OnKeyEviction(key string) error
}
