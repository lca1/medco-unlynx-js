package mylib_test


import (
	"../mylib"
	"testing"
	"github.com/lca1/unlynx/lib"
	"github.com/stretchr/testify/require"
	//"os"
	"os"
	"strconv"
	"math/rand"
)

const file = "file"
const trials = 500

func TestConvertions(t *testing.T) {
	for i:=0; i < trials; i++{
		sk, pk := lib.GenKey()
		require.True(t, sk.Equal(mylib.StringToScalar(mylib.ScalarToString(sk.Clone()))))
		//require.True(t, sk.Equal(mylib.ByteToScalar(mylib.ScalarToByte(sk.Clone()))))
		require.True(t, pk.Equal(mylib.StringToPoint(mylib.PointToString(pk.Clone()))))
	}
}

func TestConvertionFromFile(t *testing.T){
	defer func() {
		os.Remove(file)
	}()

	var i int64 =0

	for ; i < trials*10; i+=10 {
		secKey1, pubKey1 := lib.GenKey()

		err := mylib.WriteKeysToFile(secKey1, pubKey1, file)
		require.Nil(t, err)

		secKey2, pubKey2, err := mylib.ReadKeysFromFile(file)
		require.Nil(t, err)

		require.True(t, secKey1.Equal(secKey2))
		require.True(t, pubKey1.Equal(pubKey2))

		// check encryption and decryption too, just to be sure
		c := lib.EncryptInt(pubKey1, i)
		plain := lib.DecryptInt(secKey1, *c)
		require.Equal(t, i, plain)
	}
}

func TestCrypto(t *testing.T){
	var i int64 =0

	for ; i < trials; i+=1 {
		secKey, pubKey := mylib.GenKey()
		plain_start:= strconv.FormatInt(rand.Int63n(1000), 10)

		cipher :=  mylib.EncryptStr(pubKey, plain_start)
		plain_end  :=  mylib.DecryptStr(cipher, secKey)
		require.Equal(t, plain_start, plain_end)
	}
}


func TestLightEncrypt(t *testing.T){
	var i int64 = 0

	secKey, pubKey := mylib.GenKey()
	K, S := mylib.LightEncryptStr_init(pubKey)
	for ; i < trials; i+=1 {
		//cipher :=  mylib.LightEncryptInt(i*5, K, S)
		plain_start:= strconv.FormatInt(rand.Int63n(1000), 10)

		cipher :=  mylib.LightEncryptStr(plain_start, K, S)
		plain_end  :=  mylib.DecryptStr(cipher, secKey)
		require.Equal(t, plain_start, plain_end)
	}
}
