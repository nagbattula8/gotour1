package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tables := []struct {
		calls       int
		returnArray []int
	}{
		{0, []int{}},
		{4, []int{0, 1, 1, 2}},
	}

	for _, table := range tables {

		var arr []int

		f := fibonacci()

		for i := 0; i < table.calls; i++ {
			arr = append(arr, f())
		}

		if len(arr) == 0 && len(table.returnArray) == 0 {
			fmt.Println("0 case passed")
			continue
		}

		if !reflect.DeepEqual(arr, table.returnArray) {
			t.Error(arr, table.returnArray)
		} else {
			fmt.Println("input -> ", arr, " -> passed")
		}

	}
}
