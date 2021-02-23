package main

import "fmt"

/**
 * 给定一棵二叉搜索树，请找出其中第k大的节点。
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//搜索二叉树，左小右大
func kthLargest(root *TreeNode, k int) int {

	//思路：进行中序遍历，不过是先右子树再root再左子树
	//1.创建栈
	var TreeStack [99]*TreeNode
	//2.top变量，，K大变量
	TopIndex := -1
	i := 0
	p := root
	//root入栈
	for TopIndex != -1 || p != nil {
		//如果p！=nil，p入栈
		if p != nil {
			TopIndex++
			TreeStack[TopIndex] = p
			p = p.Right
		} else {
			p = TreeStack[TopIndex]
			TopIndex--
			i++
			fmt.Println(i, p.Val)
			if i == k {
				return p.Val
			}
			p = p.Left
		}
	}

	return 0
}

func main() {
	var root TreeNode
	root.Val = 5
	root.Left = new(TreeNode)
	root.Left.Val = 3
	root.Right = new(TreeNode)
	root.Right.Val = 6
	root.Left.Left = new(TreeNode)
	root.Left.Left.Val = 2
	root.Left.Right = new(TreeNode)
	root.Left.Right.Val = 4
	root.Right.Left = nil
	root.Right.Right = nil

	fmt.Println(kthLargest(&root, 3))
}

//获得二叉树的length
func TreeLen(root *TreeNode) int {

	//第一次遍历，得出支点数
	//1.创建队列
	treeQueue := []*TreeNode{}
	//2.队首队尾index,front出，rear进
	front := -1
	rear := -1
	//3.创建树节点变量
	p := root
	//4.root入队
	treeQueue = append(treeQueue, root)
	rear++

	//5.进行循环遍历
	for p != nil {
		front++
		p = treeQueue[front] //队列出队
		//将左右子树入队
		if p.Left != nil {
			treeQueue = append(treeQueue, p.Left)
			rear++
		}
		if p.Right != nil {
			treeQueue = append(treeQueue, p.Right)
			rear++
		}
	}
	//树上元素个数 == 数组最后一个index+1
	lenTreeElement := rear + 1

	return lenTreeElement
}
