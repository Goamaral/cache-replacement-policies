package cache_replacement_policies_test

import (
	"cache_replacement_policies"
	"fmt"
	"strconv"

	"math/rand"
	"testing"
	"time"
)

type Test struct {
	TestName    string
	CachePolicy cache_replacement_policies.CachePolicy
}

type DataPair struct {
	Op    int // 0-Set 1-Get
	Key   string
	Value int
}

func TestCachePolicies(t *testing.T) {
	const cacheSize = 5
	const nOps = 50
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var dataset []DataPair
	insertedValues := []int{}

	for n := 0; n < nOps; n++ {
		op := 0
		if n != 0 {
			op = rng.Intn(2)
		}

		var value int
		if op == 0 { // Set
			value = rng.Intn(cacheSize * 4)
			insertedValues = append(insertedValues, value)
		} else { // Get
			i := rng.Intn(len(insertedValues))
			value = insertedValues[i]
		}
		key := strconv.FormatInt(int64(value), 10)

		dataset = append(dataset, DataPair{Op: op, Key: key, Value: value})
	}

	tests := []Test{
		{
			TestName:    "Random Replacement",
			CachePolicy: cache_replacement_policies.NewRRCachePolicy(),
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			cacheHits := 0
			totalGetOps := 0

			cache := cache_replacement_policies.NewCache(cacheSize, test.CachePolicy)
			for _, d := range dataset {
				if d.Op == 0 { // Set
					cache.Set(d.Key, d.Value)
				} else { // Get
					if isCacheHit, _ := cache.Get(d.Key); isCacheHit {
						cacheHits++
					}
					totalGetOps++
				}
			}

			fmt.Printf("%s %.2f%% Cache Hit\n", test.TestName, float64(cacheHits*100)/float64(totalGetOps))
		})
	}
}
