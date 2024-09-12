package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	// "github.com/go-chi/jwtauth"
	// "github.com/golang-jwt/jwt/v5"
	// "github.com/go-chi/jwtauth/v5"
	// "github.com/lestrrat-go/jwx/v2/jwt"
)
   
var secretKey = []byte("secret-key")

func init() {

}

   


// From the Go standard library documentation
// https://pkg.go.dev/crypto/cipher@go1.23.0#pkg-functions

func EncryptValue(text string) ([]byte, error) {
	// key := []byte("passphrasewhichneedstobe32bytes!")
	key := []byte("keepstrackofthepreviousvalueok!!")
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, []byte(text), nil), nil
}	

func DecryptValue(ciphertext []byte) (string, error) {
	// key := []byte("passphrasewhichneedstobe32bytes!")
	key := []byte("keepstrackofthepreviousvalueok!!")
	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
		
