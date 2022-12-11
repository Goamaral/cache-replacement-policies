package cache_replacement_policies_test

import (
	"cache_replacement_policies"
	"fmt"
	"strconv"

	"math/rand"
	"testing"
	"time"
)

type CacheOp uint

const (
	SetOp CacheOp = iota
	GetOp
)

type Test struct {
	TestName    string
	CachePolicy cache_replacement_policies.CachePolicy
}

type DataPair struct {
	Op    CacheOp
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
		op := SetOp
		if n != 0 {
			op = CacheOp(rng.Intn(2))
		}

		var value int
		if op == SetOp { // Set
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
		{
			TestName:    "First In First Out",
			CachePolicy: cache_replacement_policies.NewFIFOCachePolicy(),
		},
	}
	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			cacheHits := 0
			totalGetOps := 0

			cache := cache_replacement_policies.NewCache(cacheSize, test.CachePolicy)
			for _, d := range dataset {
				if d.Op == SetOp {
					cache.Set(d.Key, d.Value)
				} else {
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
