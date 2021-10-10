package auth

import (
	"crypto/sha256"
	"math/rand"
	"time"
)

const RAND_ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPRQSTUVWXYZ1234567890_-$#~"

func GenerateSalt256() []byte {
	rand.Seed(time.Now().UnixMicro())

	rand_bytes := make([]byte, 32)
	for i := range rand_bytes {
		rand_bytes[i] = RAND_ALPHABET[rand.Int()%len(RAND_ALPHABET)]
	}
	return rand_bytes
}

func CalcHashPassword(pass_raw, salt []byte) []byte {
	// result = sha256( sha256( PASS ) + SALT )
	hpass := sha256.Sum256(pass_raw)

	sum := make([]byte, 0, len(salt)+len(hpass))
	sum = append(sum, hpass[:]...)
	sum = append(sum, salt...)

	hsum := sha256.Sum256(sum)

	return hsum[:]
}

type UserToken uint64

func GenerateNewTokenValue() UserToken {
	return UserToken(rand.Uint64())
}
