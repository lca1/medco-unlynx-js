package main

import (
	"github.com/lca1/unlynx/lib"
	"../mylib"
)

// This main creates public and secret keys and store them to file

const KeysFileName = "keys"

func main() {
	//sk, pk := mylib.GenKey()
	//println(sk)
	//println(pk)
	secKey, pubKey := lib.GenKey()

	mylib.WriteKeysToFile(secKey, pubKey, KeysFileName)
	//if !reflect.DeepEqual([]byte(k.Sk), secKey.Bytes()){
	//	byte1 := []byte(k.Sk)
	//	byte2 := secKey.Bytes()
	//	l1 := len(byte1)
	//	l2 := len(byte2)
	//	l := l1
	//	if l2 > l1 { l = l2 }
	//
	//	for i := 0; i < l; i++ {
	//		if i < l1 {print(byte1[i])}
	//		print("\t\t")
	//		if i < l2 {print(byte2[i])}
	//		println()
	//	}
	//	println("fuck")
	//	return
	//}
}
