import {Point, Scalar} from "@dedis/kyber";
import {newCurve} from "@dedis/kyber/curve";
import {PointToInt} from "./mapping";

// CipherText is an ElGamal encrypted point.
class CipherText {
	K: Point;
	C: Point;

	constructor(K: Point, C: Point) {
		this.K = K;
		this.C = C;
	}

	toString(): string {
		let cstr = "nil";
		let kstr = cstr;
		if (this.C != null) {
			cstr = this.C.toString().slice(1, 7)
		}
		if (this.K != null) {
			kstr = this.K.toString().slice(1, 7)
		}
		let str = "";
		return str.concat("CipherText{",cstr,",",kstr,"}")
	}

}

const arrayBufferToBuffer = require('arraybuffer-to-buffer');
const curve25519 = newCurve("edwards25519");

// Encryption
//______________________________________________________________________________________________________________________

/**
 * Encrypts an integer with the cothority key
 * @returns {CipherText}
 * @param pk
 * @param x
 */
export function EncryptInt(pk: Point, x: number): CipherText {
	return encryptPoint(pk, IntToPoint(x))
}

/**
 * Maps an integer to a point in the elliptic curve.
 * @returns {Point}
 * @param x
 */
function IntToPoint(x: number): Point {
	let B = curve25519.point().base();
	let i = curve25519.scalar().setBytes(toBytesInt32(x));
	return curve25519.point().mul(i, B)
}

/**
 * Creates an elliptic curve point from a non-encrypted point and encrypt it using ElGamal encryption.
 * @returns {CipherText}
 * @param pk
 * @param M
 */
function encryptPoint(pk: Point, M: Point): CipherText {
    let B = curve25519.point().base();
    let r = curve25519.scalar().pick(); // ephemeral private key
    // ElGamal-encrypt the point to produce ciphertext (K,C).
    let K = curve25519.point().mul(r, B);	// ephemeral DH public key
    let S = curve25519.point().mul(r, pk);	// ephemeral DH shared secret
    let C = curve25519.point().add(S, M);   // message blinded with secret
    return new CipherText(K,C)
}

// Decryption
//______________________________________________________________________________________________________________________

/**
 * Decrypts an integer from an ElGamal cipher text where integer are encoded in the exponent.
 * @returns {number}
 * @param prikey
 * @param cipher
 */
export function DecryptInt(prikey: Scalar, cipher: CipherText): number {
	let M = decryptPoint(prikey, cipher);
	return PointToInt[M.toString()];
}

/**
 * Decrypts an elliptic point from an El-Gamal cipher text.
 * @returns {Point}
 * @param prikey
 * @param c
 */
function decryptPoint(prikey: Scalar, c: CipherText): Point {
	let S = curve25519.point().mul(prikey, c.K);	// regenerate shared secret
	return curve25519.point().sub(c.C, S)		    // use to un-blind the message
}

// Utilities
//______________________________________________________________________________________________________________________

/**
   * Generates a random pair of keys for the user to be used during this instance.
   * @returns {[string, string]}
   */
export function GenerateKeyPair() {
	let privKey = curve25519.scalar().pick();
	let pubKey = curve25519.point().mul(privKey, null);
	return [privKey, pubKey];
}

// Representation
//______________________________________________________________________________________________________________________



// Marshal
//______________________________________________________________________________________________________________________

/**
 * Converts a Number to Buffer of
 * @returns {Buffer}
 * @param x
 */
function toBytesInt32 (x: number): Buffer {
	let arr = new ArrayBuffer(4); // an Int32 takes 4 bytes
	let view = new DataView(arr);
	view.setUint32(0, x, true); // byteOffset = 0; litteEndian = false
	return arrayBufferToBuffer(arr);
}





