package bloomfilter

type BloomFilter interface {
	Add(data []byte)

	Contains(data []byte) bool
}
