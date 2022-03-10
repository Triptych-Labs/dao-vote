package cryptog

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func Decrypt(key, hexRepr string) *string {
	ciphertext, err := hex.DecodeString(hexRepr)
	if err != nil {
		return nil
	}
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	deciphered, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil
	}
	decipheredText := string(deciphered)

	return &decipheredText
}

