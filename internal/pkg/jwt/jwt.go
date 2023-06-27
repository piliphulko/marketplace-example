package jwt

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"regexp"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var (
	JSON = jsoniter.ConfigCompatibleWithStandardLibrary
	//blockAes  crypto.GeneralCipherBlock
	keySecret []byte
)

var (
	ErrTokenFake    = errors.New("Token fake")
	ErrTokenExpired = errors.New("Token expired")
)

// InsertSecretForSignJWT inserts the key to sign the JWS
func InsertSecretForSignJWS(secret string) {
	keySecret = []byte(secret)
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Nickname string `json:"nickname"`
	Exp      int64  `json:"exp"`
}

type JWS []byte

type JWT struct {
	body []byte
}

func (jwt JWT) String() string {
	return string(jwt.body)
}

// CreateJWS creates a JWS to be signed
func CreateJWS(h Header, p Payload) (JWS, error) {

	hJson, err := JSON.Marshal(h)
	if err != nil {
		return nil, err
	}
	pJson, err := JSON.Marshal(p)
	if err != nil {
		return nil, err
	}
	hJsonBase64 := make([]byte, base64.RawURLEncoding.EncodedLen(len(hJson)))
	pJsonBase64 := make([]byte, base64.RawURLEncoding.EncodedLen(len(pJson)))
	base64.RawURLEncoding.Encode(hJsonBase64, hJson)
	base64.RawURLEncoding.Encode(pJsonBase64, pJson)
	return append(append(hJsonBase64, byte(46)), pJsonBase64...), nil
}

// SignJWS signs a JWS creating a JWT
func (jws JWS) SignJWS() (JWT, error) {
	var (
		bufJWT = bytes.Buffer{}
		mac    = hmac.New(sha256.New, keySecret)
	)

	if _, err := mac.Write(jws); err != nil {
		return JWT{nil}, err
	}
	if _, err := bufJWT.Write(jws); err != nil {
		return JWT{nil}, err
	}
	if err := bufJWT.WriteByte(46); err != nil { // Write "."
		return JWT{nil}, err
	}
	Signature := mac.Sum(nil)
	SignatureBase64URL := make([]byte, base64.RawURLEncoding.EncodedLen(len(Signature)))
	base64.RawURLEncoding.Encode(SignatureBase64URL, Signature)

	if _, err := bufJWT.Write(SignatureBase64URL); err != nil {
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

// CheckJWT checks the integrity and authenticity of the token, as well as the expiration of the token
// possible errors: ErrTokenFake, ErrTokenExpired, others standart errors
func (jwt JWT) CheckJWT() error {
	var (
		mac                        = hmac.New(sha256.New, keySecret)
		header, payload, signature = jwt.splitHeaderPayloadSignature()
		headerPayload              = append(append(header, byte(46)), payload...)

		payloadJson = make([]byte, base64.RawURLEncoding.DecodedLen(len(payload)))
		payloadType Payload
	)

	if _, err := mac.Write(headerPayload); err != nil {
		return err
	}
	signatureHMAC := mac.Sum(nil)
	signatureHMACBase64URL := make([]byte, base64.RawURLEncoding.EncodedLen(len(signatureHMAC)))
	base64.RawURLEncoding.Encode(signatureHMACBase64URL, signatureHMAC)

	if !hmac.Equal(signatureHMACBase64URL, signature) {
		return ErrTokenFake
	}
	base64.RawURLEncoding.Decode(payloadJson, payload)
	if err := JSON.Unmarshal(payloadJson, &payloadType); err != nil {
		return err
	}
	if time.Now().Unix() > payloadType.Exp {
		return ErrTokenExpired
	}

	return nil
}

// TakeNickname returns nickname from token payload
func (jwt JWT) TakeNickname() (string, error) {
	var (
		_, payload, _ = jwt.splitHeaderPayloadSignature()
		payloadJson   = make([]byte, base64.RawURLEncoding.DecodedLen(len(payload)))
		payloadType   Payload
	)
	base64.RawURLEncoding.Decode(payloadJson, payload)
	if err := JSON.Unmarshal(payloadJson, &payloadType); err != nil {
		return "", err
	}
	return payloadType.Nickname, nil
}

// BeIntoJWT converts to jwt
func BeIntoJWT(jwtString string) (JWT, error) {
	re := `^[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+\.[a-zA-Z0-9_-]+$`
	ok, err := regexp.MatchString(re, jwtString)
	if err != nil {
		return JWT{}, err
	}
	if ok {
		return JWT{body: []byte(jwtString)}, nil
	}
	return JWT{}, ErrTokenFake
}

/*
	func CreateJWS(h Header, p Payload) (JWS, error) {
		var (
			bufJWS           = bytes.Buffer{}
			encoderBase64URL = base64.NewEncoder(base64.StdEncoding, &bufJWS)
			jsonEncoder      = jsoniter.NewEncoder(encoderBase64URL)
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
*/
