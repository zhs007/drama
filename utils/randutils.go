package dramautils

import (
	"math/rand"
)

// GenHashCode - generator a hash code
func GenHashCode(length int) string {
	const HASHSTRING = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	str := ""

	for i := 0; i < length; i++ {
		ci := rand.Int() % len(HASHSTRING)
		str += HASHSTRING[ci : ci+1]
	}

	return str
}
