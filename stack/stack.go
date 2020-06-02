package stack

//Given an array, print the Next Greater Element (NGE) for every element.
//The Next greater Element for an element x is the first greater element on the right side of x in array.
//Elements for which no greater element exist, consider next greater element as -1.
//For the input array [4, 5, 2, 25}, the next greater elements for each element are as follows.
//Element       NGE
//4      -->   5
//5      -->   25
//2      -->   25
//25     -->   -1
func NextGreaterElementNaive(arr []int) []int {
	soln := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		right := -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				right = arr[j]
				break
			}
		}
		soln[i] = right
	}

	return soln
}

func NextGreaterElement(arr []int) []int {
	stack := make([]int, 0)
	soln := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {

		for len(stack) > 0 && stack[len(stack)-1] <= arr[i] {
			stack = stack[:len(stack)-1] // pop
		}
		if len(stack) == 0 {
			soln[i] = -1
		} else {
			soln[i] = stack[len(stack)-1] //top
		}
		stack = append(stack, arr[i])
	}
	return soln
}

//Given an array of integers, find the closest (not considering distance, but value) greater on left of every element.
//If an element has no greater on the left side, print -1.
func PreviousGreaterElement(arr []int) []int {
	stack := make([]int, 0)
	sol := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] <= arr[i] {
			stack = stack[:len(stack)-1] // pop
		}
		if len(stack) == 0 {
			sol[i] = -1
		} else {
			sol[i] = stack[len(stack)-1] //top
		}
		stack = append(stack, arr[i])
	}
	return sol
}

//Given an array of integers, find the closest (not considering distance, but value)
//smaller on left of every element. If an element has no smaller on the left side, print -1.
func PreviousSmallerElement(arr []int) []int {
	stack := make([]int, 0)
	sol := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] >= arr[i] {
			stack = stack[:len(stack)-1] // pop
		}
		if len(stack) == 0 {
			sol[i] = -1
		} else {
			sol[i] = stack[len(stack)-1] //top
		}
		stack = append(stack, arr[i])
	}
	return sol
}

func PreviousSmallerIndex(arr []int) []int {
	stack := make([][2]int, 0)
	sol := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] <= stack[len(stack)-1][1] {
			stack = stack[:len(stack)-1] // pop
		}
		if len(stack) == 0 {
			sol[i] = -1
		} else {
			sol[i] = stack[len(stack)-1][0] //top
		}
		stack = append(stack, [2]int{i, arr[i]})
	}
	return sol
}

func NextSmallerElement(arr []int) []int {
	stack := make([]int, 0)
	sol := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {

		for len(stack) > 0 && stack[len(stack)-1] >= arr[i] {
			stack = stack[:len(stack)-1] // pop
		}
		if len(stack) == 0 {
			sol[i] = -1
		} else {
			sol[i] = stack[len(stack)-1] //top
		}
		stack = append(stack, arr[i])
	}
	return sol
}

func NextSmallerIndex(arr []int) []int {
	stack := make([][2]int, 0)
	sol := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {

		for len(stack) > 0 && stack[len(stack)-1][1] >= arr[i] {
			stack = stack[:len(stack)-1] // pop
		}
		if len(stack) == 0 {
			sol[i] = len(arr)
		} else {
			sol[i] = stack[len(stack)-1][0] //top
		}
		stack = append(stack, [2]int{i, arr[i]})
	}
	return sol
}

//The stock span problem is a financial problem where we have a series of n daily price quotes
//for a stock and we need to calculate span of stock’s price for all n days.
//The span Si of the stock’s price on a given day i is defined as the maximum number
//of consecutive days just before the given day, for which the price of the stock on
//the current day is less than or equal to its price on the given day.
//For example, if an array of 7 days prices is given as {100, 80, 60, 70, 60, 75, 85},
//then the span values for corresponding 7 days are {1, 1, 1, 2, 1, 4, 6}
func StockSpanProblem(arr []int) []int { //NGL
	ngLStack := make([][2]int, 0) // index,value
	sol := make([]int, len(arr))
	for i, v := range arr {
		// pop
		for len(ngLStack) > 0 && ngLStack[len(ngLStack)-1][1] <= v {
			ngLStack = ngLStack[:len(ngLStack)-1]
		}
		if len(ngLStack) == 0 { //empty
			sol[i] = i + 1
		} else {
			sol[i] = i - ngLStack[len(ngLStack)-1][0] //top
		}
		ngLStack = append(ngLStack, [2]int{i, v}) // push after popping all smaller
	}

	return sol
}

//Find the largest rectangular area possible in a given histogram where the largest rectangle
//can be made of a number of contiguous bars. For simplicity, assume that all bars
//have same width and the width is 1 unit.
//Maximum area of histogram
func MaximumAreaHistogram(arr []int) int { // max((NSR-NSL)*arr[i])
	nsL := PreviousSmallerIndex(arr)
	nsR := NextSmallerIndex(arr)
	max := 0
	for i, val := range arr {
		area := val * (nsR[i] - nsL[i] - 1)
		if area > max {
			max = area
		}
	}
	return max
}
