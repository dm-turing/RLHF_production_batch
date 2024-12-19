package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	dataSize = 10000 // Set to a large value to benchmark performance
)

func benchmarkMediaLibrary(b *testing.B) {
	lib := MediaLibrary{}
	rand.Seed(time.Now().UnixNano())

	// Insert data
	for i := 0; i < dataSize; i++ {
		entry := MediaEntry{
			ID:          int64(i + 1),
			UserID:      "user",
			ContentID:   "post",
			Name:        fmt.Sprintf("media%d.jpg", i),
			FileType:    "jpg",
			FileSize:    1024,
			UploadedAt:  time.Now(),
			Geolocation: "0,0",
		}
		lib.Add(entry)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		idToFind := rand.Int63n(int64(dataSize)) + 1
		_, _ = lib.FindByID(idToFind) // FindByID

		// Remove a random entry (half the time)
		if i%2 == 0 {
			idToRemove := rand.Int63n(int64(len(lib.entries))) + 1
			lib.Remove(idToRemove)
		}
	}
}

func BenchmarkMediaLibrary_SmallData(b *testing.B) {
	const smallDataSize = 100
	dataSize = smallDataSize
	benchmarkMediaLibrary(b)
}

func BenchmarkMediaLibrary_LargeData(b *testing.B) {
	const largeDataSize = 10000
	dataSize = largeDataSize
	benchmarkMediaLibrary(b)
}
