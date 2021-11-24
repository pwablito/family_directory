package util

import (
	"crypto/sha256"
	"fmt"
)

func HashString(input string) string {
	digest := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", digest)
}
