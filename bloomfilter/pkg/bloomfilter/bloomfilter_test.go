package bloomfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSimpleBloomFilter(t *testing.T) {
	m := uint(1_000)
	p := 0.01
	bf := NewSimpleBloomFilter(m, p)

	assert.NotNil(t, bf)
}

func TestAdd(t *testing.T) {
	m := uint(1_000)
	p := 0.01
	bf := NewSimpleBloomFilter(m, p)
	
	bf.Add([]byte("user_123"))

	assert.NotNil(t, bf)
}

func TestContains(t *testing.T) {
	m := uint(1_000)
	p := 0.01
	bf := NewSimpleBloomFilter(m, p)

	bf.Add([]byte("user_123"))

	assert.True(t, bf.Contains([]byte("user_123")), "should contain user_123")
}
