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
		fmt.Println(increment.index)
		atomic.AddUint64(&letterFrequency[increment.index], increment.delta)
	}
}

func WordToLetters(wordChan chan string, incrementChan chan Increment) {
	defer close(incrementChan)
	for word := range wordChan {
		for _, letter := range word {
			newIncrement := Increment{int(unicode.ToLower(letter) - 'a'), 1}
			incrementChan <- newIncrement
		}
	}
}

func CountLetterFrequency(Words ...string) map[string]uint64 {

	letterFrequency := make([]uint64, 26)
	wg := sync.WaitGroup{}
	incrementChan := make(chan Increment)

	CounterWorkers := 2
	for i := 1; i <= CounterWorkers; i++ {
		wg.Add(1)
		go Counter(letterFrequency, incrementChan, &wg)
	}

	wordChan := make(chan string)
	WordToLettersWorkers := 1
	for i := 1; i <= WordToLettersWorkers; i++ {
		go WordToLetters(wordChan, incrementChan)
	}

	for _, word := range Words {
		wordChan <- word
	}
	close(wordChan)
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
	test := []string{"dog", "dog", "dog", "dog"}
	fmt.Println(CountLetterFrequency(test...))
}
