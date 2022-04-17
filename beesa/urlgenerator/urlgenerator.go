package urlgenerator

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// All possibel references: https://www.eddywm.com/lets-build-a-url-shortener-in-go-part-3-short-link-generation/

// ============================
// Hashing  initialUrl + userId url with sha256.
// Here userId is added to prevent providing similar shortened urls
// to separate users in case they want to shorten exact same link,
// it's a design decision,
// so some implementations do this differently.
// Derive a big integer number from the hash bytes generated during the hasing.
// Finally apply base58  on the derived big integer value and pick the first 8 characters

func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}

// Algo for initial input
func sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

// Algorithm for final output of the process
func base58Encoded(bytes []byte) string {
	enc := base58.BitcoinEncoding
	encoded, err := enc.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}
