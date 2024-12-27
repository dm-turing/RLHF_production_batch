package main

import (
	"bytes"
	"os"

	// "os"
	"testing"
)

// same FileRecoverySystem and FileMetadata structures

func TestFileStorageRetrieval(t *testing.T) {
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

	// Restore file
	err := frs.RestoreFile("recovered.txt")
	if err != nil {
		t.Fatalf("Error restoring file: %v", err)
	}

	// Check if file content matches original
	f, err := os.Open("recovered.txt")
	if err != nil {
		t.Fatalf("Error opening restored file: %v", err)
	}
	defer f.Close()

	buf := make([]byte, 5120)
	n, err := f.Read(buf)
	if err != nil {
		t.Fatalf("Error reading restored file: %v", err)
	}

	if n != 5120 {
		t.Fatalf("Expected to read 5120 bytes, got %d", n)
	}

	expectedData := bytes.Join(blockData, nil)
	if !bytes.Equal(buf, expectedData) {
		t.Errorf("Restored file content does not match original")
	}
}

func TestFileConsistencyCheck(t *testing.T) {
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

	// Modify one block to induce corruption
	frs.DataStore[2][0] = 0xFF

	err := frs.CheckConsistency()
	if err == nil {
		t.Fatalf("Expected checksum mismatch error")
	}
}

func TestChecksumIntegrity(t *testing.T) {
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

	// Restore file
	err := frs.RestoreFile("recovered.txt")
	if err != nil {
		t.Fatalf("Error restoring file: %v", err)
	}

	// Check if checksum of restored file matches original
	f, err := os.Open("recovered.txt")
	if err != nil {
		t.Fatalf("Error opening restored file: %v", err)
	}
	defer f.Close()
	restoredData := make([]byte, 1024)
	_, err = f.Read(restoredData)
	// restoredData, err := file.ReadAll(f)
	if err != nil {
		t.Fatalf("Error reading restored file: %v", err)
	}

	actualChecksum := frs.CalculateChecksum(restoredData)
	if actualChecksum != frs.Metadata.Checksum {
		t.Errorf("Checksum mismatch: expected %s, got %s", frs.Metadata.Checksum, actualChecksum)
	}
}

func TestEdgeCases(t *testing.T) {
	frs := NewFileRecoverySystem(1024)

	// Test zero-length file
	frs.StoreFileMetadata("zero_length.txt", 0)
	frs.AddBlock(make([]byte, 1024)) // Not actually used, but needed for structure
	err := frs.RestoreFile("zero_length_recovered.txt")
	if err != nil {
		t.Fatalf("Error restoring zero-length file: %v", err)
	}

	stat, err := os.Stat("zero_length_recovered.txt")
	if err != nil {
		t.Fatalf("Error stat-ing zero-length file: %v", err)
	}
	if stat.Size() != 0 {
		t.Errorf("Zero-length file has size %d", stat.Size())
	}

	// Test file with a single block
	frs.StoreFileMetadata("single_block.txt", 1024)
	blockData := make([]byte, 1024)
	for j := 0; j < 1024; j++ {
		blockData[j] = byte(j % 256)
	}
	frs.AddBlock(blockData)

	frs.Metadata.Checksum = frs.CalculateChecksum(blockData)
	err = frs.RestoreFile("single_block_recovered.txt")
	if err != nil {
		t.Fatalf("Error restoring single block file: %v", err)
	}

	f, err := os.Open("single_block_recovered.txt")
	if err != nil {
		t.Fatalf("Error opening single block file: %v", err)
	}
	defer f.Close()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		t.Fatalf("Error reading single block file: %v", err)
	}

	if n != 1024 {
		t.Fatalf("Expected to read 1024 bytes, got %d", n)
	}

	if !bytes.Equal(buf, blockData) {
		t.Errorf("Single block file content does not match original")
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
