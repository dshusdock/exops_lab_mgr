package token

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
   
   var secretKey = []byte("secret-key")
   
   func CreateToken(username string) (string, error) {
	   	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
			jwt.MapClaims{ 
		    "username": username, 
		    "exp": time.Now().Add(time.Hour * 24).Unix(), 
		})
   
	   	tokenString, err := token.SignedString(secretKey)
	   	if err != nil {
	   		return "", err
	   	}
		fmt.Printf("Token claims added: %+v\n", token)
   
		return tokenString, nil
   }

   func VerifyToken(tokenString string) error {
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
		})
	
		if err != nil {
		return err
		}
	
		if !token.Valid {
		return fmt.Errorf("invalid token")
		}
	
		return nil
 	}

	// From the Go standard library documentation
	// https://pkg.go.dev/crypto/cipher@go1.23.0#pkg-functions

	func TestEncrypt() {
		// Load your secret key from a safe place and reuse it across multiple
		// Seal/Open calls. (Obviously don't use this example key for anything
		// real.) If you want to convert a passphrase to a key, use a suitable
		// package like bcrypt or scrypt.
		// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
		key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
		plaintext := []byte("exampleplaintext")

		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err.Error())
		}

		// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
		nonce := make([]byte, 12)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			panic(err.Error())
		}

		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}

		ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
		fmt.Printf("%x\n", ciphertext)
	}	

	func TestDecrypt() {
		// Load your secret key from a safe place and reuse it across multiple
		// Seal/Open calls. (Obviously don't use this example key for anything
		// real.) If you want to convert a passphrase to a key, use a suitable
		// package like bcrypt or scrypt.
		// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
		key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
		ciphertext, _ := hex.DecodeString("c3aaa29f002ca75870806e44086700f62ce4d43e902b3888e23ceff797a7a471")
		nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err.Error())
		}

		aesgcm, err := cipher.NewGCM(block)
		if err != nil {
			panic(err.Error())
		}

		plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%s\n", plaintext)
	}
		
