package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type GeneralCipherBlock interface {
	AesEncodeBytesCFB([]byte) ([]byte, error)
	AesDecodeBytesCFB([]byte) ([]byte, error)
	AesEncodeBytesGCM([]byte) ([]byte, error)
	AesDecodeBytesGCM([]byte) ([]byte, error)
}

type generalCipherBlock struct {
	cipher.Block
}

func GetBlockCipherAes(secret string) (generalCipherBlock, error) {
	blockAes, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return generalCipherBlock{nil}, err
	}
	return generalCipherBlock{blockAes}, nil
}

func (gcb generalCipherBlock) AesEncodeBytesCFB(plaintext []byte) ([]byte, error) {

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(gcb, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func (gcb generalCipherBlock) AesDecodeBytesCFB(ciphertext []byte) ([]byte, error) {

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(gcb, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func (gcb generalCipherBlock) AesEncodeBytesGCM(plaintext []byte) ([]byte, error) {

	aead, err := cipher.NewGCM(gcb)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := make([]byte, 0, len(plaintext)+aead.Overhead()+len(nonce))
	ciphertext = append(ciphertext, nonce...)
	ciphertext = aead.Seal(ciphertext, nonce, plaintext, nil)

	return ciphertext, nil
}

func (gcb generalCipherBlock) AesDecodeBytesGCM(ciphertext []byte) ([]byte, error) {

	aead, err := cipher.NewGCM(gcb)
	if err != nil {
		return nil, err
	}

	nonce := ciphertext[:aead.NonceSize()]
	ciphertextCut := ciphertext[aead.NonceSize() : len(ciphertext)-aead.Overhead()]
	tag := ciphertext[len(ciphertext)-aead.Overhead():]

	plaintext, err := aead.Open(nil, nonce, append(ciphertextCut, tag...), nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
