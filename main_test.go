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
