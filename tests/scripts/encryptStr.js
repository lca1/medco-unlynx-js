// function worker_code() {
//     // self.importScripts("http://localhost:63342/i-b-webclient/tools/gopherjsCrypto/tests/scripts/useme.js");
//
// 	//all code here
//     onmessage = function(e) {
// 	var keys = Object.keys(e.data);
// 	if (keys.includes("import")){
// 		var import_path = e.data["import"];
// 		self.importScripts(import_path);
// 		console.log("imported cryptolib")
// 	}
// 	else if (keys.includes("plain") && keys.includes("K") && keys.includes("S")){
// 		// the passed-in data is available via e.data
// 		var i = e.data["plain"];
// 		var K = e.data["K"];
// 		var S = e.data["S"];
//
// 		//console.log(i, K, S)
// 		postMessage([i, LightEncryptStr(i, K, S)]);
// 	}
//     };
// }
//
// // This is in case of normal worker start
// // "window" is not defined in web worker
// // so if you load this file directly using `new Worker`
// // the worker code will still execute properly
// if(window!=self)
//   worker_function();
