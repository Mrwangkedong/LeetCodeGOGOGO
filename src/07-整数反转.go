package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	sb := 0

	//将x从int->string
	xtr := strconv.Itoa(x)
	//flag == 0 表示为负数
	flag := 1
	if xtr[0:1] == "-" {
		xtr = xtr[1:]
		flag = 0
	}
	k := 1
	for j := 0; j < len(xtr); j++ {
		x, _ := strconv.ParseInt(xtr[j:j+1], 10, 64)
		sb = sb + int(x)*k
		k *= 10
	}
	if sb < (0-int(math.Pow(2, 31))) || sb > int(math.Pow(2, 31)-1) {
		return 0
	}

	if flag == 0 {
		return -sb
	} else {
		return sb
	}

}

func main() {
	sb := reverse(-2147483412)
	fmt.Println(sb)

}
