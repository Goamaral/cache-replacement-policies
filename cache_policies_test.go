package cache_replacement_policies_test

import (
	"cache_replacement_policies"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const CacheSize = 5

func TestRRCachePolicy(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	cache := cache_replacement_policies.NewCache(CacheSize, cache_replacement_policies.NewRRCachePolicy())
	for n := 0; n < CacheSize*2; n++ {
		value := rand.Intn(CacheSize * 4)
		key := strconv.FormatInt(int64(value), 10)
		cache.Set(key, value)
		fmt.Printf("%d -> %+v\n", value, cache.GetItems())
	}
}
