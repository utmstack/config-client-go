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
const KeyLength = 32 // 32bytes = 256bits = AES256

// Encrypt some data using AES-256 with PKCS#7 Padding using an AES key
// derived via PBKDF2 with HMAC-SHA1
func Encrypt(input []byte, passphrase string) (output string, encryptError error) {

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

	output = base64.StdEncoding.EncodeToString(salt) + base64.StdEncoding.EncodeToString(iv) + base64.StdEncoding.EncodeToString(padded)
	return

}

// Decrypt some data using AES-256 with PKCS#7 Padding using an AES key
// derived via PBKDF2 with HMAC-SHA1
func Decrypt(input, passphrase string) (output []byte, decryptError error) {

	if len(passphrase) < 1 {
		decryptError = &InvalidPassphraseError{"Passphrase length must be greater than zero"}
		return
	}

	salt := []byte(passphrase)
	h := sha1.New()
	h.Write(salt)
	salt = h.Sum(nil)
	if len(salt) != SaltLength {
		decryptError = &InvalidSaltError{"Could not derive salt from input"}
		return
	}

	inputBytes, err := base64.StdEncoding.DecodeString(items[2])
	if err != nil {
		decryptError = &InvalidEncryptedDataError{"Could not derive cipher text from input"}
		return
	}

	// CBC mode always works in whole blocks.
	if len(inputBytes)%aes.BlockSize != 0 {
		decryptError = &InvalidEncryptedDataError{"Cipher text is not a multiple of AES block size"}
		return
	}

	key := pbkdf2.Key([]byte(passphrase), salt, IterationCount, KeyLength, sha1.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		decryptError = &InvalidAESKeyError{err.Error()}
		return
	}

	// Note - IV length is always 16 bytes for AES, regardless of key size
	iv := make([]byte, block.BlockSize())
	if len(iv) != block.BlockSize() {
		decryptError = &InvalidIVError{"Could not derive IV from input"}
		return
	}

	cbc := cipher.NewCBCDecrypter(block, iv)
	cipherText := make([]byte, len(inputBytes))
	copy(cipherText, inputBytes)
	cbc.CryptBlocks(cipherText, cipherText)
	output = PKCS5Padding(cipherText, block.BlockSize())
	return

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
