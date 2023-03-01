package cryptography

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/subtle"
	"math/big"
	"os"

	"golang.org/x/crypto/argon2"
)

// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#introduction
var (
	masterHashPasses    uint32 = 3
	masterHashMemory    uint32 = 64 * 1024
	masterHashThreads   uint8  = 4
	masterHashKeyLength uint32 = 32 // AES-256 needs 32-byte key

	rsaSize = 4096
)

func StrengthenMasterHash(masterHash []byte, salt []byte) []byte {
	if pepper := os.Getenv("CRYPTO_PEPPER"); pepper != "" {
		pepperedMasterHash := append(masterHash, []byte(pepper)...)
		return argon2.IDKey(pepperedMasterHash, salt, masterHashPasses, masterHashMemory, masterHashThreads, masterHashKeyLength)
	}
	return argon2.IDKey(masterHash, salt, masterHashPasses, masterHashMemory, masterHashThreads, masterHashKeyLength)

	// var StrengthenMasterHashIterations = 300000
	// var StrengthenMasterHashLength = 512 / 8
	// var StrengthenMasterHashDigest = sha512.New
	// return pbkdf2.Key(masterHash, salt, StrengthenMasterHashIterations, StrengthenMasterHashLength, StrengthenMasterHashDigest)
}

func CompareMasterHash(strengthenedMasterHash []byte, masterHash []byte, salt []byte) bool {
	testStrengthenedMasterHash := StrengthenMasterHash(masterHash, salt)
	return subtle.ConstantTimeCompare(strengthenedMasterHash, testStrengthenedMasterHash) == 1
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
	privateKey, _ := rsa.GenerateKey(rand.Reader, rsaSize)
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
