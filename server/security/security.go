package security

import (
	"encoding/hex"
	"math/rand"

	"golang.org/x/crypto/sha3"
)

const (
	saltSize = 64
	sha3Size = 64
)

/*
CreateSalt ...
*/
func CreateSalt() string {
	saltSlice := make([]byte, saltSize)
	rand.Read(saltSlice)
	return saltToString(saltSlice)
}

/*
HashPassword ...
*/
func HashPassword(salt, password string) string {
	passwordSlice := []byte(salt + password)
	sha3hash := sha3.Sum512(passwordSlice)
	return sha3ToString(sha3hash)
}

func saltToString(saltSlice []byte) string {
	for i := 0; i < saltSize/2; i++ {
		saltSlice[i], saltSlice[saltSize-1-i] = saltSlice[saltSize-1-i], saltSlice[i]
	}
	return hex.EncodeToString(saltSlice[:])
}

func sha3ToString(sha3Hash [64]byte) string {
	for i := 0; i < sha3Size/2; i++ {
		sha3Hash[i], sha3Hash[sha3Size-1-i] = sha3Hash[sha3Size-1-i], sha3Hash[i]
	}
	return hex.EncodeToString(sha3Hash[:])
}
