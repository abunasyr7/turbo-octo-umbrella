package main

import (
	"testing"
)

func TestGenerateRandoomElements(t *testing.T) {
	tests := []struct {
		name    string
		size      int
		want int
	}{
		{
			name: "Test with size 0",
			size: 0,
			want: 0,
		}, {
			name: "Test with size positive",
			size: 5,
			want: 5,
		},
		{
			name: "Test with size negative",
			size: -5,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if got == nil {
				t.Fatalf("generateRandomElements(%d) вернула nil, ожидался не-nil слайс", tt.size)
			} 
			
			if len(got) != tt.want {
				t.Errorf("len(got) = %d, want %d", len(got), tt.want)
			}

			for _, v := range got {
				if v < 0 || v >= 100 {
					t.Errorf("значение %d вне ожидаемого диапазона [0, 100)", v)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "обычный слайс", nums: []int{3, 7, 1, 9, 4}, want: 9},
		{name: "один элемент", nums: []int{42}, want: 42},
		{name: "все элементы одинаковые", nums: []int{5, 5, 5, 5}, want: 5},
		{name: "отрицательные числа", nums: []int{-10, -3, -7, -1}, want: -1},
		{name: "максимум в начале слайса", nums: []int{100, 2, 3}, want: 100},
		{name: "пустой слайс", nums: []int{}, want: 0},
		{name: "nil слайс", nums: nil, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.nums)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "пустой слайс", nums: []int{}, want: 0},
		{name: "один элемент", nums: []int{42}, want: 42},
		{name: "элементов меньше, чем частей", nums: []int{3, 9, 1, 4}, want: 9},
		{name: "элементов ровно столько, сколько частей", nums: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{
			name: "длина не делится на CHUNKS нацело",
			nums: []int{
				1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
				11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
				999,
			},
			want: 999,
		},
		{
			name: "большой слайс, максимум в середине",
			nums: []int{5, 3, 8, 1, 9, 2, 100, 4, 6, 7, 3, 2, 1, 9, 8, 5},
			want: 100,
		},
		{name: "отрицательные числа", nums: []int{-5, -3, -9, -1, -20, -100, -2, -4, -50}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.nums)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}