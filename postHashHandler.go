package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func readPasswordValueHttp(r *http.Request) string {
	r.ParseForm()
	password := r.Form.Get("password")
	return password
}

func EncryptAES(plaintext string) string {
	// create cipher
	c, err := aes.NewCipher([]byte(os.Getenv("passphrase")))
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	return base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(plaintext), nil))
}

func PostHashHttp(w http.ResponseWriter, r *http.Request) {

	password := readPasswordValueHttp(r)
	if password == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Warning: Can not read password value from request.")
		return
	}
	encodedString := EncryptAES(password)
	hashId := CounterMap.Set(encodedString)
	w.Write([]byte(strconv.Itoa(hashId)))
	return
}
