package main_test

import (
	. "github.com/lca1/medco-unlynx-js/javascriptLibrary"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
)

const trials = 500
const rndSeed = 1

func TestCrypto(t *testing.T) {
	rand.Seed(rndSeed)

	for i := int64(0); i < trials; i += 1 {
		secKey, pubKey := GenerateKeyPair()
		plainStart := strconv.FormatInt(rand.Int63n(1000), 10)

		cipher := EncryptInt(pubKey, plainStart)
		plainEnd := TempOverrideDecryptInt(cipher, secKey)
		require.Equal(t, plainStart, plainEnd)
	}
}
