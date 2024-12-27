package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

type FileMetadata struct {
	FileName string
	Size     int
	Blocks   []int
	Checksum string
}

type FileRecoverySystem struct {
	BlockSize int
	DataStore [][]byte
	Metadata  *FileMetadata
}

func NewFileRecoverySystem(blockSize int) *FileRecoverySystem {
	return &FileRecoverySystem{
		BlockSize: blockSize,
	}
}

func (frs *FileRecoverySystem) AddBlock(data []byte) {
	if len(data) != frs.BlockSize {
		panic(fmt.Sprintf("Invalid block size: %d", len(data)))
	}
	frs.DataStore = append(frs.DataStore, data)
}

func (frs *FileRecoverySystem) CalculateChecksum(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func (frs *FileRecoverySystem) StoreFileMetadata(fileName string, size int) {
	frs.Metadata = &FileMetadata{
		FileName: fileName,
		Size:     size,
		Blocks:   make([]int, size/frs.BlockSize),
	}
}

func (frs *FileRecoverySystem) RestoreFile(outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, blockIndex := range frs.Metadata.Blocks {
		data := frs.DataStore[blockIndex]
		_, err = file.Write(data)
		if err != nil {
			return err
		}
	}

	fmt.Printf("File restored: %s\n", outputFile)
	return nil
}

func (frs *FileRecoverySystem) CheckConsistency() error {
	checksumData := bytes.Join(frs.DataStore, nil)
	actualChecksum := frs.CalculateChecksum(checksumData)

	if actualChecksum != frs.Metadata.Checksum {
		return fmt.Errorf("Checksum mismatch: expected %s, got %s", frs.Metadata.Checksum, actualChecksum)
	}

	fmt.Println("File consistency check passed.")
	return nil
}

func main() {
	frs := NewFileRecoverySystem(1024)
	frs.StoreFileMetadata("example.txt", 5120)

	// Add sample blocks
	blockData := make([][]byte, 5)
	for i := 0; i < 5; i++ {
		blockData[i] = make([]byte, 1024)
		for j := 0; j < 1024; j++ {
			blockData[i][j] = byte((i*1024 + j) % 256)
		}
		frs.AddBlock(blockData[i])
	}

	// Calculate and store checksum
	frs.Metadata.Checksum = frs.CalculateChecksum(bytes.Join(blockData, nil))

	// Check consistency
	err := frs.CheckConsistency()
	if err != nil {
		fmt.Printf("Error checking consistency: %v\n", err)
	}

	// Restore file
	err = frs.RestoreFile("recovered.txt")
	if err != nil {
		fmt.Printf("Error restoring file: %v\n", err)
	}
}
