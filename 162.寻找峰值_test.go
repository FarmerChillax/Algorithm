package algorithm

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 3, 1},
			},
			want: 2,
		},
		{
			name: "case-2",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: 5,
		},
		{
			name: "case-3",
			args: args{
				nums: []int{4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "case-4",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 3,
		},
		{
			name: "case-5",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
