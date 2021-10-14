package main

import (
	"testing"
)

func TestFacNaive(t *testing.T) {
	var tests = []struct {
		name string
		n    int
		want int
	}{
		{"invalid fac", -1, 0},
		{"fac 0", 0, 1},
		{"fac 1", 1, 1},
		{"fac 2", 2, 2},
		{"fac 3", 3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := facNaive(tt.n)
			if got != tt.want {
				t.Errorf("wrong factorial result, got %d, want %d", got, tt.want)
			}
		})
	}

}

func BenchmarkFacNaive_10(b *testing.B) {
	var tests = []struct {
		name string
		n    int
	}{
		{"fac 10", 10},
		{"fac 11", 11},
		{"fac 13", 13},
		{"fac 15", 15},
		{"fac 20", 20},
		{"fac 25", 25},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = facNaive(tt.n)
			}
		})
	}
}

func TestFacDP(t *testing.T) {
	var tests = []struct {
		name string
		n    int
		want int
	}{
		{"invalid fac", -1, 0},
		{"fac 0", 0, 1},
		{"fac 1", 1, 1},
		{"fac 2", 2, 2},
		{"fac 3", 3, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := facDP(tt.n)
			if got != tt.want {
				t.Errorf("wrong factorial result, got %d, want %d", got, tt.want)
			}
		})
	}

}

func BenchmarkFacDP(b *testing.B) {
	var tests = []struct {
		name string
		n    int
	}{
		{"fac 10", 10},
		{"fac 11", 11},
		{"fac 13", 13},
		{"fac 15", 15},
		{"fac 20", 20},
		{"fac 23", 23},
		{"fac 25", 25},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = facDP(tt.n)
			}
		})
	}
}
