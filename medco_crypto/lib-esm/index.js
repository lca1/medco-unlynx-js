import { newCurve } from "@dedis/kyber/curve";
import { PointToInt } from "./mapping";
var CipherText = (function () {
    function CipherText(K, C) {
        this.K = K;
        this.C = C;
    }
    CipherText.prototype.toString = function () {
        var cstr = "nil";
        var kstr = cstr;
        if (this.C != null) {
            cstr = this.C.toString();
        }
        if (this.K != null) {
            kstr = this.K.toString();
        }
        var str = "";
        return str.concat("CipherText{", kstr, ",", cstr, "}");
    };
    CipherText.prototype.fromString = function (str) {
        alert(str);
        var index = str.indexOf(",");
        var kString = str.slice(11, index);
        var cString = str.slice(index + 1, str.length);
        alert(kString);
        alert(cString);
        return new CipherText(FromStringToPoint(kString), FromStringToPoint(cString));
    };
    return CipherText;
}());
var arrayBufferToBuffer = require('arraybuffer-to-buffer');
var curve25519 = newCurve("edwards25519");
export function EncryptInt(pk, x) {
    return encryptPoint(pk, IntToPoint(x));
}
function IntToPoint(x) {
    var B = curve25519.point().base();
    var i = curve25519.scalar().setBytes(toBytesInt32(x));
    return curve25519.point().mul(i, B);
}
function encryptPoint(pk, M) {
    var B = curve25519.point().base();
    var r = curve25519.scalar().pick();
    var K = curve25519.point().mul(r, B);
    var S = curve25519.point().mul(r, pk);
    var C = curve25519.point().add(S, M);
    return new CipherText(K, C);
}
export function DecryptInt(prikey, cipher) {
    var M = decryptPoint(prikey, cipher);
    return PointToInt[M.toString()];
}
function decryptPoint(prikey, c) {
    var S = curve25519.point().mul(prikey, c.K);
    return curve25519.point().sub(c.C, S);
}
export function GenerateKeyPair() {
    var privKey = curve25519.scalar().pick();
    var pubKey = curve25519.point().mul(privKey, null);
    return [privKey, pubKey];
}
function hexToBytes(hex) {
    for (var bytes = [], c = 0; c < hex.length; c += 2)
        bytes.push(parseInt(hex.substr(c, 2), 16));
    return Buffer.from(bytes);
}
export function FromStringToPoint(p) {
    var buf = hexToBytes(p);
    var res = curve25519.point().base();
    res.unmarshalBinary(buf);
    return res;
}
export function FromStringToScalar(p) {
    var buf = hexToBytes(p);
    var res = curve25519.scalar().pick();
    res.unmarshalBinary(buf);
    return res;
}
function toBytesInt32(x) {
    var arr = new ArrayBuffer(4);
    var view = new DataView(arr);
    view.setUint32(0, x, true);
    return arrayBufferToBuffer(arr);
}
//# sourceMappingURL=index.js.map