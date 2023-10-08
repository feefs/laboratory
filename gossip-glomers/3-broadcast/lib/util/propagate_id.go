package util

import (
	"broadcast/lib/state"
	"crypto/rand"
	"encoding/base64"
)

func GeneratePropagateID() (state.PropagationID, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	return state.PropagationID(base64.StdEncoding.EncodeToString(b)), nil
}
