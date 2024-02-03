package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const M = 8
	const N = 2
	const BufferSize = 10

	var buffer = make([]int, BufferSize)
	var mutex = &sync.Mutex{}
	var readCh = make(chan bool)

	// Writing goroutines
	for i := 0; i < N; i++ {
		go func(writerID int) {
			for {
				mutex.Lock()
				buffer[0]++ // Simulate writing to the buffer
				fmt.Printf("Writer %d wrote to buffer\n", writerID)
				mutex.Unlock()
				readCh <- true // Signal that data is written
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

	for i := 0; i < M; i++ {
		go func(readerID int) {
			for {
				<-readCh // Wait for data to be written
				mutex.Lock()
				fmt.Printf("Reader %d read from buffer: %d\n", readerID, buffer[0]) // Simulate reading from buffer
				mutex.Unlock()
				time.Sleep(60 * time.Millisecond)
			}
		}(i)
	}

	select {}

}
