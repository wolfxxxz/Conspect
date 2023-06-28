package main

import "fmt"

func main() {
	strthird := []string{"dss", "sdsds", "sdsds"}
	slise := []int{1, 2, 3, 4, 5, 6}
	UpdateID(slise, len((strthird)))
}

func UpdateID(arr []int, lenS int) {
	for i := 0; i <= len(arr)-1; i++ {
		arr[i] = lenS + 1
		lenS++
	}
	fmt.Println(arr)
	/*
		for i, _ := range *NewWords {
			(*NewWords)[i].ID = len
			len++
		}*/
}
