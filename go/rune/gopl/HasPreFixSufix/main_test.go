package main

import (
	"strings"
	"testing"
)

func TestHasPrefix(t *testing.T) {
	var tests = []struct {
		name   string
		s      string
		prefix string
		want   bool
	}{
		{"Test normal string", "Hello", "He", true},
		{"Test prefix longer than original string", "Hello", "HelloKitty", false},
		{"Test empty string", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasPrefix(tt.s, tt.prefix)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	var tests = []struct {
		name   string
		s      string
		suffix string
		want   bool
	}{
		{"Test normal string", "Hello", "llo", true},
		{"Test suffix longer than original string", "Hello", "lloKitty", false},
		{"Test empty string", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasSuffix(tt.s, tt.suffix)
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		name   string
		s      string
		substr string
		want   bool
	}{
		{"Test when substr is suffix", "Hello", "llo", true},
		{"Test when substr is at middle", "Hello", "ell", true},
		{"Test when substr is prefix", "Hello", "Hel", true},
		{"Test when substr is not within s", "Hello", "Heo", false},
		{"Test when s and substr are empty strings", "", "", true},
		{"Test when substr is empty string", "Hello", "", true},
	}

	t.Run("HasPrefix version of Contains", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Contains_HasPrefix(tt.s, tt.substr)
				if got != tt.want {
					t.Errorf("got %t, want %t", got, tt.want)
				}
			})
		}
	})

	t.Run("HasSuffix version of Contains", func(t *testing.T) {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := Contains_HasSuffix(tt.s, tt.substr)
				if got != tt.want {
					t.Errorf("got %t, want %t", got, tt.want)
				}
			})
		}
	})

}

func BenchmarkContains(b *testing.B) {
	sSuffix30 := "sdjaskdlfjlskadjfklasjdflHello"
	sPrefix30 := "Helloasdjfksdjfkjskddfksdjfkid"
	c30 := "123456789012345678901234567890"
	sPrefix60 := sPrefix30 + c30
	sPrefix90 := sPrefix60 + c30
	subStr := "Hello"

	b.Run("my Contains using Contains_HasPrefix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Contains_HasPrefix(sSuffix30, subStr)
		}
	})

	b.Run("my Contains using Contains_HasSuffix", func(b *testing.B) {
		b.Run("30 characters", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Contains_HasSuffix(sPrefix30, subStr)
			}
		})
		b.Run("60 characters", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Contains_HasSuffix(sPrefix60, subStr)
			}
		})
		b.Run("90 characters", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Contains_HasSuffix(sPrefix90, subStr)
			}
		})

	})
	b.Run("30 chars using strings.Contains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.Contains(sPrefix30, subStr)
		}
	})
	b.Run("60 chars using strings.Contains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.Contains(sPrefix60, subStr)
		}
	})
	b.Run("90 chars using strings.Contains", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.Contains(sPrefix90, subStr)
		}
	})

}
