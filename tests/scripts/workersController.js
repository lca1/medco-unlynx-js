i2b2 = {};
i2b2.MedCo = {};
i2b2.MedCo.ctrlr = {};

// a structure to handle the workers (background threads)
i2b2.MedCo.ctrlr.background = {
    // from outside you should just call the following
    // 1. initWorkers() and initEncrypt() (in whatever order)
    // 2. toBeEncrypted() x N (call this functions whenever you have something to encrypt)
    // 3. done() (wait/check if encryption has finished)

    // here is stored the list of workers of with the workload can be spread
    workers: [],

    // This is a map from plaintexts to the corresponding ciphertext to avoid encrypting again.
    // It is updated every time a worker encrypts new integer.
    encryptionCache: {},

    // A list of plaintexts that have to be encrypted (notice that some of these my have the corresponding ciphertext
    // in the encryptionCache). When a plaintext is encrypted it is saved in encryptionCache and removed from toEncrypt
    toEncrypt: [],

    // store the number of remaining plaintext to encrypt ("toEncrypt.length" may be different from "remaining" since
    // when a plaintext is sent to a worker it is removed from toEncrypt but "remaining" is not decreased yet)
    remaining: 0,

    // K and S are the parameters that will be used to encrypt
    K: null,
    S: null,
    // aggregate public key of the cothority (used to decrypt)
    AggregateKey: "GKrufk6bsuMwegxcPqr7B1aFWKit1szJlugZ01HkSPA=",

    // init initializes the workers (to create them one time and then just send them the strings to
    // encrypt/decrypt). Initialise a worker may take some seconds because each one of them ahas to import
    // the crypto javascript
    initWorkers: function(num_workers){
        var context = i2b2.MedCo.ctrlr.background;

        // first kill the existing workers, if any
        for (var i=0; i < context.workers.length; i++) {
            context.workers[i].terminate();
        }
        context.workers = [];

        // todo be careful the importPath to points to the correct location of the file
        var importPath = document.location["href"];
        // you have to tell the workers where the crypto library is so that they can import it
        importPath = importPath.substr(0, importPath.lastIndexOf('/')) + "/scripts/cryptolib.js";

        for (var i=0; i < num_workers; i++){
            // start a new worker
            var w = new Worker(URL.createObjectURL(new Blob(["("+worker_code.toString()+")()"], {type: 'text/javascript'})));

            // register to each worker an encrypt function that will look for something to encrypt and send it
            // to the worker
            w.encrypt = function (){
                // look for something to encrypt
                for (var i=0; i < context.toEncrypt.length; i++){
                    if (context.toEncrypt[i] in context.encryptionCache){ // it has been encrypted in the past
                        context.toEncrypt.shift();
                        context.remaining -= 1;
                    }
                    else{
                        w.postMessage({"enc": {"plain": context.toEncrypt.shift(), "K": context.K, "S": context.S}})
                        break
                    }
                }
            };

            // register event for the result
            w.onmessage = function(e) {
                var task = Object.keys(e.data)[0];

                switch (task) {
                    case "enc":
                        var plain = e.data[task]["plain"];
                        var cipher = e.data[task]["cipher"];

                        context.encryptionCache[plain] = cipher;
                        context.remaining -= 1;

                        document.getElementById("encrypted").innerHTML += JSON.stringify(e.data[task]) + "<br>"; // todo delete

                        // start another encryption
                        w.encrypt();
                        break;

                    case "dec":
                        // todo
                        break;

                    default:
                        alert("Message received from the worker not recognized:" + JSON.stringify(e.data))
                }
            };

            w.postMessage({"init": {"import": importPath}});

            context.workers.push(w)
        }
    },

    initEncrypt: function(){
        var context = i2b2.MedCo.ctrlr.background;
        // init the ephemeral key of the light encryption
        [context.K, context.S] = LightEncryptStr_init(context.AggregateKey);
        context.encryptionCache = {}; // old ciphertexts are not valid anymore
    },

    // receives a list of plaintexts to be encrypted and append it to the toEncrypt list
    toBeEncrypted: function(toEncrypt){
        var context = i2b2.MedCo.ctrlr.background;
        // other plaintext to be encrypted are appended to the list
        context.toEncrypt = context.toEncrypt.concat(toEncrypt);
        context.remaining += toEncrypt.length;

        // sends to each worker something to encrypt
        for(var i=0; i < context.workers.length; i++){
            context.workers[i].encrypt(); // function registered to the worker (see init)
        }
    },

    // checks if there are other plaintexts to be encrypted
    done: function() {
        return i2b2.MedCo.ctrlr.background.remaining == 0;
    }

};


// all worker code here (this function is just used to create workers in the i2b2.MedCo.ctrlr.background.init)
function worker_code() {
    // a worker responds to 3 different messages, init, encrypt, decrypt
    onmessage = function(e) {
        // e.data is in one of the following forms -> response:
        // - {"init": {"import": "http://.../cryptolib.js"}} -> no response
        // - {"enc": {"plain": "123", "K": "...", "S": "..."}} -> {"enc": {"plain": "<plaintext>", "cipher": "<ciphertext>"}}
        // - todo {"dec": { ... }} -> {"dec": {"<ciphertext>": "<plaintext>"}}

        var task = Object.keys(e.data)[0];

        switch (task){
            case "init":
                self.importScripts(e.data[task]["import"]); // e.g. self.importScripts("http://localhost:63342/i-b-webclient/tools/gopherjsCrypto/tests/scripts/useme.js");
                console.log("worker says: imported cryptolib");
                break;

            case "enc":
                var plain = e.data[task]["plain"];
                var K = e.data[task]["K"];
                var S = e.data[task]["S"];
                var cipher = LightEncryptStr(plain, K, S);

                console.log("worker says: plain: " + plain + "; cipher: ", cipher);

                // {"enc": {"<plaintext>": "<ciphertext>"}}
                var response = {};
                response[task] = {};
                response[task]["plain"] = plain;
                response[task]["cipher"] = cipher;
                postMessage(response);

                break;

            case "dec":
                // todo
                break;

            default:
                alert("worker says: message not recognized: " + JSON.stringify(e.data))
        }
    };
}