package main

import (
	"fmt"
	"sort"
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

// MediaLibrary is a collection of media entries.
type MediaLibrary struct {
	entries []MediaEntry
}

// Add adds a new media entry to the library and sorts the entries by ID.
func (lib *MediaLibrary) Add(entry MediaEntry) {
	lib.entries = append(lib.entries, entry)
	sort.Slice(lib.entries, func(i, j int) bool {
		return lib.entries[i].ID < lib.entries[j].ID
	})
}

// FindByID finds a media entry by ID using binary search.
func (lib *MediaLibrary) FindByID(id int64) (MediaEntry, bool) {
	low := 0
	high := len(lib.entries) - 1

	for low <= high {
		mid := (low + high) / 2
		if lib.entries[mid].ID == id {
			return lib.entries[mid], true
		} else if lib.entries[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return MediaEntry{}, false
}

func main() {
	library := MediaLibrary{}

	mediaEntry1 := MediaEntry{ID: 1}
	mediaEntry2 := MediaEntry{ID: 2}
	mediaEntry3 := MediaEntry{ID: 3}
	library.Add(mediaEntry1)
	library.Add(mediaEntry2)
	library.Add(mediaEntry3)

	entry, found := library.FindByID(2)
	if found {
		fmt.Printf("Found media entry: %#v\n", entry)
	} else {
		fmt.Println("Media entry not found.")
	}
}
