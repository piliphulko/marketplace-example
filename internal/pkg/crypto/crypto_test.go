package crypto

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var testBlock GeneralCipherBlock

func TestMain(m *testing.M) {
	var err error
	testBlock, err = GetBlockCipherAes("12345678qwertyui")
	if err != nil {
		log.Fatal(err)
	}

	testRun := m.Run()
	os.Exit(testRun)
}

func TestAesCFB(t *testing.T) {

	plaintext := []byte("test some text")

	encodeText, err := testBlock.AesEncodeBytesCFB(plaintext)
	require.Nil(t, err)

	decodeText, err := testBlock.AesDecodeBytesCFB(encodeText)
	require.Nil(t, err)

	require.NotEqual(t, encodeText, decodeText)
	require.Equal(t, plaintext, decodeText)
}

func TestAesGCM(t *testing.T) {

	plaintext := []byte("test some text")

	encodeText, err := testBlock.AesEncodeBytesGCM(plaintext)
	require.Nil(t, err)

	decodeText, err := testBlock.AesDecodeBytesGCM(encodeText)
	require.Nil(t, err)

	require.NotEqual(t, encodeText, decodeText)
	require.Equal(t, plaintext, decodeText)
}
