package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomString() string {
	buff := make([]byte, 5)
	rand.Read(buff)
	return base64.RawURLEncoding.EncodeToString(buff)
}
