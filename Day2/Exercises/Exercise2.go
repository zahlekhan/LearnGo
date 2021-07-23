package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func GetRating(ratingChan chan uint64, num int) {
	var wg sync.WaitGroup
	source := rand.NewSource(42)
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//Simulate call
			rand := rand.New(source)
			duration := time.Duration(rand.Intn(100))
			time.Sleep(duration * time.Millisecond)
			rating := uint64(rand.Intn(5))

			ratingChan <- rating
		}()
	}
	wg.Wait()
	close(ratingChan)
}

func Adder(totalRating *uint64, ratingChan chan uint64, done chan<- bool) {
	for rating := range ratingChan {
		atomic.AddUint64(totalRating, rating)
	}
	done <- true
}

func AverageRating(ratingChan chan uint64, numOfStudents int) float64 {
	totalRating := uint64(0)
	done := make(chan bool)

	go Adder(&totalRating, ratingChan, done)

	<-done
	return float64(totalRating) / float64(numOfStudents)
}

func main() {
	ratingChan := make(chan uint64)
	numOfStudents := 200
	go GetRating(ratingChan, numOfStudents)

	fmt.Println(AverageRating(ratingChan, numOfStudents))
}
