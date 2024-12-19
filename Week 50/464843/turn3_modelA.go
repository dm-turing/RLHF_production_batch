package main

import (
	"fmt"
	"sort"
	"time"
)

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

// MediaLibrary is a collection of media entries kept in a sorted order.
type SortedMediaLibrary struct {
	entries []MediaEntry
}

// Add adds a new media entry to the library and maintains sorted order.
func (lib *SortedMediaLibrary) Add(entry MediaEntry) {
	lib.entries = append(lib.entries, entry)
	sort.Slice(lib.entries, func(i, j int) bool { return lib.entries[i].ID < lib.entries[j].ID })
}

// FindByID finds a media entry by ID using binary search.
func (lib *SortedMediaLibrary) FindByID(id int64) (MediaEntry, bool) {
	i := sort.Search(len(lib.entries), func(j int) bool { return lib.entries[j].ID >= id })
	if i < len(lib.entries) && lib.entries[i].ID == id {
		return lib.entries[i], true
	}
	return MediaEntry{}, false
}

// Remove removes a media entry by ID and maintains sorted order.
func (lib *SortedMediaLibrary) Remove(id int64) bool {
	for i, entry := range lib.entries {
		if entry.ID == id {
			lib.entries = append(lib.entries[:i], lib.entries[i+1:]...)
			sort.Slice(lib.entries, func(i, j int) bool { return lib.entries[i].ID < lib.entries[j].ID })
			return true
		}
	}
	return false
}

func main() {
	lib := SortedMediaLibrary{}

	// Example usage
	for i := 0; i < 10; i++ {
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

	// Find media entry by ID
	entry, found := lib.FindByID(5)
	if found {
		fmt.Printf("Found media entry: %#v\n", entry)
	} else {
		fmt.Println("Media entry not found.")
	}

	// Remove a media entry
	lib.Remove(5)
	fmt.Println("Media entry removed.")
}
