package common_helper

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateActivationKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
