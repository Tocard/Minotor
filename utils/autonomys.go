package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/mr-tron/base58"
	"log"
	"strings"
)

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

func DecodeSS58(address string) ([]byte, error) {
	decoded, err := base58.Decode(address)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base58 address: %w", err)
	}

	if len(decoded) < 35 {
		return nil, errors.New("invalid SS58 address length")
	}
	pubKey := decoded[1 : 1+32]
	log.Printf("Converted %s to 0x%s", address, hex.EncodeToString(pubKey))

	return pubKey, nil
}
