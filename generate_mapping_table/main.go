package main

import (
	"github.com/urfave/cli"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/suites"
	"log"
	"os"
	"strconv"
)

var PointToInt = make(map[string]int64,0)
var suite = suites.MustFind("Ed25519")

//"mapping.ts"

func printUsage() {
	println("Usage: main <path to mapping.ts> <nbr mappings to generate> <negative values: [0,1]>")
}

func parseArguments(args cli.Args) (string, int64, bool) {
	if len(args) != 4 {
		printUsage()
		log.Fatal("Wrong number of arguments")
	}
	mappingPath := args[1]
	nbrMappings, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		printUsage()
		log.Fatal("Error converting <nbr mappings to generate>")
	}

	tmp, err := strconv.ParseInt(args[3], 10, 64)
	if err != nil {
		printUsage()
		log.Fatal("Error converting <negative values: [0,1]>")
	}

	var checkNeg bool
	if tmp == 0 {
		checkNeg = false
	} else if tmp == 1 {
		checkNeg = true
	} else {
		printUsage()
		log.Fatal("<negative values> should be 0 or 1")
	}
	return mappingPath, nbrMappings, checkNeg
}

func writeMapToJSFile(path string, pointToInt map[string]int64) error {
	/*export let PointToInt: Record<string, number> = {
		"edc876d6831fd2105d0b4389ca2e283166469289146e2ce06faefe98b22548df": 5,
		"f47e49f9d07ad2c1606b4d94067c41f9777d4ffda709b71da1d88628fce34d85": 6
	}*/

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("export let PointToInt: Record<string, number> = {\n")
	if err != nil {
		return err
	}
	for k, v := range pointToInt {
		_, err = f.WriteString("\t" + `"` + k + `": ` + strconv.FormatInt(v, 10) + ",\n")
		if err != nil {
			return err
		}
	}
	_, err = f.WriteString("};")
	if err != nil {
		return err
	}
	return nil
}

// Run this main to generate and populate the mapping table "point -> integer"
func main() {
	// parse arguments
	mappingPath, nbrMappings, checkNeg := parseArguments(os.Args)

	// populate the .js file
	var Bi kyber.Point
	B := suite.Point().Base()
	var m int64

	for Bi, m = suite.Point().Null(), 0; m < nbrMappings; Bi, m = Bi.Add(Bi, B), m+1 {
		PointToInt[Bi.String()] = m
		if checkNeg {
			neg := suite.Point().Mul(suite.Scalar().SetInt64(int64(-m)), B)
			PointToInt[neg.String()] = -m
		}
	}
	err := writeMapToJSFile(mappingPath, PointToInt)
	if err != nil {
		log.Fatal(err)
	}
}