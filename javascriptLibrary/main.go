// +build js

package main

import (
	"encoding/base64"
	"github.com/BurntSushi/toml"
	"github.com/lca1/unlynx/lib"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/suites"
	"go.dedis.ch/kyber/v3/util/encoding"
	"strconv"
	"strings"
	"syscall/js"
)

func main() {

	c := make(chan struct{}, 0)

	js.Global().Set("AggregateKeys", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return AggregateKeys(args[0].String())
	}))
	js.Global().Set("GenerateKeyPair", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return GenerateKeyPair()
	}))
	js.Global().Set("EncryptInt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return EncryptInt(args[0].String(), args[1].String())
	}))
	js.Global().Set("DecryptInt", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return DecryptInt(args[0].String(), args[1].String())
	}))

	println("MedCo-Unlynx Javascript WASM-Go crypto engine initialized")
	<-c
}

var isTablePopulated = false

// ----------------------------- functions exported to javascript -----------------------------

// aggregate public keys of nodes into one public key
func AggregateKeys(rosterToml string) string {
	// convert input string to structure GroupToml
	group := &GroupToml{}
	_, err := toml.Decode(rosterToml, group)
	if err != nil {
		println(err)
		return ""
	}

	if len(group.Servers) <= 0 {
		println("Empty or invalid group file", err)
		return ""
	}

	// convert all strings representing the public key to kyber.Point and sum them up
	var agg kyber.Point
	for i, s := range group.Servers {
		// Backwards compatibility with old group files.
		if s.Suite == "" {
			s.Suite = "Ed25519"
		}

		suite, err := suites.Find(s.Suite)
		if err != nil {
			println(err)
			return ""
		}

		pubR := strings.NewReader(s.Public)
		public, err := encoding.ReadHexPoint(suite, pubR)
		if err != nil {
			println(err)
			return ""
		}

		if i == 0 {
			agg = public
		} else {
			if public != nil {
				agg = agg.Add(agg, public)
			}
		}
	}
	b, err := agg.MarshalBinary()
	return base64.URLEncoding.EncodeToString(b)
}

// generate a random pair of keys
func GenerateKeyPair() string {
	sk, pk := libunlynx.GenKey()
	seckey, _ := libunlynx.SerializeScalar(sk)
	pubkey, _ := libunlynx.SerializePoint(pk)
	return seckey + "," + pubkey
}

// encrypt an integer
func EncryptInt(pubkey string, plain string) string {
	m, _ := strconv.ParseInt(plain, 10, 64)
	pointPubkey, _ := libunlynx.DeserializePoint(pubkey)

	return libunlynx.EncryptInt(pointPubkey, m).Serialize()
}

// decrypt an integer
func DecryptInt(ciphertext string, seckey string) string {

	// populate the table with the one created (if it gives an error is just because the mapping table is big)
	if !isTablePopulated {
		libunlynx.PointToInt.PutAll(PointToInt)
		isTablePopulated = true
	}

	scalarSeckey, _ := libunlynx.DeserializeScalar(seckey)
	cipherCipherText := libunlynx.CipherText{}
	cipherCipherText.Deserialize(ciphertext)

	return strconv.FormatInt(libunlynx.DecryptInt(scalarSeckey, cipherCipherText), 10)
}
