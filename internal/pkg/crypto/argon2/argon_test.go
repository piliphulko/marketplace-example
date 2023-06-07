package argon2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTotal(t *testing.T) {
	ap, err := CreareArgon2Record([]byte("asdfghjk"))
	require.Nil(t, err)
	b, err := ap.Bytes()
	require.Nil(t, err)
	apNew, err := GetParamsArgon2(b)
	require.Nil(t, err)
	require.True(t, apNew.CheckPass([]byte("asdfghjk")))
	require.False(t, apNew.CheckPass([]byte("asdfghjh")))
}
