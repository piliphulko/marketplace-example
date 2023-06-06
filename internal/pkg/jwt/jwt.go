package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

var (
	JSON = jsoniter.ConfigCompatibleWithStandardLibrary
	//blockAes  crypto.GeneralCipherBlock
	keySecret []byte
)

var (
	ErrFakeToken = errors.New("fake token")
)

func InsertSecretForSignJWT(secret string) {
	keySecret = []byte(secret)
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Nickname string `json:"nickname"`
	Exp      string `json:"exp"`
}

type JWS []byte

type JWT struct {
	body []byte
}

func (jwt JWT) String() string {
	return string(jwt.body)
}

func CreateJWS(h Header, p Payload) (JWS, error) {
	var (
		bufJWS           = bytes.Buffer{}
		encoderBase64URL = base64.NewEncoder(base64.URLEncoding, &bufJWS)
		jsonEncoder      = JSON.NewEncoder(encoderBase64URL)
	)
	defer encoderBase64URL.Close()

	if err := jsonEncoder.Encode(&h); err != nil {
		return nil, err
	}
	bufJWS.WriteByte(46) // Write "."

	if err := jsonEncoder.Encode(&p); err != nil {
		return nil, err
	}

	return bufJWS.Bytes(), nil
}

func (jws JWS) SignJWS() (JWT, error) {
	var (
		bufJWT           = bytes.Buffer{}
		encoderBase64URL = base64.NewEncoder(base64.URLEncoding, &bufJWT)
	)
	defer encoderBase64URL.Close()

	mac := hmac.New(sha256.New, keySecret)

	if _, err := mac.Write(jws); err != nil {
		return JWT{nil}, err
	}
	/*
		signature, err := blockAes.AesEncodeBytesGCM(jws)
		if err != nil {
			return "", nil
		}
	*/
	if _, err := bufJWT.Write(jws); err != nil {
		return JWT{nil}, err
	}
	if err := bufJWT.WriteByte(46); err != nil { // Write "."
		return JWT{nil}, err
	}
	if _, err := encoderBase64URL.Write(mac.Sum(nil)); err != nil {
		return JWT{nil}, err
	}
	return JWT{bufJWT.Bytes()}, nil
}

func (jwt JWT) splitHeaderPayloadSignature() ([]byte, []byte, []byte) {
	var (
		split                      = bytes.Split(jwt.body, []byte("."))
		header, payload, signature = split[0], split[1], split[2]
	)
	return header, payload, signature
}

func (jwt JWT) CheckJWT() error {
	var (
		mac              = hmac.New(sha256.New, keySecret)
		bufHMAC          = bytes.Buffer{}
		encoderBase64URL = base64.NewEncoder(base64.URLEncoding, &bufHMAC)

		header, payload, signature = jwt.splitHeaderPayloadSignature()
		headerPayload              = append(append(header, byte(46)), payload...)
	)
	defer encoderBase64URL.Close()

	if _, err := mac.Write(headerPayload); err != nil {
		return err
	}

	if _, err := encoderBase64URL.Write(mac.Sum(nil)); err != nil {
		return err
	}

	if !hmac.Equal(bufHMAC.Bytes(), signature) {
		return ErrFakeToken
	}
	return nil
}
