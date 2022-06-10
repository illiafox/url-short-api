package generator

import (
	"math/rand"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

func KeyLength(a *assert.Assertions, l int) {
	key, err := Key(l)

	a.NoError(err, "generate key")

	a.Equal(l, len(key), "check key length")

	a.True(utf8.Valid(key), "valid utf8 format")
}

const limit = 1000

func TestKey(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().Unix())

	KeyLength(assert.New(t), rand.Intn(limit))
}

func FuzzKey(f *testing.F) {
	f.Add(1)

	abs := func(n int) int {
		if n < 0 {
			return -n
		}

		return n
	}

	f.Fuzz(func(t *testing.T, length int) {
		length = abs(length)

		KeyLength(assert.New(t), length)
	})
}
