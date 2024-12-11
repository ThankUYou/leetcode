package binarytree

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144. 二叉树的前序遍历, easy
/*
 * 前序遍历：中 --》 左 --》右
 */
func preorderTraversal(root *TreeNode) (res []int) {
	var preTravel func(*TreeNode)
	preTravel = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preTravel(node.Left)
		preTravel(node.Right)
	}
	preTravel(root)
	return
}

// 94. 二叉树的中序遍历, easy
/*
 * 中序遍历：左 --》 中 --》右
 */
func inorderTraversal(root *TreeNode) (res []int) {
	var midTravel func(*TreeNode)
	midTravel = func(node *TreeNode) {
		if node == nil {
			return
		}
		midTravel(node.Left)
		res = append(res, node.Val)
		midTravel(node.Right)
	}

	midTravel(root)
	return
}

// 145. 二叉树的后序遍历, easy
/*
 * 后序遍历：左 --》右 --》 中
 */
func postorderTraversal(root *TreeNode) (res []int) {
	var lastTravel func(*TreeNode)
	lastTravel = func(node *TreeNode) {
		if node == nil {
			return
		}
		lastTravel(node.Left)
		lastTravel(node.Right)
		res = append(res, node.Val)
	}
	lastTravel(root)
	return
}

// 872. 叶子相似的树, easy
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var travel func(*TreeNode) []int
	travel = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}
		// 叶子结点添加
		if node.Left == nil && node.Right == nil {
			return []int{node.Val}
		}
		left := travel(node.Left)
		right := travel(node.Right)
		return append(left, right...)
	}

	res1 := travel(root1)
	res2 := travel(root2)

	return checkSame(res1, res2)
}

func checkSame(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// LCP 44. 开幕式焰火
// 多少种不同的颜色 ==》map
func numColor(root *TreeNode) int {
	cnt := make(map[int]int, 0)
	var travel func(*TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		cnt[node.Val]++
		travel(node.Left)
		travel(node.Right)
	}
	travel(root)

	return len(cnt)
}

// 404. 左叶子之和
// 左叶子怎么判断?当我们遍历到节点 node 时
// 如果它的左子节点是一个叶子结点，那么就将它的左子节点的值累加计入答案
func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func dfs(node *TreeNode) (ans int) {
	if node.Left != nil {
		if isLeafNode(node.Left) {
			ans += node.Left.Val
		} else {
			ans += dfs(node.Left)
		}
	}
	if node.Right != nil && !isLeafNode(node.Right) {
		ans += dfs(node.Right)
	}
	return
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}

// 671. 二叉树中第二小的节点
func findSecondMinimumValue(root *TreeNode) int {
	vals := []int{}
	var travel func(*TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		travel(node.Left)
		travel(node.Right)
	}
	travel(root)
	slices.Sort(vals)
	return min2(vals)
}

func min2(arr []int) int {
	minN := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > minN {
			return arr[i]
		}
	}
	return -1
}
