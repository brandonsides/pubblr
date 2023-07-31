package database

import "math/rand"

const SALT_LENGTH = 32

var GenerateSalt = func() string {
	ret := make([]byte, SALT_LENGTH)
	for i := 0; i < SALT_LENGTH; i++ {
		ret[i] = 'a' + byte(rand.Intn(26))
	}
	return string(ret)
}
