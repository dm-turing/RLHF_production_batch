package main  

import (  
    "fmt"
    "io/ioutil"
    "log"
    "math/rand"
    "os"
    "path/filepath"
    "sync"
    "time"
)

const (
    numGoroutines   = 5   // Number of concurrent readers and writers
    dataDir        = "data" // Directory containing log files
    readWriteRatio = 0.7  // Ratio of reads to writes (choose a value between 0 and 1)
)

var (
    fileList []string // List of log files to read/write from
    mu       sync.Mutex // Mutex for locking the fileList
    wg       sync.WaitGroup
)

func init() {
    rand.Seed(time.Now().UnixNano())
    // Initialize the fileList with all log files from the directory
    files, err := filepath.Glob(filepath.Join(dataDir, "*.log"))
    if err != nil {
        log.Fatalf("Error finding log files: %v", err)
    }
    fileList = files
}

func main() {
    wg.Add(numGoroutines)
    for i := 0; i < numGoroutines; i++ {
        go worker(i)
    }

    wg.Wait()
}

func worker(id int) {
    defer wg.Done()
    for {
        doReadWrite()
        // Add a short random delay to simulate interleaved I/O operations
        time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    }
}

func doReadWrite() {
    // Select random file from the list
    mu.Lock()
    n := len(fileList)
    if n == 0 {
        mu.Unlock()
        return
    }
    fileIndex := rand.Intn(n)
    selectedFile := fileList[fileIndex]
    mu.Unlock()

    doOperation(selectedFile)
}

func doOperation(filename string) {
    // Perform read or write operation with selective blocking
    op := rand.Float64()
    if op < readWriteRatio {
        doRead(filename)
    } else {
        doWrite(filename)
    }
}

func doRead(filename string) {
    // Lock only during file access to avoid simultaneous reads
    mu.Lock()
    defer mu.Unlock()

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Printf("Error reading file %s: %v", filename, err)
    } else {
        // Consume the read data for demonstration purposes
        _ = data
    }
}

func doWrite(filename string) {
    // Block write operation if there are more read threads than write threads
    mu.Lock()
    writeCount := 0
    for _, file := range fileList {
        if filepath.Base(file) == filepath.Base(filename) && dataDir == filepath.Dir(file) {
            writeCount++
        }
    }
    // Let a small ratio of writes continue if reads are heavily prioritized
    if float64(writeCount) / float64(numGoroutines) >= 2*readWriteRatio && rand.Float64() < 0.9 {
        mu.Unlock()
        return
    }
    mu.Unlock()

    content := fmt.Sprintf("This is a new write from goroutine %d.\n", filepath.Base(filename))
    err := ioutil.WriteFile(filename, []byte(content), 0644)
    if err != nil {
        log.Printf("Error writing file %s: %v",