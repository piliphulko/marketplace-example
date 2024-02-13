package jwt

import (
	"testing"
	"time"

	"github.com/piliphulko/marketplace-example/api/apierror"
	"github.com/stretchr/testify/require"
)

func TestTotal(t *testing.T) {
	InsertSecretForSignJWS("12345678qwertyui")

	jws, err := CreateJWS(Header{Alg: "SHA256", Typ: "JWT"}, Payload{Nickname: "test", Exp: time.Now().Add(1 * time.Hour).Unix()})
	require.Nil(t, err)

	jwt, err := jws.SignJWS()
	require.Nil(t, err)

	err = jwt.CheckJWT()
	require.Nil(t, err)

	login, err := jwt.TakeNickname()
	require.Nil(t, err)
	require.Equal(t, "test", login)

	jwt.body = []byte("eyJhbGciOiJBRVMiLCJ0eXAiOiJKV1Qi.fQp7Im5pY2tuYW1lIjoidGVzdCIsImV4cCI6IjI0MjIi.QQ8-1MOWCthrrOaMsF_S--yKKf-FYTDvx6RDGn")
	err = jwt.CheckJWT()
	require.ErrorIs(t, err, apierror.ErrTokenFake)

	jws, err = CreateJWS(Header{Alg: "SHA256", Typ: "JWT"}, Payload{Nickname: "test", Exp: time.Now().Unix() - 1})
	require.Nil(t, err)

	jwt, err = jws.SignJWS()
	require.Nil(t, err)

	err = jwt.CheckJWT()
	require.ErrorIs(t, err, apierror.ErrTokenExpired)
}

var bendE *JWS

func BenchmarkJWS(b *testing.B) {
	InsertSecretForSignJWS("12345678qwertyui")
	var jws JWS
	var err error
	for i := 0; i != 10000; i++ {
		jws, err = CreateJWS(Header{Alg: "SHA256", Typ: "JWT"}, Payload{Nickname: "test", Exp: time.Now().Add(time.Minute).Unix()})
		require.Nil(b, err)
		jwt, err := jws.SignJWS()
		require.Nil(b, err)
		require.Nil(b, jwt.CheckJWT())
	}
	bendE = &jws
}

/*
+ CreateJWS
BenchmarkJWS-8   	1000000000	         0.02022 ns/op	       0 B/op	       0 allocs/op
+ SignJWS
BenchmarkJWS-8   	1000000000	         0.04697 ns/op	       0 B/op	       0 allocs/op
+ CheckJWT
BenchmarkJWS-8   	1000000000	         0.08082 ns/op	       0 B/op	       0 allocs/op
*/
