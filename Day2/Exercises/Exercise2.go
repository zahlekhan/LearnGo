package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Adder(totalRating *uint64, ratingChan chan uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	for rating := range ratingChan {
		atomic.AddUint64(totalRating, rating)
	}
}

func AverageRating(Ratings ...uint64) float64 {
	totalRating := uint64(0)
	wg := sync.WaitGroup{}
	ratingChan := make(chan uint64)

	numOfCounter := 10
	for i := 1; i <= numOfCounter; i++ {
		wg.Add(1)
		go Adder(&totalRating, ratingChan, &wg)
	}

	for _, rating := range Ratings {
		ratingChan <- rating
	}

	close(ratingChan)
	wg.Wait()

	return CalculateAverage(totalRating, len(Ratings))
}

func CalculateAverage(total uint64, count int) float64 {
	return float64(total) / float64(count)
}

func main() {
	ratings := []uint64{1, 2, 3, 4, 5}
	fmt.Println(AverageRating(ratings...))
}
