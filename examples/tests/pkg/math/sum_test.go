// Go in action
// @jeffotoni
// 2019-01-24

package math

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test_1: ", args{2, 2}, 4},
		{"test_2: ", args{-2, 6}, 4},
		{"test_3: ", args{-4, 8}, 4},
		{"test_4: ", args{5, 7}, 12},
		{"test_5: ", args{8, 8}, 15}, // forcing the error
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
