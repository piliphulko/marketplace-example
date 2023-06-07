package argon2

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"

	"golang.org/x/crypto/argon2"
)

type paramsArgon2 struct {
	Salt        []byte
	PassEncoded []byte
	Time        uint32
	Memory      uint32
	Threads     uint8
	KeyLen      uint32
}

type ParamsArgon2 interface {
	Bytes() ([]byte, error)
	CheckPass([]byte) bool
}

// Bytes returns a slice of bytes that can be written to the database for future authentication
func (pa paramsArgon2) Bytes() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	// Записываем структуру в буфер
	if err := enc.Encode(pa); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// CheckPass checks if the entry password matches
func (pa paramsArgon2) CheckPass(b []byte) bool {
	return bytes.Equal(pa.PassEncoded, argon2.IDKey(b, pa.Salt, pa.Time, pa.Memory, pa.Threads, pa.KeyLen))
}

// GetParamsArgon2 gets interface ParamsArgon2 from a previously recorded record using the Bytes() method
func GetParamsArgon2(b []byte) (ParamsArgon2, error) {
	var (
		bufReader = bytes.NewReader(b)
		pa        paramsArgon2
		dec       = gob.NewDecoder(bufReader)
		err       = dec.Decode(&pa)
	)
	if err != nil {
		return nil, err
	}
	return pa, nil
}

// CreareArgon2Record creates a new entry from the password for authentication
func CreareArgon2Record(pass []byte) (ParamsArgon2, error) {
	var (
		pa paramsArgon2
	)
	pa.Salt = make([]byte, 16)
	pa.Time = 3
	pa.Memory = 64 * 1024 * 2
	pa.Threads = 4
	pa.KeyLen = 32
	if _, err := rand.Read(pa.Salt); err != nil {
		return nil, err
	}
	pa.PassEncoded = argon2.IDKey(pass, pa.Salt, pa.Time, pa.Memory, pa.Threads, pa.KeyLen)

	return pa, nil
}
