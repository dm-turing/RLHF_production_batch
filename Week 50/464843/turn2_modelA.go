package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

// MediaEntry represents a single rich media entry.
type MediaEntry struct {
	ID          int64     `json:"id"`
	UserID      string    `json:"userId"`
	ContentID   string    `json:"contentId"`
	Name        string    `json:"name"`
	FileType    string    `json:"fileType"`
	FileSize    int64     `json:"fileSize"`
	UploadedAt  time.Time `json:"uploadedAt"`
	Geolocation string    `json:"geolocation"`
}

// MediaLibrary is a collection of media entries using a map for efficient lookups and removals.
type MediaLibrary struct {
	entries    map[int64]MediaEntry
	entryOrder []int64
}

// Add adds a new media entry to the library.
func (lib *MediaLibrary) Add(entry MediaEntry) {
	entry.ID = int64(len(lib.entries) + 1) // Simple auto-increment
	lib.entries[entry.ID] = entry
	lib.entryOrder = append(lib.entryOrder, entry.ID)
}

// FindByID finds a media entry by ID.
func (lib *MediaLibrary) FindByID(id int64) (MediaEntry, bool) {
	return lib.entries[id], true
}

// ListFiltered lists media entries filtered by specific criteria.
func (lib *MediaLibrary) ListFiltered(userID string, contentID string, after time.Time) []MediaEntry {
	var filtered []MediaEntry
	for _, id := range lib.entryOrder {
		entry := lib.entries[id]
		if entry.UserID == userID && entry.ContentID == contentID && entry.UploadedAt.After(after) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

// Remove removes a media entry by ID.
func (lib *MediaLibrary) Remove(id int64) bool {
	delete(lib.entries, id)
	for i, entryID := range lib.entryOrder {
		if entryID == id {
			lib.entryOrder = append(lib.entryOrder[:i], lib.entryOrder[i+1:]...)
			return true
		}
	}
	return false
}

// Save saves a media entry to a specified folder.
func (entry *MediaEntry) Save(folder string) error {
	filename := entry.Name
	filePath := path.Join(folder, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer f.Close()

	_, err = f.Write([]byte("placeholder content")) // Replace with actual file content
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}

func main() {
	library := MediaLibrary{
		entries: make(map[int64]MediaEntry),
	}

	// Example usage
	filePath := "sample_image.jpg"
	fileName := path.Base(filePath)
	fileType := strings.Split(fileName, ".")[1]
	fileSize := int64(len([]byte("placeholder content"))) // Replace with actual file size
	geo := "40.7128, -74.0060"                            // Example geolocation

	mediaEntry := MediaEntry{
		UserID:      "user123",
		ContentID:   "post456",
		Name:        fileName,
		FileType:    fileType,
		FileSize:    fileSize,
		UploadedAt:  time.Now(),
		Geolocation: geo,
	}

	library.Add(mediaEntry)
	mediaEntry.Save("/path/to/media") // Replace with actual file save path

	// Find media entry by ID
	entry, found := library.FindByID(1)
	if found {
		fmt.Printf("Found media entry: %#v\n", entry)
	} else {
		fmt.Println("Media entry not found.")
	}

	// List filtered entries
	filtered := library.ListFiltered("user123", "post456", time.Now().AddDate(0, -1, 0))
	fmt.Println("Filtered media entries:")
	for _, entry := range filtered {
		fmt.Printf("%#v\n", entry)
	}

	// Remove a media entry
	library.Remove(1)
	fmt.Println("Media entry removed.")
}
