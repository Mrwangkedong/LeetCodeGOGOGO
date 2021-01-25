package main

import (
	"fmt"
)

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/

//三重循环
func lengthOfLongestSubstring2(s string) int {
	maxlen := 0

	//利用双重循环abcabcbb
	for i := 0; i < len(s); i++ {
		l := 1
		for j := i + 1; j < len(s); j++ {
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					j = len(s) //在k跳出循环后，j也跳出循环
					break
				}
			}
			if j != len(s) {
				l++
			}
		}

		if maxlen < l {
			maxlen = l
		}

	}

	return maxlen
}

//利用map
func lengthOfLongestSubstring(s string) int {

	maxlen := 0

	//利用双重循环abcabcbb
	for i := 0; i < len(s); i++ {
		l := 1

		mapp := make(map[string]int) //string -- 元素  int -- 位置
		mapp[string(s[i])] = i       //存储

		for j := i + 1; j < len(s); j++ {
			//判断当前s[j]是否在map中
			if _, has := mapp[string(s[j])]; !has {
				//如果不存在,存下来
				mapp[string(s[j])] = j
				l++
			} else {
				//如果存在，就计算距离
				break
			}
		}

		if maxlen < l {
			maxlen = l
		}

	}

	return maxlen
}

func main() {
	maxlen := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(maxlen)
}
