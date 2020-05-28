package binarySearch

// Binary Search  Ascending order
func SearchAsc(arr []int, key int) int {

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2
		if key == arr[mid] {
			return mid
		} else if key > arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// Binary Search  Desc order
func SearchDesc(arr []int, key int) int {

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2
		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

// FirstOccurance in asc sorted array
func FirstOccurance(arr []int, key int) int {

	start := 0
	end := len(arr) - 1
	result := -1
	for start <= end {
		mid := start + (end-start)/2
		if key > arr[mid] {
			start = mid + 1
		} else {
			result = mid
			end = mid - 1
		}
	}
	return result
}

//Find the Rotation Count in Rotated Sorted array
//Consider an array of distinct numbers sorted in increasing order.
//The array has been rotated (clockwise) k number of times. Given such an array, find the value of k.
func RotationCount(arr []int) int {
	start := 0

	end := len(arr) - 1

	//sorted case
	if arr[end] > arr[start] || end == 0 {
		return 0
	}

	for start <= end {
		mid := start + (end-start)/2
		pivot := arr[mid]
		if pivot < arr[nextIndex(arr, mid)] && pivot < arr[prevIndex(arr, mid)] {
			return len(arr) - mid
		} else if arr[end] > pivot { // go where not sorted
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0
}

func nextIndex(arr []int, i int) int {
	return (i + 1) % len(arr)
}

func prevIndex(arr []int, i int) int {
	return (i + len(arr) - 1) % len(arr)
}

//Rotated Sorted Array Search
//Asked in: Facebook Google Microsoft Amazon
//Given an array of integers A of size N and an integer B.
//array A is rotated at some pivot unknown to you beforehand.
//(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2 ).
//You are given a target value B to search. If found in the array, return its index, otherwise return -1.
func RotatedArrSearch(arr []int, key int) int {
	r := RotationCount(arr)
	if r == 0 {
		return SearchAsc(arr, key)
	}

	a := SearchAsc(arr[:len(arr)-r], key) //search in left arr
	b := SearchAsc(arr[len(arr)-r:], key) //search in right arr

	if a == -1 && b == -1 {
		return -1
	}
	if a == -1 {
		return len(arr) - r + b
	}
	return a
}

//SEARCH IN A NEARLY SORTED ARRAY:
//Given an array which is sorted, but after sorting some elements are moved to either of the adjacent positions,
//i.e., arr[i] may be present at arr[i+1] or arr[i-1]. Write an efficient function to search an element in this array.
//Basically the element arr[i] can only be swapped with either arr[i+1] or arr[i-1].
//For example consider the array {2, 3, 10, 4, 40}, 4 is moved to next position and 10 is moved to previous position.
//Example :
//Input: arr[] =  {10, 3, 40, 20, 50, 80, 70}, key = 40
//Output: 2
//Output is index of 40 in given array
func NearlySortedArrSearch(arr []int, key int) int {

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2
		prev := -1
		next := -1
		if mid < len(arr)-1 {
			next = arr[mid+1]
		}
		if mid > 0 {
			prev = arr[mid-1]
		}
		if arr[mid] == key {
			return mid
		} else if prev == key {
			return mid - 1
		} else if next == key {
			return mid + 1
		} else if prev == -1 {
			return -1
		} else if next == -1 {
			return -1
		} else if minThreeInts(prev, arr[mid], next) > key { //search left
			end = mid - 2
		} else {
			start = mid + 2 //search right
		}
	}
	return -1
}

func minThreeInts(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= c && b <= a {
		return b
	}
	return c
}

//FIND FLOOR OF AN ELEMENT IN A SORTED ARRAY:
//Given a sorted array and a value x, the floor of x is the largest element in array smaller than or equal to x.
//Write efficient functions to find floor of x.
//Example:
//Input : arr[] = {1, 2, 8, 10, 10, 12, 19}, x = 5
//Output : 2
//2 is the largest element in arr[] smaller than 5.
func FloorInSortedArr(arr []int, key int) int {
	start := 0
	end := len(arr) - 1
	result := -1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == key {
			return key
		} else if arr[mid] > key {
			end = mid - 1 //search left
		} else {
			result = arr[mid]
			start = mid + 1
		}
	}
	return result
}
