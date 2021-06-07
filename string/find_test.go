package string

import "testing"

func TestBruteForceSearch(t *testing.T) {
	main := "abcd227fac"
	pattern := "ac"
	t.Log(BruteForceSearch(main, pattern))

	main = "abcd227fa"
	pattern = "ac"
	t.Log(BruteForceSearch(main, pattern))
}

func TestRabinKarpSearch(t *testing.T) {
	main := "abcd227fac"
	pattern := "ac"
	t.Log(RabinKarpSearch(main, pattern))

	main = "abcd227fa"
	pattern = "ac"
	t.Log(RabinKarpSearch(main, pattern))
}
