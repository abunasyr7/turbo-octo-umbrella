package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
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

	for i :=0; i < size; i++ {
		result[i] = rand.Intn(100)
	}

	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int) {
	if len(data) == 0 {
		return 0
	}

	max := data[0]

	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
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
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
