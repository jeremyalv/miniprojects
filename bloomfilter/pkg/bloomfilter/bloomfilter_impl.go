package bloomfilter

import (
	"fmt"
	"hash/fnv"
	"math"
	"sync"
)

type simpleBloomFilter struct {
	bitset []bool // bit array
	k uint // number of hash functions
	m uint // size of bitset
	lock sync.RWMutex // thread safety
}

// Create a new simple bloom filter for n elements and p false positive probability
func NewSimpleBloomFilter(n uint, p float64) BloomFilter {
	// m = - (n * ln(p)) / (ln(2)^2)
	m := uint(math.Ceil(-1 * (float64(n) * math.Log(p)) / math.Pow(math.Log(2), 2)))
	
	// k = (m / n) * ln(2)
	k := uint(math.Ceil( float64(m) / float64(n) * math.Log(2) ))

	fmt.Printf("created a bloom filter with m = %v, k = %v\n", m, k)

	return &simpleBloomFilter{
		bitset: make([]bool, m),
		k: k,
		m: m,
	}
}

func (bf *simpleBloomFilter) Add(data []byte) {
	bf.lock.Lock()
	defer bf.lock.Unlock()

	h1 := bf.getHashValues(data)

	for i := uint(0); i < bf.k; i++ {
		pos := (h1 + uint64(i)*h1) % uint64(bf.m)

		bf.bitset[pos] = true;
	}
}

func (bf *simpleBloomFilter) Contains(data []byte) bool {
	bf.lock.RLock()
	defer bf.lock.RUnlock()

	h1 := bf.getHashValues(data)

	for i := uint(0); i < bf.k; i++ {
		pos := (h1 + uint64(i)*h1) % uint64(bf.m)

		if !bf.bitset[pos] {
			return false
		}
	}

	return true
}

func (bf *simpleBloomFilter) getHashValues(data []byte) uint64 {
	h := fnv.New64a()
    h.Write(data)
    h1 := h.Sum64()

	return h1
}
