package util

import (
	"crypto/rand"
	"encoding/base64"

	"broadcast/types"
)

func GeneratePropagateID() (types.PropagateID, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	return types.PropagateID(base64.StdEncoding.EncodeToString(b)), nil
}
