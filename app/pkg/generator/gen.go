package generator

import (
	"crypto/rand"
)

var (
	CHARS  = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")
	LENGTH = byte(len(CHARS))
)

// func init() {
// 	rand.New(rand.NewSource(time.Now().Unix())).
// 	 	Shuffle(len(CHARS), func(i, j int) {
// 		 	CHARS[i], CHARS[j] = CHARS[j], CHARS[i]
// 	})
// }

func Key(length int) ([]byte, error) {
	key := make([]byte, length)

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	for i, n := range key {
		key[i] = CHARS[n%LENGTH]
	}

	return key, nil
}
