// Link: https://www.youtube.com/watch?v=Yc1yd2kSX4E&t=571s

package main

// Make Imports

import (
	"fmt"

	concurrency "github.com/spencerguo11/go_learning_spencer/CONCURRENCY"
)

// Main function

func main() {
	/*
		go time.Sleep(5 * time.Second) // by adding go in front of the Sleep function, it will run this function concurrently
		fmt.Println("Goodbye")
	*/

	/*
		// Create Wait Group
		var wg sync.WaitGroup

		// Sleep Concurrently
		wg.Add(1)
		go concurrency.Sleep(2*time.Second, &wg)

		// Wait for Wait Group
		wg.Wait()
		fmt.Println("Goodbye!")

	*/

	// Create Input Data
	numbers := []float64{2.0, 4.0, 8.0, 16.0, 32.0, 64.0, 128.0, 256.0}

	// Create Generator
	chGenerator := concurrency.Generator(numbers)

	// Spawn Workers
	var workers []<-chan float64

	for i := 0; i < 4; i++ {
		chWorker := concurrency.Worker(chGenerator)
		workers = append(workers, chWorker)

	}

	// Collect Results
	chResult := concurrency.Merge(workers)

	// Block + Print Results
	for sq := range chResult {
		fmt.Println(sq)
	}
}
