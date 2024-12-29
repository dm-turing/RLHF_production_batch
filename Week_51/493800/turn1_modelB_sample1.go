// Chunk Server
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// Chunk server implementation
type ChunkServer struct {
	chunks map[string][]byte // Map from chunk ID to chunk data
	mux    sync.RWMutex      // Mutex for concurrent access
}

func (cs *ChunkServer) ReadChunk(chunkID string, chunk *[]byte) error {
	cs.mux.RLock()
	defer cs.mux.RUnlock()
	*chunk, ok := cs.chunks[chunkID]
	if !ok {
		return fmt.Errorf("chunk %s not found", chunkID)
	}
	return nil
}

func (cs *ChunkServer) WriteChunk(args *ChunkWriteArgs, reply *bool) error {
	cs.mux.Lock()
	defer cs.mux.Unlock()
	cs.chunks[args.ChunkID] = args.Chunk
	*reply = true
	return nil
}

// Chunk Write arguments
type ChunkWriteArgs struct {
	ChunkID string
	Chunk   []byte
}

func main() {
	// Create a new chunk server
	cs := &ChunkServer{
		chunks: make(map[string][]byte),
	}

	// Register the RPC server
	rpc.Register(cs)

	// Start listening on a port
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error listening:", err)
	}
	defer listener.Close()

	// Serve RPC connections forever
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting:", err)
		}
		go rpc.ServeConn(conn)
	}
}
