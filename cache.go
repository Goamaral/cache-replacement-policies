package cache_replacement_policies

type Cache interface {
	Set(key string, value any)
	Get(key string) (bool, any)
}

type cacheItem struct {
	key   string
	value any
}

type cache struct {
	items  map[string]any
	size   int
	policy CachePolicy
}

func NewCache(size int, policy CachePolicy) Cache {
	return &cache{
		items:  map[string]any{},
		size:   size,
		policy: policy,
	}
}

func (c *cache) Set(key string, value any) {
	if _, found := c.items[key]; found {
		return
	}
	if len(c.items) == c.size {
		key := c.policy.PickKeyToInvalidate()
		delete(c.items, key)
		c.policy.OnKeyInvalidate(key)
	}
	c.items[key] = value
	c.policy.OnKeySet(key)
}

func (c *cache) Get(key string) (bool, any) {
	item, found := c.items[key]
	return found, item
}
