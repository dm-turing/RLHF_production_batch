// Name Server
package main

import (
	"fmt"
	"sync"
)

// Name server implementation
type NameServer struct {
	fileToChunks map[string][]string // Map from file name to chunk IDs
	mux          sync.RWMutex        // Mutex for concurrent access
}

func (ns *NameServer) RegisterChunk(args *RegisterChunkArgs, reply *bool) error {
	ns.mux.Lock()
	defer ns.mux.Unlock()
	ns.fileToChunks[args.FileName] = append(ns.fileToChunks[args.FileName], args.ChunkID)
	*reply = true
	return nil
}

func (ns *NameServer) GetChunks(fileName string, chunkIDs *[]string) error {
	ns.mux.RLock()
	defer ns.mux.RUnlock()
	*chunkIDs, ok := ns.fileToChunks[fileName]
	if !ok {
		return fmt.Errorf("file %s not found", fileName)
	}
	return nil
}

// Register Chunk arguments
