package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func CalculateSHA256Checksum(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return ""
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Printf("Error calculating SHA256 checksum: %v\n", err)
		return ""
	}

	checksum := hex.EncodeToString(hash.Sum(nil))
	return checksum
}
