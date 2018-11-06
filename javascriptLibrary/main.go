package main

import (
	"encoding/base64"
	"github.com/BurntSushi/toml"
	"github.com/dedis/kyber"
	"github.com/dedis/kyber/suites"
	"github.com/dedis/kyber/util/encoding"
	"github.com/gopherjs/gopherjs/js"
	"github.com/lca1/unlynx/lib"
	"strconv"
	"strings"
)

func main() {

	// declare functions to transpile
	js.Global.Set("AggregateKeys", AggregateKeys)
	js.Global.Set("GenerateKeyPair", GenerateKeyPair)
	js.Global.Set("EncryptInt", EncryptInt)
	js.Global.Set("DecryptInt", DecryptInt)
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
func GenerateKeyPair() (seckey string, pubkey string) {
	sk, pk := libunlynx.GenKey()
	seckey, _ = libunlynx.SerializeScalar(sk)
	pubkey, _ = libunlynx.SerializePoint(pk)
	return
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
