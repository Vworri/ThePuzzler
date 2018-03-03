package puzzler

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"

	"golang.org/x/crypto/sha3"
)

func Md5Hash(message string) string {
	h := md5.New()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha1Hash(message string) string {
	h := sha1.New()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256Hash(message string) string {
	h := sha256.New()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func Sha512Hash(message string) string {
	h := sha512.New()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func Sha3256(message string) string {
	h := sha3.New256()
	io.WriteString(h, message)
	return fmt.Sprintf("%x", h.Sum(nil))
}
