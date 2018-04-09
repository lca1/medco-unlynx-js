package main

import (
	"gopkg.in/dedis/crypto.v0/abstract"
	"gopkg.in/dedis/onet.v1/network"
	//"encoding/json"
	//"io/ioutil"
	//"github.com/lca1/unlynx/lib"
	"encoding/json"
	"os"
	"path/filepath"
)

// Run this main to generate and populate mappingTable/mapping.go

const MaxHomomorphicInt int64 = 100000
var PointToInt = make(map[string]int64, MaxHomomorphicInt)
var suite = network.Suite
const Nmappings = 10000

func main() {
	// populate the .go file
	var Bi abstract.Point
	B := suite.Point().Base()
	var m int64

	for Bi, m = suite.Point().Null(), 0; m < Nmappings; Bi, m = Bi.Add(Bi, B), m+1 {
		PointToInt[Bi.String()] = m
	}
	marsh, err := json.Marshal(PointToInt)
	if err != nil {
		println(err.Error())
		return
	}

	// open file and write
	absPath, err := filepath.Abs("./tools/gopherjsCrypto/mappingTable/mapping.go")
	if err != nil {
		println(err.Error())
		return
	}
	f, err := os.OpenFile(absPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		println(err.Error())
		return
	}
	defer f.Close()

	// Write go code to initialize the maximum point and integer and the mapping table
	_, err = f.WriteString(
		"package mappingTable \n" +
			//"var greatestM = " +
			//"var greatestInt = " +
			"var PointToInt = map[string]int64")
	if err != nil {
		println(err.Error())
		return
	}
	_, err = f.Write(marsh)
	if err != nil {
		println(err.Error())
		return
	}

	//_, err = f.WriteString(
	//	"func initUnlynx(){" +
	//			"lib.PointToInt=PointToInt;" +
	//			"CurrentGreatestM=" +
	//			"CurrentGreatestInt=" +
	//		"}")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
}
