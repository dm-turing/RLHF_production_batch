package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func calculateChecksum(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	hash.Write(data)
	checksum := hex.EncodeToString(hash.Sum(nil))
	return checksum, nil
}

func main() {
	filePath := "recovered.txt"
	checksum, err := calculateChecksum(filePath)
	if err != nil {
		fmt.Println("Error calculating checksum:", err)
	} else {
		fmt.Println("Calculated Checksum:", checksum)
	}
}
