package random

import (
	"math/rand"

	"github.com/google/uuid"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Uuid() string {
	return uuid.NewString()
}

func PositiveInt() int {
	return rand.Intn(100) + 1
}

func String() string {
	n := PositiveInt()
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
