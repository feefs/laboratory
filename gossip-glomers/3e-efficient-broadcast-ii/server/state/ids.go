package state

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateBroadcastID() (BroadcastID, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	return BroadcastID(base64.StdEncoding.EncodeToString(b)), nil
}

func GeneratePropagateID() (PropagationID, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	return PropagationID(base64.StdEncoding.EncodeToString(b)), nil
}
