package main

import (
	"log"
	"time"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	low := 0
	high := len(arr) - 1

	target := 5

	log.Println("Linear Search")

	timeNowStart := time.Now()
	log.Println("Time Start: ", timeNowStart)
	for _, v := range arr {
		if v == target {
			timeNowEnd := time.Now()
			log.Println("Time Finish:", timeNowEnd)
			log.Println("Time Taken: ", timeNowEnd.Sub(timeNowStart))
			break
		}
	}

	log.Println("Binary Search")

	timeNowStart = time.Now()
	log.Println("Start Time: ", timeNowStart)
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			timeNowEnd := time.Now()
			log.Println("Time Finish:", timeNowEnd)
			log.Println("Time Taken: ", timeNowEnd.Sub(timeNowStart))
			break
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

}
