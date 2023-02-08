package cryptography

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/subtle"
	"math/big"

	"golang.org/x/crypto/pbkdf2"
)

var StrengthenMasterHashIterations = 150000
var StrengthenMasterHashLength = 512 / 8
var StrengthenMasterHashDigest = sha512.New

var RSASize = 4096

func StrengthenMasterHash(masterHash []byte, salt []byte) []byte {
	return pbkdf2.Key(masterHash, salt, StrengthenMasterHashIterations, StrengthenMasterHashLength, StrengthenMasterHashDigest)
}

func RandomBytes(n int) []byte {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return bytes
}

func RandomString(n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[random.Int64()]
	}
	return string(result)
}

func GenerateSignature(value string) (*rsa.PublicKey, []byte) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, RSASize)
	publicKey := &privateKey.PublicKey
	hashed := sha512.Sum512([]byte(value))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashed[:])
	return publicKey, signature
}

func VerifySignature(publicKey *rsa.PublicKey, signature []byte, value string) bool {
	hashed := sha512.Sum512([]byte(value))
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashed[:], signature) == nil
}

func ImportPublicKey(n []byte, e int) *rsa.PublicKey {
	nBuilder := new(big.Int)
	nBuilder.SetBytes(n)
	key := rsa.PublicKey{
		N: nBuilder,
		E: e,
	}
	return &key
}

func ConstantTimeCompare(x []byte, y []byte) bool {
	return subtle.ConstantTimeCompare(x, y) == 1
}
