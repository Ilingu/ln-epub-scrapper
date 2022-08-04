package utils

import (
	"reflect"
	"testing"
)

type testCasesInput struct {
	arr    any
	slices int
}

type testCasesShape struct {
	input  testCasesInput
	expect []any
}

func TestCutArray(t *testing.T) {
	for i, test := range testCases {
		out := CutArray(test.input.arr.([]int), test.input.slices)

		isOk := true
		for j := range test.expect {
			if !reflect.DeepEqual(out[j], test.expect[j]) {
				isOk = false
			}
		}

		if !isOk {
			t.Errorf("Test #%d\nwant: %v\ngot: %v\n", i, test.expect, out)
		}
	}
}

var testCases = []testCasesShape{
	{
		input:  testCasesInput{arr: []int{1, 2, 3, 4, 5, 6, 7, 8}, slices: 4},
		expect: []any{[]int{1, 2}, []int{3, 4}, []int{5, 6}, []int{7, 8}},
	},
	{
		input:  testCasesInput{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, slices: 2},
		expect: []any{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}},
	},
	{
		input:  testCasesInput{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, slices: 3},
		expect: []any{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
	},
	{
		input:  testCasesInput{arr: []int{1, 2, 3, 4, 5, 6, 7}, slices: 2},
		expect: []any{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7}},
	},
	{
		input:  testCasesInput{arr: []int{1, 2, 3, 4, 5, 6}, slices: 5},
		expect: []any{[]int{1}, []int{2}, []int{3}, []int{4}, []int{5}, []int{6}},
	},
	{
		input:  testCasesInput{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, slices: 4},
		expect: []any{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{10, 11, 12}, []int{13, 14, 15}},
	},
	{
		input:  testCasesInput{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, slices: 3},
		expect: []any{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{10, 11}},
	},
}

func TestGetImageCover(t *testing.T) {
	filePath, ok := GetImageCover("https://i0.wp.com/perpetualdaydreams.com/wp-content/uploads/2020/11/921BE86A-9B4C-4E3B-BA92-26C1CF6E081A.jpeg?resize=212%2C300&ssl=1")
	if !ok {
		t.Fatal("cannot get image cover ❌")
	}

	t.Log(filePath)
}
