package main

import (
	"fmt"
	"testing"
)

func TestBubblingSort(t *testing.T) {
	sortArr := []int{1, 7, 4, 2, 5}
	fmt.Println(sortArr)
	BubblingSort(sortArr)
	fmt.Println(sortArr)
}
