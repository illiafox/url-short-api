package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuzzCache(f *testing.F) {
	f.Add("key", "value")

	cache := New()

	f.Fuzz(func(t *testing.T, key, value string) {
		cache.Set(key, value)

		assert.Equal(t, value, cache.Get(key), "check cache")
	})

}
