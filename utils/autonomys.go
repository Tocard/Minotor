package utils

import (
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/blake2b"
	"log"
	"math/big"
	"strings"
)

const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const ss58Prefix = "SS58PRE"

// decodeBase58 properly handles leading zeros and returns a fixed-length byte slice
func decodeBase58(input string) []byte {
	bigInt := big.NewInt(0)
	base := big.NewInt(58)

	for _, char := range input {
		index := strings.IndexRune(base58Alphabet, char)
		if index == -1 {
			return nil
		}
		bigInt.Mul(bigInt, base)
		bigInt.Add(bigInt, big.NewInt(int64(index)))
	}

	decoded := bigInt.Bytes()

	// Handle leading zeroes in Base58
	zeroCount := 0
	for _, char := range input {
		if char == '1' {
			zeroCount++
		} else {
			break
		}
	}

	return append(make([]byte, zeroCount), decoded...)
}

// DecodeSS58 extracts the public key and network prefix from an SS58 address
func DecodeSS58(address string) ([]byte, uint16, error) {
	decoded := decodeBase58(address)
	if len(decoded) < 35 || len(decoded) > 37 {
		log.Println("Decoded length:", len(decoded), "Data:", decoded)
		return nil, 0, errors.New("invalid SS58 address length")
	}

	var prefix uint16
	var prefixLen int
	if decoded[0] < 64 {
		prefix = uint16(decoded[0])
		prefixLen = 1
	} else {
		prefix = uint16(decoded[0]) | (uint16(decoded[1]) << 8)
		prefixLen = 2
	}

	pubKey := decoded[prefixLen : len(decoded)-2]
	checksum := decoded[len(decoded)-2:]

	hash, _ := blake2b.New(64, nil)
	hash.Write([]byte(ss58Prefix))
	hash.Write(decoded[:len(decoded)-2])
	fullChecksum := hash.Sum(nil)

	if checksum[0] != fullChecksum[0] || checksum[1] != fullChecksum[1] {
		log.Println("Expected checksum:", fullChecksum[:2], "Got:", checksum)
		return nil, 0, errors.New("checksum mismatch")
	}

	return pubKey, prefix, nil
}

func CleanAddressesArray(addresses []string) [][]byte {
	var _addresses [][]byte
	for _, address := range addresses {
		if strings.HasPrefix(address, "0x") {
			address = strings.Replace(address, "0x", "", -1)
		}
		pubKey, err := hex.DecodeString(address)
		if err != nil {
			log.Printf("Failed to parse %s\n: %s", address, err.Error())
			continue
		} else {
			_addresses = append(_addresses, pubKey)
		}
	}
	return _addresses
}
