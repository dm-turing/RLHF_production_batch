package main

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type Task struct {
	Data string
}

func main() {
	// Create a buffer pool for tasks to reduce memory allocation overhead
	var taskPool *sync.Pool = &sync.Pool{
		New: func() interface{} { return new(Task) },
	}

	// Prepare pipes for data transmission between stages
	readerPipe, writerPipe := io.Pipe()
	defer readerPipe.Close()
	defer writerPipe.Close()

	// Number of worker goroutines for each stage
	readerCount := 2
	transformerCount := 2
	writerCount := 2

	// Create and start worker goroutines
	wg := &sync.WaitGroup{}
	start := time.Now()

	// Reader stage
	for i := 0; i < readerCount; i++ {
		wg.Add(1)
		go reader(wg, taskPool, writerPipe)
	}

	// Transformer stage
	for i := 0; i < transformerCount; i++ {
		wg.Add(1)
		go transformer(wg, taskPool, writerPipe)
	}

	// Writer stage
	for i := 0; i < writerCount; i++ {
		wg.Add(1)
		go writer(wg, taskPool, readerPipe)
	}

	// Generate tasks and send them to the reader stage
	for i := 0; i < 1000; i++ {
		task := taskPool.Get().(*Task)
		task.Data = fmt.Sprintf("Task %d: Read", i)
		if _, err := writerPipe.Write([]byte(task.Data + "\n")); err != nil {
			fmt.Println("Error writing to writerPipe:", err)
			taskPool.Put(task)
			break
		}
		taskPool.Put(task)
	}

	// Close the writerPipe to signal the end of the input
	if err := writerPipe.Close(); err != nil {
		fmt.Println("Error closing writerPipe:", err)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Completed in %v\n", elapsed)
}

func reader(wg *sync.WaitGroup, taskPool *sync.Pool, writerPipe *io.PipeWriter) {
	defer wg.Done()
	defer taskPool.Put(taskPool.Get())

	r := io.NewLineReader(writerPipe)

	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			fmt.Println("Error reading from writerPipe:", err)
			if err == io.EOF {
				return
			}
		}
		if isPrefix {
			fmt.Println("Reader encountered prefix")
			continue
		}

		task := taskPool.Get().(*Task)
		task.Data = string(line)
		if _, err := writerPipe.Write([]byte(task.Data + "\n")); err != nil {
			fmt.Println("Error writing to writerPipe:", err)
			taskPool.Put(task)
			break
		}
		taskPool.Put(task)
	}
}

func transformer(wg *sync.WaitGroup, taskPool *sync.Pool, writerPipe *io.PipeWriter) {
	defer wg.Done()
	defer taskPool.Put(taskPool.Get())

	r := io.NewLineReader(writerPipe)

	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			fmt.Println("Error reading from writerPipe:", err)
			if err == io.EOF {
				return
			}
		}
		if isPrefix {
			fmt.Println("Transformer encountered prefix")
			continue
		}

		task := taskPool.Get().(*Task)
		task.Data = string(line) + " Transformed"
		if _, err := writerPipe.Write([]byte(task.Data + "\n")); err != nil {
			fmt.Println("Error writing to writerPipe:", err)
			taskPool.Put(task)
			break
		}
		taskPool.Put(task)
	}
}

func writer(wg *sync.WaitGroup, taskPool *sync.Pool, readerPipe *io.PipeReader) {
	defer wg.Done()
	defer taskPool.Put(taskPool.Get())

	r := io.NewLineReader(readerPipe)

	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			fmt.Println("Error reading from readerPipe:", err)
			if err == io.EOF {
				return
			}
		}
		if isPrefix {
			fmt.Println("Writer encountered prefix")
			continue
		}

		task := taskPool.Get().(*Task)
		task.Data = string(line) + " Written"
		fmt.Println(task.Data)
		taskPool.Put(task)
	}
}
