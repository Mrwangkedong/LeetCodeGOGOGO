package main

import "fmt"

func duplicateZeros(arr []int) {

	//建立两个游标，一个原数组，一个
	i := 0 //跟随不变的数组  brr
	j := 0 //跟随要改变的数组  arr
	brr := []int{}

	for _, item := range arr {
		brr = append(brr, item)
	}
	fmt.Println(brr)

	for j < len(arr) {
		if brr[i] == 0 {
			arr[j] = 0
			j++
			if j >= len(arr) {
				break
			}
			arr[j] = 0
			j++

			i++
		} else {
			arr[j] = brr[i]
			j++
			i++
		}
	}

	fmt.Println(arr)

}

func main() {
	a := []int{1, 2, 3}

	duplicateZeros(a)

}
