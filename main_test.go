package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"When the target is met", args{[]int{1, 3, 5, 7}, 6}, []int{0, 2}},
		{"When the target is not met", args{[]int{1, 2, 3, 8}, 6}, nil},
		{"When there is an empty massive", args{[]int{}, 0}, nil},
	}
	for _, test := range tests {
		got := TwoSum(test.args.nums, test.args.target)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("The TwoSum function doesn't work correctly %s, got %v, want %v", test.name, got, test.want)
		}

	}
}
