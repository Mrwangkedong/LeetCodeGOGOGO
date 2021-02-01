package main

import (
	"fmt"
	"strconv"
)

func main() {
	//正整数N
	num := 0
	_, _ = fmt.Scan(&num)
	//第一个输入数
	s1 := ""
	_, _ = fmt.Scan(&s1)
	//第二个输入数
	s2 := ""
	_, _ = fmt.Scan(&s2)
	//将s1,s2转移到数组b1,b2中
	b1 := []int64{}
	b2 := []int64{}
	//判断原始结果
	origin := []int{}
	for i := 0; i < num; i++ {
		x, _ := strconv.ParseInt(string(s1[i]), 10, 64)
		b1 = append(b1, x)
		y, _ := strconv.ParseInt(string(s2[i]), 10, 64)
		b2 = append(b2, y)
		if x == 0 && y == 0 {
			origin = append(origin, 0)
		} else {
			origin = append(origin, 1)
		}
	}

	//设置故障数量 sb
	sb := 0
	//利用for循环模拟故障
	for j := 0; j < num-1; j++ {
		for k := j + 1; k < num; k++ {
			//进行交换
			b1[j], b1[k] = b1[k], b1[j]
			//观察改变j后的数组是否与原来数组相同
			if b1[j] == 0 && b2[j] == 0 {
				if origin[j] != 0 {
					sb++
				}
				continue
			} else {
				if origin[j] != 1 {
					sb++
				}
			}
			//观察改变k后的数组是否与原来数组相同
			if b1[k] == 0 && b2[k] == 0 {
				if origin[k] != 0 {
					sb++
				}
				continue
			} else {
				if origin[k] != 1 {
					sb++
				}
			}
			//为了下一次循环，j,k再交换回来
			b1[j], b1[k] = b1[k], b1[j]
		}
	}

	fmt.Println(sb)

}
