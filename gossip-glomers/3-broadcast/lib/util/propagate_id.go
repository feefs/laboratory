package util

import (
	"broadcast/lib/state"
	"crypto/rand"
	"encoding/base64"
)

func GeneratePropagateID() (state.PropagateID, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	return state.PropagateID(base64.StdEncoding.EncodeToString(b)), nil
}
