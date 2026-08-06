package main

import (
	"errors"
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
		wantErr bool
	} {
		{
			name: 	"simple slice",
			nums:    []int{3, 7, 1, 9, 4},
			want:    9,
			wantErr: false,
		}, 	
		{
			name:    "один элемент",
			nums:    []int{42},
			want:    42,
			wantErr: false,
		},
		{
			name:    "все элементы одинаковые",
			nums:    []int{5, 5, 5, 5},
			want:    5,
			wantErr: false,
		},
		{
			name:    "отрицательные числа",
			nums:    []int{-10, -3, -7, -1},
			want:    -1,
			wantErr: false,
		},
		{
			name:    "максимум в начале слайса",
			nums:    []int{100, 2, 3},
			want:    100,
			wantErr: false,
		},
		{
			name:    "пустой слайс",
			nums:    []int{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "nil слайс",
			nums:    nil,
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maximum(tt.nums)

			if (err != nil) != tt.wantErr {
				t.Fatalf("maximum(%v) error = %v, wantErr %v", tt.nums, err, tt.wantErr)
			}

			if tt.wantErr && !errors.Is(err, ErrEmptySlice) {
				t.Errorf("maximum(%v) error = %v, ожидалась ErrEmptySlice", tt.nums, err)
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

