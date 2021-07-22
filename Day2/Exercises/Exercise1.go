package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unicode"
)

type Increment struct {
	index int
	delta uint64
}

func Counter(letterFrequency []uint64, incrementChan chan Increment, wg *sync.WaitGroup) {
	defer wg.Done()
	for increment := range incrementChan {
		atomic.AddUint64(&letterFrequency[increment.index], increment.delta)
	}
}

func CountLetterFrequency(Words ...string) map[string]uint64 {

	letterFrequency := make([]uint64, 26)
	wg := sync.WaitGroup{}
	incrementChan := make(chan Increment)

	numOfCounter := 10
	for i := 1; i <= numOfCounter; i++ {
		wg.Add(1)
		go Counter(letterFrequency, incrementChan, &wg)
	}

	for _, word := range Words {
		for _, letter := range word {
			newIncrement := Increment{int(unicode.ToLower(letter) - 'a'), 1}
			incrementChan <- newIncrement
		}
	}

	close(incrementChan)
	wg.Wait()

	return GenerateFrequency(letterFrequency)
}

func GenerateFrequency(letterFrequency []uint64) map[string]uint64 {
	result := make(map[string]uint64)
	for idx, val := range letterFrequency {
		result[string(rune(idx+'a'))] = val
	}
	return result
}

func main() {
	test := []string{"dog", "dog", "dog", "dog", "dog"}
	fmt.Println(CountLetterFrequency(test...))
}
