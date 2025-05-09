package main

func BinarySearch(arr []int, x int) bool {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := (l + r) / 2

		if arr[mid] == x {
			return true
		} else if arr[mid] > x {
			r = mid - 1
		} else if arr[mid] < x {
			l = mid + 1
		}
	}

	return false
}

func main() {}
