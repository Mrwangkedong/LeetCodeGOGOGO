package main

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
	TreeStack := []*TreeNode{}
	//2.top变量，K大变量
	TopIndex := -1
	i := 0
	p := root
	//3.root先进栈
	TreeStack = append(TreeStack, root)
	TopIndex++

	//4.进行循环遍历
	for p != nil {
		//取出栈顶元素
		p = TreeStack[TopIndex]
		//看看有没有右子树，如果有添加进去栈，如果没有，遍历当前,i++,添加左子树进栈
		if p.Right != nil {
			TreeStack = append(TreeStack, p.Right)
			TopIndex++
		} else {
			i++
			TopIndex--
			//对比i是否等于k,如果等于，返回当前节点的val值
			if i == k {
				return p.Val
			}
			//添加左子树
			if p.Left != nil {
				TreeStack = append(TreeStack, p.Left)
				TopIndex++
			}
		}
	}

	return 0
}

func main() {

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
