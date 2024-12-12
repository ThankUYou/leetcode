package binarytree

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 104. 二叉树的最大深度, easy
// DFS, 顶向下
func maxDepth(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		ans = max(ans, depth)
		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}
	dfs(root, 0)
	return
}

// 111. 二叉树的最小深度, easy
// DFS. 左右两子树最小值
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	minD := math.MaxInt
	if root.Left != nil {
		minD = min(minD, minDepth(root.Left))
	}
	if root.Right != nil {
		minD = min(minD, minDepth(root.Right))
	}

	return minD + 1
}

// 112. 路径总和, easy
// 我的思路是求出所有和存放数组中，然后检测是否和targetsum相等
func hasPathSum(root *TreeNode, targetSum int) bool {
	res := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			res = append(res, sum)
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, 0)

	for _, v := range res {
		if v == targetSum {
			return true
		}
	}
	return false
}

func hasPathSumII(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == root.Right { // root 是叶子
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

// 129. 求根节点到叶节点数字之和
func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum = sum*10 + node.Val
		// 叶子结点
		if node.Left == node.Right {
			res += sum
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}
	dfs(root, 0)
	return res
}

// 199. 二叉树的右视图
// 先遍历右结点再遍历左结点
func rightSideView(root *TreeNode) []int {
	res := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return res
}

// 1448. 统计二叉树中好节点的数目
// 保留路径上的最大值，当前值与最大值比较
func goodNodes(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, maxVal int) (res int) {
		if node == nil {
			return 0
		}
		// 等于的情况也算好结点
		if node.Val >= maxVal {
			res++
			maxVal = node.Val
		}
		res += dfs(node.Left, maxVal) + dfs(node.Right, maxVal)
		return
	}
	return dfs(root, math.MinInt32)
}

// 1457. 二叉树中的伪回文路径
// 伪回文 --》 位运算
func pseudoPalindromicPaths(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, mask int) int {
		if node == nil {
			return 0
		}
		mask ^= 1 << node.Val
		if node.Left == nil && node.Right == nil {
			if mask&(mask-1) == 0 {
				return 1
			}
			return 0
		}
		return dfs(node.Left, mask) + dfs(node.Right, mask)
	}
	return dfs(root, 0)
}

// 1026. 节点与其祖先之间的最大差值
// 记录每条路径上的最大和最小值，最大差值 = 最大值-最小值
func maxAncestorDiffI(root *TreeNode) (res int) {
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, minVal, maxVal int) {
		if node == nil {
			return
		}
		minVal = min(minVal, node.Val)
		maxVal = max(maxVal, node.Val)
		res = max(res, max(node.Val-minVal, maxVal-node.Val))
		dfs(node.Left, minVal, maxVal)
		dfs(node.Right, minVal, maxVal)
	}
	dfs(root, root.Val, root.Val)
	return
}

// 1026. 节点与其祖先之间的最大差值优化
// 记录了最大最小值，当这条路径结束了就可以更新最大差值了，不需要每个结点都更新差值再比较
func maxAncestorDiffII(root *TreeNode) (res int) {
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, minVal, maxVal int) {
		if node == nil {
			res = max(res, maxVal-minVal)
			return
		}
		minVal = min(minVal, node.Val)
		maxVal = max(maxVal, node.Val)
		dfs(node.Left, minVal, maxVal)
		dfs(node.Right, minVal, maxVal)
	}
	dfs(root, root.Val, root.Val)
	return
}

// 1022. 从根到叶的二进制数之和
// 遍历所有的路径，每条路径的值求和
func sumRootToLeaf(root *TreeNode) int {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, res int) int {
		if node == nil {
			return 0
		}
		res = res*2 + node.Val
		// 到叶子结点了，这条路径结束了，路径上的值求和完毕
		if node.Left == nil && node.Right == nil {
			return res
		}
		return dfs(node.Left, res) + dfs(node.Right, res)
	}

	return dfs(root, 0)
}

// 623. 在二叉树中增加一行
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{
			Val:   val,
			Left:  root,
			Right: nil,
		}
	}
	if depth == 2 {
		root.Left = &TreeNode{val, root.Left, nil}
		root.Right = &TreeNode{val, nil, root.Right}
	} else {
		root.Left = addOneRow(root.Left, val, depth-1)
		root.Right = addOneRow(root.Right, val, depth-1)
	}
	return root
}

// 1315. 祖父节点值为偶数的节点和
// 暴力dfs
func sumEvenGrandparent(root *TreeNode) int {
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		res := 0
		if node.Val%2 == 0 {
			if node.Left != nil {
				if node.Left.Left != nil {
					res += node.Left.Left.Val
				}
				if node.Left.Right != nil {
					res += node.Left.Right.Val
				}
			}
			if node.Right != nil {
				if node.Right.Left != nil {
					res += node.Right.Left.Val
				}
				if node.Right.Right != nil {
					res += node.Right.Right.Val
				}
			}
		}
		return res + dfs(node.Left) + dfs(node.Right)
	}
	return dfs(root)
}

// 988. 从叶结点开始的最小字符串
// 遍历所有结果，最后排序
func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var dfs func(*TreeNode) []string
	dfs = func(tn *TreeNode) []string {
		if tn == nil {
			return []string{}
		}
		s := string(tn.Val + 'a')
		if tn.Left == nil && tn.Right == nil {
			return []string{s}
		}
		l := dfs(tn.Left)
		r := dfs(tn.Right)
		var res []string
		for i := range l {
			res = append(res, l[i]+s)
		}
		for i := range r {
			res = append(res, r[i]+s)
		}
		return res
	}
	res := dfs(root)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	return res[0]
}
