package main

import "fmt"

func exchange(nums []int) []int {
	Num := len(nums)
	//定义两个游标，一个是指向前面的偶数，一个在后面找奇数
	i := 0
	//将j定位到第一个偶数得到下一位
	for i < len(nums) {
		if nums[i]%2 != 0 {
			i++
		} else {
			break
		}
	}
	j := i + 1
	//开始循环nums，进行交换
	for j < Num {
		//找到偶数
		for {
			if nums[i]%2 != 0 {
				i++
			} else {
				break
			}
		}
		//找到奇数
		for j < Num {
			if nums[j]%2 != 0 {
				break
			} else {
				j++
			}
		}
		if j == Num {
			break
		}
		//进行交换
		nums[i], nums[j] = nums[j], nums[i]

	}
	return nums
}
func main() {
	a := []int{1}
	a = exchange(a)
	fmt.Println(a)
}
