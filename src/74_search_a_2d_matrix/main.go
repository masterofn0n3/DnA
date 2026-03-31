package main

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := (left + right) / 2
		current := matrix[mid]
		if exist := binarySearch(current, target); exist {
			return true
		} else {
			if target < current[0] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}

func binarySearch(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}
	return false
}
