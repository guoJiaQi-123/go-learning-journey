package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSum(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{6, 15}
		checkSum(t, got, want)
	})
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	numbersToSumLen := len(numbersToSum)
	sums = make([]int, numbersToSumLen)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tails := numbers[1:]
		sums = append(sums, Sum(tails))
	}
	return sums
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func Sum(numbers []int) int {
	sum := 0
	for idx, number := range numbers {
		// idx 索引 number 值
		fmt.Println(idx, number)
		sum += number
	}
	return sum
}

//func Sum(numbers [5]int) int {
//	sum := 0
//	for i := 0; i < len(numbers); i++ {
//		sum += numbers[i]
//	}
//	return sum
//}
