package main

import "testing"

func Test_sliding_window_increases(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name: "example",
			input: []int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliding_window_increases(tt.input); got != tt.want {
				t.Errorf("sliding_window_increases() = %v, want %v", got, tt.want)
			}
		})
	}
}
