/**
 * @Author: chentong
 * @Date: 2023/11/03 11:45
 */

package math

import (
	"testing"
)

func TestCeil(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "Test Case 1",
			a:    7,
			b:    3,
			want: 3,
		},
		{
			name: "Test Case 2",
			a:    15,
			b:    2,
			want: 8,
		},
		{
			name: "Test Case 3",
			a:    -7,
			b:    3,
			want: -2,
		},
		{
			name: "Test Case 4",
			a:    7,
			b:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ceil(tt.a, tt.b); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "Test Case 1",
			a:    7,
			b:    3,
			want: 2,
		},
		{
			name: "Test Case 2",
			a:    15,
			b:    2,
			want: 7,
		},
		{
			name: "Test Case 3",
			a:    -7,
			b:    3,
			want: -3,
		},
		{
			name: "Test Case 4",
			a:    7,
			b:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Floor(tt.a, tt.b); got != tt.want {
				t.Errorf("Ceil() = %v, want %v", got, tt.want)
			}
		})
	}
}
