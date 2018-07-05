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
	//// given file path
	//rosterFilePath := "src/main/tools/gopherjsCrypto/group.toml"
	//println(mylib.AggregateKeysFromFile(rosterFilePath))
	//
	////// given file content
	//rosterFileContent := "[[servers]]\n Address = \"tls://10.90.38.8:2000\"\n Suite = \"Ed25519\"\n Public = \"d4ca39db7834fdad06ef8de54e34b4a0942816efe801ed8c1607d197e0d0bb4f\"\n Description = \"Unlynx Server 0\"\n[[servers]]\n Address = \"tls://10.90.38.10:2000\"\n Suite = \"Ed25519\"\n Public = \"cfa45916a96c14b4b9a8417c6ffff4108d73bc048190d0c1c350f955a8e516d7\"\n Description = \"Unlynx Server 1\"\n [[servers]]\n Address = \"tls://10.90.38.11:2000\"\n Suite = \"Ed25519\"\n Public = \"2580a4dc353b979410896d6d71f80b9254ee6999be8361bd2f0c956cf88ea113\"\n Description = \"Unlynx Server 2\""
	//println(mylib.AggregateKeys(rosterFileContent))

	//--> check how big is the mapping table
	//js.Global.Set("MyPrint", myprint)
	//print(len(mappingTable.PointToInt))


	//--> encrypt some variants
	//encryptAndTest()
}

func transpileFunctions(){
	js.Global.Set("AggKeys", mylib.AggregateKeys)
	js.Global.Set("AggKeysFromFile", mylib.AggregateKeysFromFile)

	js.Global.Set("GenKey", mylib.GenKey)
	js.Global.Set("EncryptStr", mylib.EncryptStr)
	js.Global.Set("DecryptStr", mylib.DecryptStr)
	// also the light encryption
	js.Global.Set("LightEncryptStr_init", mylib.LightEncryptStr_init)
	js.Global.Set("LightEncryptStr", mylib.LightEncryptStr)
}


//func encryptAndTest(){
//	//pk :="Sdmuk6sblX10oI47OACV/7YJ5EV9xjC4yv3PxaKc+w4=";
//	//sk :="zjvEJSm8BHpVqqS7PPy4xSE6n7QIorre1yuhuZ8bzg4=";
//	K := "4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyA=";
//	S := "25HhZV/HKHaGn+nfJS0hl5VGys3uCjMZtex6THA5/0c=";
//
//	plains_str := []string{"-7054948997223410688","-7054948998062267136","-7054968999892742144","-7054948999337337856","-7054948997064022784","-7054953138544961536","-7054948997064020736","-7054923607457132544","-7054904773018905600","-7054898625779855360","-7054948987408734208","-7054923379857424384","-7054917142457610240","-7054861692282335232","-7054948050082458624","-7054922546600210432","-7054949048695910400","-7054923381098933248","-7054861517262417920","-7054918645662608384","-7054905185268658176","-7054904954414166016","-7054898626853597184","-7054904932905773056","-7054861823278837760"};
//
//	ciphers := make([]string, len(plains_str))
//	for i, _ := range plains_str{
//		//plain, _ := strconv.ParseInt(pl, 10, 64)
//		ciphers[i] = mylib.LightEncryptStr(plains_str[i], K, S) //strconv.Atoi(pl)
//	}
//
//	//ciphers_js := []string{"4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAyv5RvMN3DwobTYDhMq6KwE7Jbi5K2DemSR+e+yMCklQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyBrfx43POcVJBeVWbCv80Mx+z4QHgPRYYjOcyYX9JNAzw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAx8ZRWteSDfrNQ9lduvoj0/Wq7XJW1Ct8Wv2IyZb4T0Q==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCrGkFfOH5Rccw5NIPl6cNeh6kAW6AkxFtvJ7ysi2ZOsw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCb/fNT38MS0oT1B5+D4+kRbrSqlaYsk7YLYCQiXdhACg==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyBfKtKjZNbqqKcNzQq5ChYqbqI+7V0sFsJJ/JgRNS9fHw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyB+qN/r1RYzFiW9jqN/o3o7D+8U3h8iJQRYKSEXMe6EMA==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDc7zOG56Otw4cPhnjRfZ4WcUg1NguUpI5fwzUUpWlu5Q==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCdOjUH0ApoM8hP2JdrX8qltiyghDDTGSrwDxBOVNVthw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyBbeRwtjCS4Bcc++6QhilACbs9az2JCMkbuOSSbPWKeQg==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAoGWkJCpj6w9YnOeTbPORuDTy+IEFaEztSwCDg1zHGxA==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDqCtcWqfDtSzWSg2TZlmNwoPcIm0YJM1VnAu0pM4JzxQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDUuPl/3dPVEivoycWWNBcbHIklX0QXQTsppLd2zBWvDw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyBDOGVAecu1k33lB5MgcMAv947zzirc9zG+XZI0LF2wPw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCB7O+RPPkWXzB77OUbVsIOAuxtos7xuwLlXhG25GEGnQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCccW3uc6pMdJ8/v2+j2Jji/fFg93ZrK9NMZJnzzLQCug==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyC76GIXmrw/adl5aBhNRsljBzHTU1UlilFpBh3+XQXbxw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDFR+1/JfJYkJm6PjBBBIVFCuppYo5Q/SkXYjk4Zx/Xrg==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAymMpHIGtkR1pUlVFEly5JxV5awErRk+8ZKI2lyoXtlQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCkuDdatdqfD0vng59j3RTG/qJWvQvPOOiRJFOXNcoZiQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyC/WGiMxn3fk/3EWOxaNN0Gq68d6g0/DPoOm3SI34NsTg==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCSJgbnpNrzKpD0nKqrP8P4i7wuQtYcFTzdfQ9ftBA3ug==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyANhOFURQ8tCCy5AFw21WFxzU6GA9WCAHCyQw5JP1wX5Q==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDtVbDm8bPsDbVEDhcO0G92/bKx8K3/otFtdCSLeG/9vQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyC+bqr3iEhbhbaGrJ9ROldzfvZm+zeKwfNtchxVCI2CYQ=="};
//	ciphers_js := []string{"4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAyv5RvMN3DwobTYDhMq6KwE7Jbi5K2DemSR+e+yMCklQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyD9FpF5q0l80A0ds10NjvV5H2u6AFBaZKi4BO+/Z54Vvg==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAx8ZRWteSDfrNQ9lduvoj0/Wq7XJW1Ct8Wv2IyZb4T0Q==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCrGkFfOH5Rccw5NIPl6cNeh6kAW6AkxFtvJ7ysi2ZOsw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyA/jzsdPKJ1FaPEitVyleZ7x2qNx2JiSTyrvc5OIO46Qw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyBfKtKjZNbqqKcNzQq5ChYqbqI+7V0sFsJJ/JgRNS9fHw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAecn+rs24+kvtBxMNlbyFGvFJss5tKFSIL8uHlUrVkog==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDc7zOG56Otw4cPhnjRfZ4WcUg1NguUpI5fwzUUpWlu5Q==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCdOjUH0ApoM8hP2JdrX8qltiyghDDTGSrwDxBOVNVthw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyBbeRwtjCS4Bcc++6QhilACbs9az2JCMkbuOSSbPWKeQg==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAoGWkJCpj6w9YnOeTbPORuDTy+IEFaEztSwCDg1zHGxA==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDqCtcWqfDtSzWSg2TZlmNwoPcIm0YJM1VnAu0pM4JzxQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDUuPl/3dPVEivoycWWNBcbHIklX0QXQTsppLd2zBWvDw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyBDOGVAecu1k33lB5MgcMAv947zzirc9zG+XZI0LF2wPw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCB7O+RPPkWXzB77OUbVsIOAuxtos7xuwLlXhG25GEGnQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCccW3uc6pMdJ8/v2+j2Jji/fFg93ZrK9NMZJnzzLQCug==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyC76GIXmrw/adl5aBhNRsljBzHTU1UlilFpBh3+XQXbxw==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDFR+1/JfJYkJm6PjBBBIVFCuppYo5Q/SkXYjk4Zx/Xrg==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyAymMpHIGtkR1pUlVFEly5JxV5awErRk+8ZKI2lyoXtlQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCkuDdatdqfD0vng59j3RTG/qJWvQvPOOiRJFOXNcoZiQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyC/WGiMxn3fk/3EWOxaNN0Gq68d6g0/DPoOm3SI34NsTg==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyCSJgbnpNrzKpD0nKqrP8P4i7wuQtYcFTzdfQ9ftBA3ug==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyANhOFURQ8tCCy5AFw21WFxzU6GA9WCAHCyQw5JP1wX5Q==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyDtVbDm8bPsDbVEDhcO0G92/bKx8K3/otFtdCSLeG/9vQ==","4zKXNX74MKVSWMZGScdiOBea+kra3W7CdZvVkhX+yyC+bqr3iEhbhbaGrJ9ROldzfvZm+zeKwfNtchxVCI2CYQ=="}
//
//	println(strings.Join(ciphers,","))
//	println(strings.Join(ciphers_js,","), "\n")
//	diff := ciphersDifferent(ciphers, ciphers_js)
//	for i,_ := range diff{
//		println(plains_str[i], "correct enc: ", ciphers[i], "wrong enc: ", ciphers_js[i])
//	}
//
//	//for _, i := range diff{
//	//	println("\n", plains[i], ":")
//	//	for j, _ := range plains{
//	//		if !intInSlice(j, diff){
//	//			println(i, "<", j, plains[i]<plains[j])
//	//		}
//	//	}
//	//	print()
//	//}
//}
//
//func ciphersDifferent(a []string, b []string) []int {
//	if len(a) != len(b) {
//		return []int{-1}
//	}
//
//	diffCiphers := []int{}
//	for i, v := range a {
//
//		if v != b[i] {
//			diffCiphers = append(diffCiphers, i)
//		}
//	}
//	return diffCiphers
//}
//
//func intInSlice(a int, list []int) bool {
//	for _, b := range list {
//		if b == a {
//			return true
//		}
//	}
//	return false
//}
