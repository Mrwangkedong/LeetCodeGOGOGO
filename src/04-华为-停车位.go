package main

import (
	"fmt"
	"strconv"
	"strings"
)

//停车位

func main() {
	a := ""
	//定义一个数组，存放所有的车位
	list := []string{}
	//定义一个数组，存储所有value == 1的index
	one_list := []int{}
	//输入
	_, _ = fmt.Scan(&a)
	//存入当前停车场的车辆情况
	list = strings.Split(a, ",")
	if len(list) > 100 {
		list = list[:100]
	}
	for i := 0; i < len(list); i++ {
		//将value==1的index存入one_list
		if list[i] == strconv.Itoa(1) {
			one_list = append(one_list, i)
		}
	}

	//进行判断
	max_num := 0 //两个1之间0的个数
	for j := 1; j < len(one_list); j++ {
		k := one_list[j] - one_list[j-1]
		if k > max_num {
			max_num = k
		}
	}

	max_num--
	fmt.Println((max_num + 1) / 2)

}
