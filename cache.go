package cache_replacement_policies

type Cache interface {
	Set(key string, value any)
	Get(key string) (bool, any)
	GetItems() []cacheItem
}

type cacheItem struct {
	key   string
	value any
}

type cache struct {
	itemMapper map[string]int // TODO: Search better data structure to improve key list fetching
	items      []cacheItem
	size       int
	policy     CachePolicy
}

func NewCache(size int, policy CachePolicy) Cache {
	return &cache{
		itemMapper: map[string]int{},
		items:      []cacheItem{},
		size:       size,
		policy:     policy,
	}
}

func (c *cache) Set(key string, value any) {
	if _, found := c.itemMapper[key]; found {
		return
	}
	if len(c.itemMapper) < c.size {
		i := len(c.items)
		c.items = append(c.items, cacheItem{key, value})
		c.itemMapper[key] = i
	} else {
		invalidIndex := c.policy.PickIndexToInvalidate(c.items)
		delete(c.itemMapper, c.items[invalidIndex].key)
		c.itemMapper[key] = invalidIndex
		c.items[invalidIndex] = cacheItem{key, value}
	}
}

func (c *cache) Get(key string) (bool, any) {
	i, found := c.itemMapper[key]
	if found {
		return true, c.items[i]
	}
	return false, nil
}

func (c *cache) GetItems() []cacheItem {
	return c.items
}
