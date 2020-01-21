require('./wasm_exec.js');

const go = new Go();
module.exports.CryptoEngine = {
    GenerateKeyPair: function() {},
    EncryptInt: function() {},
    DecryptInt: function() {},
    AggregateKeys: function() {},
    Ready: Promise
};

module.exports.CryptoEngine.Ready = WebAssembly.instantiateStreaming(fetch("assets/medco-unlynx-js.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    module.exports.CryptoEngine.GenerateKeyPair = function() {
        return GenerateKeyPair();
    };
    module.exports.CryptoEngine.EncryptInt = function(key, text) {
        return EncryptInt(key, text);
    };
    module.exports.CryptoEngine.DecryptInt = function(cipher, key) {
        return DecryptInt(cipher, key);
    };
    module.exports.CryptoEngine.AggregateKeys = function(text) {
        return AggregateKeys(text);
    };
    Promise.resolve();
});
