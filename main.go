package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"hash"
)

var mHasher hash.Hash32

func init() {
	// mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
	mHasher = murmur3.New32WithSeed(uint32(11))
}

func murmurHash(key string, size int32) int {
	mHasher.Write([]byte(key))
	result := mHasher.Sum32() % uint32(size)
	//actual index where we are setting byte true or false
	mHasher.Reset()
	return int(result)
}

// define Bf
type BloomFilter struct {
	filter []bool
	size   int32
}

// intialise bf ,when we define a bf we need to tell how big the bloom filter is, allocate a bloom filter that will hold these many bytes
func NewBloomFilter(size int) *BloomFilter {
	//create a new bloom filter
	return &BloomFilter{
		make([]bool, size),
		int32(size),
	}
}

// create add method to BF
func (b *BloomFilter) Add(key string) {
	idx := murmurHash(key, b.size)
	fmt.Println("wrote", key, "At index", idx)
	b.filter[idx] = true
}

func (b *BloomFilter) Exists(key string) bool {
	idx := murmurHash(key, b.size)
	return b.filter[idx]
}

func main() {
	bloom := NewBloomFilter(16)
	keys := []string{"a", "b", "c", "d", "e", "f"}
	for _, key := range keys {
		bloom.Add(key)
	}
	for _, key := range keys {
		fmt.Println(key, bloom.Exists(key))
	}

	fmt.Println(bloom.Exists("z"))

}
