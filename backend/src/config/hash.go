package config

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/argon2"
)

type params struct {
    memory      uint32
    iterations  uint32
    parallelism uint8
    saltLength  uint32
    keyLength   uint32
}

func HashPassword(password string) (string, error) {
	p := &params{
        memory:      64 * 1024,
        iterations:  3,
        parallelism: 2,
        saltLength:  16,
        keyLength:   32,
    }

    hash, err := generateFromPassword("password123", p)
    if err != nil {
        log.Fatal(err)
    }

	safeString := base64.StdEncoding.EncodeToString(hash)
	return safeString, nil
}

func generateFromPassword(password string, p *params) (hash []byte, err error) {
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return nil, err
	}

	hash = argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return hash, nil
}

func generateRandomBytes(n uint32)([]byte, error){
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	
	if err != nil{
		return nil, err
	}

	return bytes, nil
}