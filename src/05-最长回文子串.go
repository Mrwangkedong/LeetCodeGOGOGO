package main

/*
给你一个字符串 s，找到 s 中最长的回文子串。

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。

输入：s = "cbbd"
输出："bb"

*/

func longestPalindrome(s string) string {
	result := string(s[0])

	//设置一个黑名单map，包含某个不是回文子串的范围，若小子串在大子串里面，则大子串肯定不是回文子串
	//black_map[1] = 3   遇到  black_map[1] = 5  则进行更新为black_map[1] = 5
	black_map := make(map[int]int)

	//设置一个map，key为a-z，value为他们出现的位置的int数组
	letter_map := make(map[string][]int)

	//遍历字符串，进行letter_map的填充
	//在填充的时候，可以进行判断是否为回文子串！！！
	for index, i := range s {
		//判断是否存在，如果不存在就新建
		if _, has := letter_map[string(i)]; !has {
			letter_map[string(i)] = []int{index}
		} else {
			//如果存在，开始进行判断，是否已前一个是回环
			exist := 0                                                        //0表示不存在，1表示存在
			last_index := letter_map[string(i)][len(letter_map[string(i)])-1] //表示上一次出现的位置
			//先判断是否存在于黑名单中的子串属于他
			for k, v := range black_map {
				//存在  [last_index k v index]
				if k >= last_index && v <= index {
					exist = 1 //表示存在
					break     //跳出
				}
			}
			//如果存在小回环，且如果black_map[index-1]的值小于index，则扩大范围;如果不存在，就创建
			if exist == 1 {
				if _, has := black_map[last_index]; has && (black_map[last_index] < index) || (!has) {
					black_map[last_index] = index
				}
			} else {
				//如果不存在小回环，就自行进行判断
				for j := 1; j < (index-letter_map[string(i)][last_index])/2; j++ {
					//如果有一处不相对称，就说明不存在回环
					if s[last_index+j] != s[index-j] {
						black_map[last_index] = index
						break
					}
				}

			}

		}
	}

	return string(result)
}

func main() {
	longestPalindrome("asd")
}
