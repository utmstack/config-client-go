package utm_crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"golang.org/x/crypto/pbkdf2"
)

const IterationCount = 65536
const SaltLength = 16
const KeyLength = 16 // 32bytes = 256bits = AES256

// Encrypt some data using AES-256 with PKCS#7 Padding using an AES key
// derived via PBKDF2 with HMAC-SHA1

func Encrypt(text, passphrase string) (output string, encryptError error) {

	input := []byte(text)
	if len(passphrase) < 1 {
		encryptError = &InvalidPassphraseError{"Passphrase length must be greater than zero"}
		return
	}

	salt := make([]byte, SaltLength)
	readSalt, err := rand.Read(salt)
	if err != nil || readSalt != SaltLength {
		encryptError = &CouldNotObtainRandomSaltError{err.Error()}
	}

	key := pbkdf2.Key([]byte(passphrase), salt, IterationCount, KeyLength, sha1.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		encryptError = &InvalidAESKeyError{err.Error()}
		return
	}

	iv := make([]byte, block.BlockSize())
	readIV, err := rand.Read(iv)
	if err != nil || readIV != block.BlockSize() {
		encryptError = &CouldNotObtainRandomIVError{err.Error()}
	}

	padded := PKCS5Padding(input, block.BlockSize())

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(padded, padded)

	output = base64.StdEncoding.EncodeToString(salt) + "#" + base64.StdEncoding.EncodeToString(iv) + "#" + base64.StdEncoding.EncodeToString(padded)
	return

}

// Decrypt some data using AES-256 with PKCS#7 Padding using an AES key
// derived via PBKDF2 with HMAC-SHA1
func Decrypt(input, passphrase string) ([]byte, error) {

	if len(passphrase) < 1 {
		decryptError := &InvalidPassphraseError{"Passphrase length must be greater than zero"}
		return nil, decryptError
	}

	inputBytes := []byte(input)
	salt := []byte(passphrase)
	h := sha1.New()
	h.Write(salt)
	salt = h.Sum(nil)

	// CBC mode always works in whole blocks.
	//if len(inputBytes)%aes.BlockSize != 0 {
	//	decryptError := &InvalidEncryptedDataError{"Cipher text is not a multiple of AES block size"}
	//	return nil, decryptError
	//}

	key := pbkdf2.Key([]byte(passphrase), salt, IterationCount, KeyLength, sha1.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		decryptError := &InvalidAESKeyError{err.Error()}
		return nil, decryptError
	}

	// Note - IV length is always 16 bytes for AES, regardless of key size
	iv := salt[:16]
	if err != nil || len(iv) != block.BlockSize() {
		decryptError := &InvalidIVError{"Could not derive IV from input"}
		return nil, decryptError
	}

	cbc := cipher.NewCBCDecrypter(block, iv)
	cipherText := make([]byte, len(inputBytes))
	copy(cipherText, inputBytes)
	cbc.CryptBlocks(cipherText, PKCS5Padding(cipherText, block.BlockSize()))

	return cipherText, nil

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
