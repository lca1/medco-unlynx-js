# Transpile Unlynx
To make use of the homomorphic encryption used by Unlynx we had to transpile the needed function to Javascript so to use them from within the browser. Moreover, since the decryption solves a discrete logarithm problem we had to produce a table with 10.000 entries to map points (ciphers) to the corresponting integer (plaintext). This table allowed to accelerate the otherwise very slow decription. 
The `gopherjsCrypto/` folder is structured as follows:
- `generate/`: folder containing two go scripts, one to genarate private/public keys (of the webclient user) and one to generate the reverse mapping table. The generated table is stored in `mappingTable/mapping.go` as a go variable called  `PointToInt`.
- `mylib/`: contains the wrappers to all the needed Unlynx functions as well as the tests to check everything works properly. We had to wrap them to simplify their interface so that, when transpiled, in Javascript we could deal **only with strings** (and not with more cumbersome structures, e.g. abstract.Point, which gave problems when used in Javascript). 
- `tests/`: contains the html page and javascript code to check that the transpiled functions work when called from whithin a browser. After generating the `cryptolib.js` and `cryptolib.js.map` copy them in the `tests/scripts/` folder and open in the browser `tests/example.html` to test the transpiled functions. In this folder there is also the code to test the background encryption (to avoid freezing the page while decrypting) before it *got integrated in the MedCo plugin*.
- `crypto_javascript.go`: script which is compiled with gopherjs to generate `cryptolib.js` and `cryptolib.js.map` with the following command:
    ```sh
    $ gopherjs build -m crypto_javascript.go -o cryptolib.js
    ```
### Libraries
To transpile the Unlynx code I used the following setting:
- Installed [`GopherJS 1.9-1`](https://github.com/gopherjs/gopherjs/tree/go1.9)
- Used [`Onet`](https://github.com/dedis/onet) on *master* branch.
- Used [`Unlynx`](https://github.com/lca1/unlynx) on *newI2B2* branch. I had issues in transpiling `onet` so I ended up in usng this version of Unlynx shich imports `onet.v1` instead of `onet`.

