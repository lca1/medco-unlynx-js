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
	js.Global.Set("DecryptInt", TempOverrideDecryptInt)

	js.Module.Get("exports").Set("AggregateKeys", AggregateKeys)
	js.Module.Get("exports").Set("GenerateKeyPair", GenerateKeyPair)
	js.Module.Get("exports").Set("EncryptInt", EncryptInt)
	js.Module.Get("exports").Set("DecryptInt", TempOverrideDecryptInt)
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
		//libunlynx.PointToInt.PutAll(PointToInt)
		isTablePopulated = true
	}

	scalarSeckey, _ := libunlynx.DeserializeScalar(seckey)
	cipherCipherText := libunlynx.CipherText{}
	cipherCipherText.Deserialize(ciphertext)

	return strconv.FormatInt(libunlynx.DecryptInt(scalarSeckey, cipherCipherText), 10)
}

// todo: waiting on GopherJS bug resolution
// --- 19/11/18: temporarily override decrypt function (because of bug in GopherJS + ConcurrentMap)

const MaxHomomorphicInt int64 = 100000
var PointToIntOverride = make(map[string]int64, MaxHomomorphicInt)
var SuiTe = suites.MustFind("Ed25519")
var currentGreatestInt int64
var currentGreatestM kyber.Point

func TempOverrideDecryptInt(ciphertext string, seckey string) string {

	if !isTablePopulated {
		PointToIntOverride = PointToInt
		isTablePopulated = true
	}

	scalarSeckey, _ := libunlynx.DeserializeScalar(seckey)
	cipherCipherText := libunlynx.CipherText{}
	cipherCipherText.Deserialize(ciphertext)

	// decryptPoint()
	S := SuiTe.Point().Mul(scalarSeckey, cipherCipherText.K)
	M := SuiTe.Point().Sub(cipherCipherText.C, S)

	// discreteLog()
	B := SuiTe.Point().Base()
	var ok bool
	var m int64

	if m, ok = PointToInt[M.String()]; ok {
		return strconv.FormatInt(m, 10)
	}

	//initialise
	if currentGreatestInt == 0 {
		currentGreatestM = SuiTe.Point().Null()
	}
	foundPos := false
	guess := currentGreatestM
	guessInt := currentGreatestInt

	start := true
	for i := guessInt; i < MaxHomomorphicInt && !foundPos; i = i + 1 {
		// check for 0 first
		if !start {
			guess = SuiTe.Point().Add(guess, B)
		}
		start = false

		guessInt = i
		PointToInt[guess.String()] = guessInt
		if guess.Equal(M) {
			foundPos = true
		}
	}
	currentGreatestM = guess
	currentGreatestInt = guessInt

	if !foundPos {
		println("out of bound encryption, bound is ", MaxHomomorphicInt)
		return strconv.FormatInt(0, 10)
	}

	return strconv.FormatInt(guessInt, 10)
}
