package main

import (
	"github.com/gopherjs/gopherjs/js"
	"./mylib"
	//"./mappingTable"

	//"io/ioutil"
	//"os"
)

// USE:
// unlynx: branch newI2B2
// onet: branch master

func main() {
	 //--> transpile the functions: gopherjs build -m crypto_javascript.go -o cryptolib.js
	transpileFunctions()


	// --> compute the aggregate key of the cothority
	// given file path
	//rosterFilePath := "src/main/tools/gopherjsCrypto/group.toml"
	//println(mylib.AggregateKeysFromFile(rosterFilePath))

	//--> check how big is the mapping table
	//print(len(mappingTable.PointToInt))
}

func transpileFunctions(){
	js.Global.Set("AggKeys", mylib.AggregateKeys)
	//js.Global.Set("AggKeysFromFile", mylib.AggregateKeysFromFile) // not working, in javascript you can read a file only after a GET request

	js.Global.Set("GenKey", mylib.GenKey)
	js.Global.Set("EncryptStr", mylib.EncryptStr)
	js.Global.Set("DecryptStr", mylib.DecryptStr)
	// also the light encryption
	js.Global.Set("LightEncryptStr_init", mylib.LightEncryptStr_init)
	js.Global.Set("LightEncryptStr", mylib.LightEncryptStr)
}
