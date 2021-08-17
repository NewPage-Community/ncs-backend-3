package model

import (
	"math/rand"
	"time"
)

// GetRandomString - create random string
func getRandomString(lenght int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	bytesLen := len(bytes)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := int64(0); i < lenght; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}
