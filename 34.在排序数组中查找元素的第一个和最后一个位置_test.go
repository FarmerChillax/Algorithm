package algorithm

import (
	"testing"
)

func Test_getLeftBorder(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 1,
		},
		{
			name: "case-2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: -2,
		},
		{
			name: "case-3",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 6,
			},
			want: 2,
		},
		{
			name: "case-4",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeftBorder(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("getLeftBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRightBorder(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// TODO: Add test cases.
		{
			name: "case-1",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 3,
		},
		{
			name: "case-2",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: -2,
		},
		{
			name: "case-3",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 6,
			},
			want: 4,
		},
		{
			name: "case-4",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRightBorder(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("getRightBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
