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

func Adder(totalRating *uint64, ratingChan chan uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	for rating := range ratingChan {
		atomic.AddUint64(totalRating, rating)
	}
}

func AverageRating(ratingChan chan uint64, numOfStudents int) float64 {
	totalRating := uint64(0)
	wg := sync.WaitGroup{}

	numOfAdder := 10
	for i := 1; i <= numOfAdder; i++ {
		wg.Add(1)
		go Adder(&totalRating, ratingChan, &wg)
	}

	wg.Wait()
	return float64(totalRating) / float64(numOfStudents)
}

func main() {
	ratingChan := make(chan uint64)
	numOfStudents := 200
	go GetRating(ratingChan, numOfStudents)

	fmt.Println(AverageRating(ratingChan, numOfStudents))
}
