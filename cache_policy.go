package cache_replacement_policies

type CachePolicy interface {
	PickKeyToInvalidate() string
	OnKeySet(key string)
	OnKeyInvalidate(key string) error
}
