package main

import (
	"fmt"

	"github.com/jeremyalv/spellchecker/pkg/bloomfilter"
)

func main() {
	bf := bloomfilter.NewSimpleBloomFilter(1_000, 0.01)

	bf.Add([]byte("user_123"))
	bf.Add([]byte("user_456"))


	// Test existence
	item1 := []byte("user_123")
	item2 := []byte("user_999")

	fmt.Printf("Contains %s? %v\n", item1, bf.Contains(item1)) // true
	fmt.Printf("Contains %s? %v\n", item2, bf.Contains(item2)) // high probability of outputting false
}

