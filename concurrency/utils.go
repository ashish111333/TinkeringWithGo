package concurrency

import (
	"crypto/rand"
)

func RandString(prefix string) string {
	if prefix == "" {
		return rand.Text()
	}
	return prefix + rand.Text()
}
