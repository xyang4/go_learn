package main

/*
   冒泡排序，由大到小
*/
func BubblingSort(arr []int) {

	num := len(arr)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if arr[i] < arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}
