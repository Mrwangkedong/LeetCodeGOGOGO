package main

import (
	"fmt"
	"strings"
)

/*
题目描述：
	给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
例子：
	输入:	"tree"
	输出:	"eert"
例子2：
	输入：	"his s he a ha he  ha ae"
	输出：	”        hhhhhaaaaeeessi“
*/

func frequencySort(s string) string {
	//返回值
	sb := ""
	//创建一个26*2的数组,根据字符的ASCII编码进行存入，存入他出现的次数
	list := [123]int{}
	//记录出现的max_num
	max_num := 0
	//遍历
	for _, item := range s {
		//在其相对应的index上增加1
		list[int(item)]++
		num := list[int(item)]
		if max_num < num {
			max_num = num
		}
	}
	//从max_num开始向下循环
	for i := max_num; i > 0; i-- {
		for j := 0; j < 123; j++ {
			if list[j] == i {
				sb += strings.Repeat(string(rune(j)), i)
			}
		}
	}

	return sb
}

//48开始 122

func main() {
	fmt.Println(frequencySort("his s he a ha he  ha ae"))
}
