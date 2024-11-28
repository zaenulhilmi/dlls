package utils

import (
	"crypto/sha256"
	"dlls/contracts"
	"encoding/hex"
)

func NewHasher() contracts.Hasher {
	return &hasherImpl{}
}

type hasherImpl struct{}

func (h *hasherImpl) Hash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (h *hasherImpl) Compare(password, hashedPassword string) bool {
	return h.Hash(password) == hashedPassword
}
