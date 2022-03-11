package main

/**
冒泡排序
*/
func bubbleSort(arr []*Person) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j].fatRate > arr[j+1].fatRate {
				arr[j].fatRate, arr[j+1].fatRate = arr[j+1].fatRate, arr[j].fatRate
			}
		}
	}
}

/**
快速排序
*/
func quickSort(arr []*Person, start, end int) {
	middle := (start + end) / 2
	midFatRate := arr[middle].fatRate
	s, e := partation(arr, start, end, midFatRate)
	if s == e {
		s++
		e--
	}
	if e > start {
		quickSort(arr, start, e)
	}
	if s < end {
		quickSort(arr, s, end)
	}
}

func partation(arr []*Person, start int, end int, midFatRate float64) (int, int) {
	for start <= end {
		for arr[start].fatRate < midFatRate {
			start++
		}
		for arr[end].fatRate > midFatRate {
			end--
		}
		if start >= end {
			break
		}

		arr[start].fatRate, arr[end].fatRate = arr[end].fatRate, arr[start].fatRate
		start++
		end--
	}
	return start, end
}
