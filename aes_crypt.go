package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func Encrypt(text []byte) ([]byte, []byte, error) {

	buffer := make([]byte, 16)

	_, err := rand.Read(buffer)
	if err != nil {
		return nil, nil, err
	}
	key := []byte(fmt.Sprintf("%x", buffer))

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Помилка при ініціалізації блочного шифру AES:", err)
		return nil, nil, err
	}

	// Створення рандомного вектору ініціалізації (IV)
	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		fmt.Println("Помилка при генерації вектора ініціалізації (IV):", err)
		return nil, nil, err
	}

	// Шифрування даних
	ciphertext := make([]byte, aes.BlockSize+len(text))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)
	copy(ciphertext[:aes.BlockSize], iv)

	return key, ciphertext, nil
}

func Decrypt(encrypted []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := encrypted[:aes.BlockSize]

	blockMode := cipher.NewCTR(block, iv)

	plaintext := make([]byte, len(encrypted)-aes.BlockSize)
	blockMode.XORKeyStream(plaintext, encrypted[aes.BlockSize:])

	return plaintext, nil
}
