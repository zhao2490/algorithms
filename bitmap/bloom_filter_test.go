package bitmap

import "testing"

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(3, 100)
	d1, d2, d3 := []byte("Hello"), []byte("zhao"), []byte("zzz")
	bf.Add(d1)
	t.Log(bf.FalsePositiveRate())
	bf.Add(d3)
	t.Log(bf.FalsePositiveRate())
	for _, i := range [][]byte{d1, d2, d3} {
		t.Log(bf.Check(i))
	}
}
