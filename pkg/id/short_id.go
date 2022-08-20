package id

import (
	"crypto/rand"
)

const (
	// DefaultShortNameLen is defult length of short name
	DefaultShortNameLen = 8
)
const (
	alphaTab   = "abcdfghjklmnpqrstvwxzABCDFGHJKLMNPQRSTVWXZ012456789"
	upperTab   = "ABCDFGHJKLMNPQRSTVWXZ012456789"
	trimBits   = 0
	tabLen     = len(alphaTab)
	headTabLen = tabLen - 10 // first byte dont allow number character

)

// ShortID  generate a random string with length n.
func ShortID(n int) string {
	s, err := ShortIDWithError(n)
	if err != nil {
		panic(err)
	}
	return s
}

// UpperShortID  generate a random string with length n.
func UpperShortID(n int) string {
	s, err := ShortIDWithDic(n, upperTab, len(upperTab)-10)
	if err != nil {
		panic(err)
	}
	return s
}

// ShortIDWithError  generate a random string with length n
func ShortIDWithError(n int) (string, error) {
	return ShortIDWithDic(n, alphaTab, headTabLen)
}

func withDefaultLength(n int) int {
	if n <= 0 {
		n = DefaultShortNameLen
	}
	return n
}

// ShortIDWithDic generate a shortID with dictionary.
func ShortIDWithDic(n int, dic string, headDicLen int) (string, error) {
	n = withDefaultLength(n)
	b := make([]byte, n)
	if nr, err := rand.Read(b); err != nil || nr != len(b) {
		return "", err
	}

	mod := headDicLen
	dicLen := len(dic)
	for i, v := range b {
		idx := (int(v>>trimBits) % mod)
		mod = dicLen
		b[i] = dic[idx]
	}
	return string(b), nil
}
