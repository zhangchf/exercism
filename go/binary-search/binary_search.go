package binarysearch

import "fmt"

func SearchInts(slice []int, key int) (index int) {
	start, end := 0, len(slice)-1
	mid := (start + end) / 2
	for {
		if start >= end {
			if end < 0 || slice[end] < key {
				end++
			}
			return end
		}
		mid = (start + end) / 2
		midValue := slice[mid]
		if midValue == key {
			for mid >= 0 && slice[mid] == key {
				mid--
			}
			return mid + 1
		} else if midValue > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0
}

func Message(slice []int, key int) (msg string) {
	if len(slice) == 0 {
		return "slice has no values"
	}
	index := SearchInts(slice, key)
	if index == len(slice) {
		return fmt.Sprintf("%d > all %d values", key, len(slice))
	}
	foundValue := slice[index]
	if index == 0 {
		if foundValue == key {
			return fmt.Sprintf("%d found at beginning of slice", key)
		}
		return fmt.Sprintf("%d < all values", key)
	}
	if foundValue == key {
		if index == len(slice)-1 {
			return fmt.Sprintf("%d found at end of slice", key)
		}
		return fmt.Sprintf("%d found at index %d", key, index)
	}
	return fmt.Sprintf("%d > %d at index %d, < %d at index %d", key, slice[index-1], index-1, slice[index], index)
}
