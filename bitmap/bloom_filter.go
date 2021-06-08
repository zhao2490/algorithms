package bitmap

import (
	"hash"
	"hash/fnv"
	"math"
)

/*
	布隆过滤器
	优点：对位图的优化 可以使用更小的内存
	缺点：对存在的情况有误判，如果某个数字通过布隆过滤器判断不存在，一定不存在，但是判断存在，有可能不存在
*/
type BloomFilter struct {
	bitmap []bool      // the bloom-filter bitmap
	k      int         // number of hash functions
	n      int         // number of elements in the filter
	m      int         // size of the bloom filter
	hashfn hash.Hash64 // hash function
}

// Returns a new Bloom Filter object, if you pass the number of hash functions to use the maximum size of Bloom Filter
func NewBloomFilter(numHashFuncs, bfSize int) *BloomFilter {
	bf := new(BloomFilter)
	bf.bitmap = make([]bool, bfSize)
	bf.k, bf.m = numHashFuncs, bfSize
	bf.n = 0
	bf.hashfn = fnv.New64()
	return bf
}

func (bf *BloomFilter) getHash(b []byte) (uint32, uint32) {
	bf.hashfn.Reset()
	bf.hashfn.Write(b)
	hash64 := bf.hashfn.Sum64()
	h1 := uint32(hash64 & ((1 << 32) - 1))
	h2 := uint32(hash64 >> 32)
	return h1, h2
}

func (bf *BloomFilter) Add(e []byte) {
	h1, h2 := bf.getHash(e)
	for i := 0; i < bf.k; i++ {
		ind := (h1 + uint32(i)*h2) % uint32(bf.m)
		bf.bitmap[ind] = true
	}
	bf.n++
}

// Checks if an element (in byte-array form) exists in the Bloom Filter
func (bf *BloomFilter) Check(x []byte) bool {
	h1, h2 := bf.getHash(x)
	result := true
	for i := 0; i < bf.k; i++ {
		ind := (h1 + uint32(i)*h2) % uint32(bf.m)
		result = result && bf.bitmap[ind]
	}
	return result
}

// Returns the current False Positive Rate of the Bloom Filter
func (bf *BloomFilter) FalsePositiveRate() float64 {
	return math.Pow(1-math.Exp(-float64(bf.k*bf.n)/float64(bf.m)),
		float64(bf.k))
}
