package main

import (
	"encoding/json"
	"errors"
	"gopkg.in/dedis/crypto.v0/abstract"
	"gopkg.in/dedis/onet.v1/network"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"
	"strconv"
)

//const MaxHomomorphicInt int64 = 100000
var PointToInt map[string]int64 // = make(map[string]int64, MaxHomomorphicInt)
var suite = network.Suite

//const Nmappings = 10000

func printUsage() {
	println("Usage: generateMappingTable <path to mapping.go> <nb mappings to generate>")
}

func parseArguments(args cli.Args) (mappingPath string, nbMappings int64, err error) {

	if len(args) != 3 {
		printUsage()
		err = errors.New("wrong number of arguments")
		return
	}
	mappingPath = args[0]
	nbMappings, errConv := strconv.ParseInt(args[1], 10, 64)
	if errConv != nil {
		printUsage()
		err = errConv
		return
	}
	return
}

// Run this main to generate and populate the mapping table "point -> integer"
func main() {

	// parse arguments
	mappingPath, nbMappings, err := parseArguments(os.Args)
	if err != nil {
		println(err)
		return
	}

	// populate the .go file
	var Bi abstract.Point
	B := suite.Point().Base()
	var m int64

	for Bi, m = suite.Point().Null(), 0; m < nbMappings; Bi, m = Bi.Add(Bi, B), m+1 {
		PointToInt[Bi.String()] = m
	}
	marsh, err := json.Marshal(PointToInt)
	if err != nil {
		println(err.Error())
		return
	}

	// open file and write
	absPath, err := filepath.Abs(mappingPath)
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
		"package main \n" +
			"var PointToInt = map[interface{}]interface{}")
	if err != nil {
		println(err.Error())
		return
	}
	_, err = f.Write(marsh)
	if err != nil {
		println(err.Error())
		return
	}
}
