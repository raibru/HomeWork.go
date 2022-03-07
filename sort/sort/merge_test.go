package sort

import (
	"testing"

	"github.com/raibru/HomeWork/sort/test"
)

func Test_mergeSort(t *testing.T) {
	// t.Error("Check Failure")
	// Given
	cases := make([]test.CaseSort, 0)
	err := test.ReadTestData("sort-data-success.json", &cases)
	if err != nil {
		t.Fatalf("Error read json data: %v", err)
	}

	for _, tc := range cases {
		t.Logf("INFO:  Merge sort data   -> %v", tc.Input)
		t.Logf("                  expect -> %v", tc.Expect)

		// When
		qs := MergeSort{}
		err := qs.Sort(tc.Input)

		t.Logf("                  sorted -> %v", tc.Input)

		// Then
		if err != nil {
			t.Fatalf("Test MergeSort return error: %v", err)
		}
		for i, exp := range tc.Expect {
			if exp != tc.Input[i] {
				t.Fatalf("FAIL: (%d) Test MergeSort expect %d but get %d", i, exp, tc.Input[i])
			}
		}
	}
}
