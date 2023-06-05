package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"os"
)

func GenerateECDSAKeysAndSign(text []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey, *big.Int, *big.Int, error) {
	// Генеруємо приватний ключ
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Отримуємо публічний ключ з приватного
	publicKey := &privateKey.PublicKey

	// Підписуємо текст за допомогою приватного ключа
	hash := sha256.Sum256(text)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return privateKey, publicKey, r, s, nil
}

func SavePublicKeyToFile(publicKey *ecdsa.PublicKey, filename string) error {
	// Створюємо файл для запису
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Створюємо PEM-блок для серіалізації публічного ключа
	derBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}

	// Записуємо PEM-блок у файл
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	return nil
}

func SavePrivateKeyToFile(privateKey *ecdsa.PrivateKey, filename string) error {
	// Створюємо файл для запису
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Серіалізуємо приватний ключ
	derBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return err
	}

	// Створюємо PEM-блок для серіалізації приватного ключа
	block := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: derBytes,
	}

	// Записуємо PEM-блок у файл
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	return nil
}
