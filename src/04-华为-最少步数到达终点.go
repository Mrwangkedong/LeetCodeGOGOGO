package main

import (
	"fmt"
)

func main() {
	a := 0
	list := []int{}
	for true {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			list = append(list, a)
		}
	}
	step_num := len(list) //步数初始
	for i := 1; i < len(list)/2; i++ {
		next_step := i //第一步
		x := 1         //步数计数
		for true {
			x++                                     //进入下一步
			next_step = next_step + list[next_step] //计算下一步
			if next_step >= len(list) {
				break
			} else if next_step == len(list)-1 {
				if x < step_num {
					step_num = x
				}
				break
			}
		}
	}
	if step_num == len(list) {
		fmt.Println(-1)
	} else {
		fmt.Println(step_num)
	}

}
