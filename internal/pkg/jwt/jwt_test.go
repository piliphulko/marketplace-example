package jwt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
func TestMain(m *testing.M) {

		InsertSecretForSignJWT("12345678qwertyui")

		testRun := m.Run()
		os.Exit(testRun)
	}
*/
func TestTotal(t *testing.T) {
	InsertSecretForSignJWT("12345678qwertyui")

	jws, err := CreateJWS(Header{Alg: "AES", Typ: "JWT"}, Payload{Nickname: "test", Exp: "2422"})
	require.Nil(t, err)

	jwt, err := jws.SignJWS()
	require.Nil(t, err)

	err = jwt.CheckJWT()
	require.Nil(t, err)

	jwt.body = []byte("eyJhbGciOiJBRVMiLCJ0eXAiOiJKV1Qi.fQp7Im5pY2tuYW1lIjoidGVzdCIsImV4cCI6IjI0MjIi.QQ8-1MOWCthrrOaMsF_S--yKKf-FYTDvx6RDGn")
	err = jwt.CheckJWT()
	require.ErrorIs(t, err, ErrFakeToken)
}
