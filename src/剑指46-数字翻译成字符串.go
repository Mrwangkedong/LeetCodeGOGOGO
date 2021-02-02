package main

import (
	"fmt"
	"strconv"
)

//
/*
题目描述：
	给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。
一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

示例：
	输入: 12258
	输出: 5
	解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
*/

func translateNum(num int) int {
	//把int转化成字符串
	str := strconv.Itoa(num)
	//如果是空则返回0
	//如果是一位就返回1
	if len(str) <= 1 {
		return len(str)
	} else {
		left := 0
		right := 0
		//舍弃第一位，进行递归
		num1, _ := strconv.ParseInt(str[1:], 10, 32)
		left = translateNum(int(num1))
		//如果前两位0<=x && x<=25,舍弃前两位i，进行递归
		if item, _ := strconv.ParseInt(str[:2], 10, 32); int(item) <= 25 {
			num2, _ := strconv.ParseInt(str[2:], 10, 32)
			right = translateNum(int(num2))
		}
		return left + right
	}
}
func main() {
	fmt.Println(translateNum(12))
}
