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
		fmt.Printf("%x\n", nonce)
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

	var value []byte

	func TestEncrypt2() {
		fmt.Println("Encryption Program v0.01")

		text := []byte("My Super Secret")
		key := []byte("passphrasewhichneedstobe32bytes!")
	
		// generate a new aes cipher using our 32 byte long key
		c, err := aes.NewCipher(key)
		// if there are any errors, handle them
		if err != nil {
			fmt.Println(err)
		}
	
		// gcm or Galois/Counter Mode, is a mode of operation
		// for symmetric key cryptographic block ciphers
		// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
		gcm, err := cipher.NewGCM(c)
		// if any error generating new GCM
		// handle them
		if err != nil {
			fmt.Println(err)
		}
	
		// creates a new byte array the size of the nonce
		// which must be passed to Seal
		nonce := make([]byte, gcm.NonceSize())
		// populates our nonce with a cryptographically secure
		// random sequence
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			fmt.Println(err)
		}
	
		// here we encrypt our text using the Seal function
		// Seal encrypts and authenticates plaintext, authenticates the
		// additional data and appends the result to dst, returning the updated
		// slice. The nonce must be NonceSize() bytes long and unique for all
		// time, for a given key.
		// fmt.Println(gcm.Seal(nonce, nonce, text, nil))

		value = gcm.Seal(nonce, nonce, text, nil)
		fmt.Println(value)
		fmt.Println(len(value))

		// the WriteFile method returns an error if unsuccessful
		// err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
		// // handle this error
		// if err != nil {
		// 	// print it out
		// 	fmt.Println(err)
		// }
	}

	func TestDecrypt2() {
		fmt.Println("Decryption Program v0.01")

		key := []byte("passphrasewhichneedstobe32bytes!")
		// ciphertext, err := ioutil.ReadFile("myfile.data")
		ciphertext := value
		// if our program was unable to read the file
		// print out the reason why it can't
		// if err != nil {
		// 	fmt.Println(err)
		// }
	
		c, err := aes.NewCipher(key)
		if err != nil {
			fmt.Println(err)
		}
	
		gcm, err := cipher.NewGCM(c)
		if err != nil {
			fmt.Println(err)
		}
	
		nonceSize := gcm.NonceSize()
		if len(ciphertext) < nonceSize {
			fmt.Println(err)
		}
	
		nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
		plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(plaintext))
	}

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
		
