import { Point, Scalar } from "@dedis/kyber";
declare class CipherText {
    K: Point;
    C: Point;
    constructor(K: Point, C: Point);
    toString(): string;
    fromString(str: string): CipherText;
}
export declare function EncryptInt(pk: Point, x: number): CipherText;
export declare function DecryptInt(prikey: Scalar, cipher: CipherText): number;
export declare function GenerateKeyPair(): (Scalar | Point)[];
export declare function FromStringToPoint(p: string): Point;
export declare function FromStringToScalar(p: string): Scalar;
export {};
