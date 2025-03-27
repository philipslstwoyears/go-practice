package main

import "fmt"

func binarySearch(slice []int, target int) error {
	left, right := 0, len(slice)-1
	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == target {
			fmt.Printf("Значение: %d находиться в слайсе на %d месте", target, mid)
			break
		} else if slice[mid] < target {
			left = mid + 1
		} else if slice[mid] > target {
			right = mid - 1
		}
	}
	return nil
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 9
	fmt.Println(binarySearch(slice, target))
}
