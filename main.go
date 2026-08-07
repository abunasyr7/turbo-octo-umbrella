package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var ErrEmptySlice = errors.New("cannot find maximum of an empty slice")

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return make([]int, 0)
	}

	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = rand.Intn(100)
	}

	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]

	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS

	if chunkSize == 0 {
		return maximum(data)
	}

	maxes := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)

		go func(chunk []int, idx int) {
			defer wg.Done()
			maxes[idx] = maximum(chunk)
		}(data[start:end], i)
	}

	wg.Wait()

	return maximum(maxes)

}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	singleThreadMax := maximum(data)
	singleThreadElapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", singleThreadMax, singleThreadElapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	multiThreadMax := maxChunks(data)
	multiThreadElapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d µs\n", multiThreadMax, multiThreadElapsed)
}
