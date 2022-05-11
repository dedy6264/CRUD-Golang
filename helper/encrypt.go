package helper

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))

}
func EncryptAes(key string) string {
	b := "sejutacita"
	md := Encrypt(key)
	c, err := aes.NewCipher([]byte(md))
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	out := make([]byte, len(b))
	fmt.Println("########")
	c.Encrypt(out, []byte(b))
	return hex.EncodeToString(out)
}

func EncryptSHA1(key string) string {
	md := Encrypt(key)
	var sha = sha1.New()
	sha.Write([]byte(md))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	return encryptedString
}
