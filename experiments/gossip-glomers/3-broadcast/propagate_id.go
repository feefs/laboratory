package main

import (
	"crypto/rand"
	"encoding/base64"
)

func generatePropagateID() (propagateID, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b[:])
	if err != nil {
		return "", err
	}

	return propagateID(base64.StdEncoding.EncodeToString(b)), nil
}
