package hash

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var p = &params{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func Hash(password string) (*string, error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return nil, err
	}
	h := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	hash := fmt.Sprintf("%s$%s", string(h), string(salt))
	return &hash, nil
}

func Compare(hash string, password string) bool {
	s := strings.Split(hash, "$")
	salt := []byte(s[1])
	hashByte := []byte(s[0])
	passwordByte := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)
	return subtle.ConstantTimeCompare(hashByte, passwordByte) == 1
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
