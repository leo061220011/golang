package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := aesGCM.Seal(nil, nonce, data, nil)
	return append(nonce, ciphertext...), nil
}

func decrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("тут")
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("данные слишком короткие")
	}
	nonce := data[:nonceSize]
	ciphertext := data[nonceSize:]
	return aesGCM.Open(nil, nonce, ciphertext, nil)

}

func main() {
	//hash sha256
	/*
		data := []byte("hello from Go!")
		hash := sha256.Sum256(data)
		fmt.Printf("Исходный: %s \n Hash: %x", data, hash)
	*/

	//HMAC
	/*
		data := []byte("hello from Go! how are you")
		key := []byte("secret-key")

		h := hmac.New(sha256.New, key)
		h.Write(data)
		result := h.Sum(nil)
		fmt.Printf("%x", result)
	*/
	//AES
	key := []byte("my-secret-keyyyy")
	data := []byte("Hello from Go! Hello from Go! Hello from Go!")

	encryptedData, err := encrypt(data, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x\n", encryptedData)

	decryptedData, err := decrypt(encryptedData, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Исходный: %s", decryptedData)
}
