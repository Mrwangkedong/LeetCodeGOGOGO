package abc

import "fmt"

type kkk struct {
	name string
}

var Name string = "Kiki"

//二叉树定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func D() {
	fmt.Println(ReName(Name))
}
