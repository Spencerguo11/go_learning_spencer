package concurrency

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

func Sleep(t time.Duration, wg *sync.WaitGroup) {
	// Sleep
	time.Sleep(t)

	// Decrement Wait Group
	wg.Done()

}

// Square

func Square(number float64) float64 {
	// Define Time to Sleep
	sleepAmount := rand.Intn(5) + 1

	// Sleep
	time.Sleep(time.Duration(sleepAmount) * time.Second)

	// Square number
	square := math.Pow(number, 2)

	// Return Result
	return square
}

// Generator
func Generator(numbers []float64) <-chan float64 { // This means the return channel is read only
	// Create Channel
	chOut := make(chan float64) // at this point, the channel is readable and editable.

	// Feed Data into CHannel via Goroutine
	go func() { // loop through slice of intergers and write them into a channel
		for _, number := range numbers {
			chOut <- number
		}
		close(chOut)
	}()

	// Return Channel
	return chOut
}

// Worker
func Worker(chIn <-chan float64) <-chan float64 {
	// Create Output Channel
	chOut := make(chan float64)

	// Handle Incoming Data
	go func() {
		// Iterate Over Input Channel
		for numberIn := range chIn {
			// Square Number
			chOut <- Square(numberIn)
		}
		// Close Channel
		close(chOut)
	}()

	// Return Ouput Channel
	return chOut
}

// Internal Handover
func mergeCopy(chIn <-chan float64, chOut chan<- float64, wg *sync.WaitGroup) {
	for square := range chIn {
		chOut <- square
	}
	wg.Done()

}

// Merge
func Merge(chsIn []<-chan float64) <-chan float64 {
	// Create Wait Group for Workers
	var wg sync.WaitGroup
	wg.Add(len(chsIn))

	// Create Output Channel
	chOut := make(chan float64)

	// Copy Data
	for _, chIn := range chsIn {
		go mergeCopy(chIn, chOut, &wg)
	}

	// Start Goroutine to Wait for Workers to Finish
	go func() {
		wg.Wait()
		close(chOut)
	}()

	// Return Result Channel
	return chOut
}
