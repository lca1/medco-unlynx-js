package mylib

import (
	"gopkg.in/dedis/crypto.v0/abstract"
	"github.com/lca1/unlynx/lib"
	networkv1 "gopkg.in/dedis/onet.v1/network"
	"encoding/json"
	"io/ioutil"
	"../mappingTable"
	"gopkg.in/dedis/crypto.v0/random"
	"strconv"

	"encoding/base64"
	"github.com/BurntSushi/toml"
	//"github.com/dedis/onet"
	"github.com/dedis/kyber/suites"
	"strings"
	"github.com/dedis/kyber/util/encoding"
	"github.com/dedis/kyber"
)

var suite = networkv1.Suite

type Keys struct{
	SecKey string // scalar to []byte (as a string it does not work when writing and reading back from file)
	PubKey string // point to string
}

type CipherString struct{
	K, C string // Point to string
}

// AGGREGATE PUBLIC KEYS FROM group.toml
func AggregateKeysFromFile(rosterFilePath string) string{
	b, err := ioutil.ReadFile(rosterFilePath)
	if err != nil {
		println("Error while opening group file", err)
		return ""
	}

	return AggregateKeys(string(b))
}

// copy here the needed classes to avoid importing onet (gopherjs gives errors when importing onet...)

// GroupToml holds the data of the group.toml file.
type GroupToml struct {
	Servers []*ServerToml `toml:"servers"`
}
// ServerToml is one entry in the group.toml file describing one server to use for
// the cothority.
type ServerToml struct {
	Address     string
	Suite       string
	Public      string
	Description string
}
func AggregateKeys(rosterToml string) string{
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

		if i==0 {
			agg = public
		} else {
			if public != nil {
				agg = agg.Add(agg, public)
			}
		}
	}
	b, err := agg.MarshalBinary()
	return base64.StdEncoding.EncodeToString(b)
}


// READ WRITE KEYS
func WriteKeysToFile(secKey abstract.Scalar, pubKey abstract.Point, filename string) error{
	k := Keys{ScalarToString(secKey), PointToString(pubKey)}

	JSONkeys, err := json.Marshal(k)
	if err != nil {
		println("error when marshalling the keys from file: " + err.Error())
		return err
	}

	err = ioutil.WriteFile(filename, JSONkeys, 0644)
	if err != nil {
		println("error when writing the keys to file" + err.Error())
		return err
	}
	return nil
}

func ReadKeysFromFile(filename string) (secKey abstract.Scalar, pubKey abstract.Point, err error){
	JSONkeys := []byte{}
	JSONkeys, err = ioutil.ReadFile(filename)
	if err != nil {
		println("error when reading from file" + err.Error())
		return
	}

	k := Keys{}
	err = json.Unmarshal(JSONkeys, &k)
	if err != nil {
		println("error when unmarshalling the keys" + err.Error())
		return
	}
	secKey = StringToScalar(k.SecKey)
	pubKey = StringToPoint(k.PubKey)
	return
}

// PUBLIC KEY
func PointToString(p abstract.Point) string {
	str, err := lib.SerializePoint(p)
	if err != nil {
		println("error when converting the point to string" + err.Error())
		panic(err)
		return ""
	}
	return str
}

func StringToPoint(str string) abstract.Point{
	if str == "" {
		nullPoint := networkv1.Suite.Point().Null()
		//js.Global.Call("alert", "nil point")
		return nullPoint
	}

	point, err := lib.DeserializePoint(str)
	if err != nil {
		println("error when converting the string to point" + err.Error())
		panic(err)
		return nil
	}
	return point
}


// SECRET KEY
func ScalarToString(scalar abstract.Scalar) string{
	str, err := lib.SerializeScalar(scalar)
	if err != nil {
		println("error when converting the secret key to string")
		return "error when converting the secret key to string"
	}

	return str
}

func StringToScalar(str string) abstract.Scalar{
	secret, err := lib.DeserializeScalar(str)
	if err != nil {
		println("error when decoding the string")
		//js.Global.Call("alert", "error when decoding the string")
		return nil
	}
	return secret
}


// CIPHERTEXT

func CipherToString(c lib.CipherText) string{
	return c.Serialize()
}

func StringToCipher(c string) lib.CipherText{
	cipher := lib.CipherText{}
	err := cipher.Deserialize(c)
	if err != nil {
		println("error when deserializing the ciphertext")
		//js.Global.Call("error when unmarshalling the binary")
		return lib.CipherText{}
	}
	return cipher
}

// CRYPTO FUNCTIONS

// sec key type: string
// pub key type: string

func GenKey() (seckey string, pubkey string){
	sk, pk := lib.GenKey()
	seckey = ScalarToString(sk)
	pubkey = PointToString(pk)
	return
}

func EncryptStr(pubkey string, plain string) string{
	m, _ := strconv.ParseInt(plain, 10, 64)
	c := lib.EncryptInt(StringToPoint(pubkey), m)
	//return CipherToCipherString(*c)
	return CipherToString(*c)
}

func DecryptStr(ciphertext string, seckey string) string{
	// populate the table with the one created (if it gives an error is just because the mapping table is big)
	lib.PointToInt=mappingTable.PointToInt
	//return lib.DecryptInt(StringToScalar(seckey), CipherStringToCipher(ciphertext))
	return strconv.FormatInt(lib.DecryptInt(StringToScalar(seckey), StringToCipher(ciphertext)), 10)
}


	// LIGHT VERSIONS OF THE ENCRYPT FUNCTION (more static)
// ciphertext = (K,C) = (k*B, S + M) = (k*B, pubkey*k + m*B)
// S = pubkey*k (ephemeral DH shared secret)
// M = m*B

// B and pubkey are fixed
// if you fix also k then you also fixed K and S and compute only M and S+M

func lightEncryptStr_init_(pubkey abstract.Point)(K, S abstract.Point){
	// generate a new ephemeral key and compute: (K, S) = (k*B, pubkey*k)
	B := suite.Point().Base()
	k := suite.Scalar().Pick(random.Stream) // ephemeral private key
	K = suite.Point().Mul(B, k)      // ephemeral DH public key
	S = suite.Point().Mul(pubkey, k) // ephemeral DH shared secret
	return
}

func lightEncryptStr_(m int64, K, S abstract.Point)*lib.CipherText{
	// computes ciphertext = (K, S + m*B)
	C := S.Add(S, lib.IntToPoint(m))   // message blinded with secret
	return &lib.CipherText{K, C}
}

// wrappers
func LightEncryptStr_init(pubkey string) (K_str, S_str string){
	K, S := lightEncryptStr_init_(StringToPoint(pubkey))
	K_str, S_str = PointToString(K), PointToString(S)
	return
}

func LightEncryptStr(m string, K, S string) string{
	m_, _ := strconv.ParseInt(m, 10, 64)
	c := lightEncryptStr_(m_, StringToPoint(K), StringToPoint(S))
	return CipherToString(*c)
}

//// crappy one, only strings!!!
//func LightEncryptInt(m int64, K, S string) string{
//	c := LightEncryptStr_(m, StringToPoint(K), StringToPoint(S))
//	return CipherToString(*c)
//}