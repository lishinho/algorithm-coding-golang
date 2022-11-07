package main

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func makeTree(values ...interface{}) *TreeNode {
	// Construct the nodes
	var nodes []*TreeNode
	for _, value := range values {
		switch v := value.(type) {
		case int:
			nodes = append(nodes, &TreeNode{Val: v})
		default:
			nodes = append(nodes, nil)
		}
	}

	// Arrange into a binary tree
	if count := len(nodes); count > 0 {
		for parentIndex, node := range nodes {
			if childIndex := (parentIndex * 2) + 1; childIndex < count {
				node.Left = nodes[childIndex]
				if childIndex++; childIndex < count {
					node.Right = nodes[childIndex]
				}
			}
		}
		return nodes[0]
	}
	return nil
}

func GetLevelOrderTreeNodeVal(root *TreeNode) [][]int {
	return levelOrder1(root)
}

// 解法一 BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		res = append(res, tmp)
	}
	return res
}

// 解法二 DFS
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	var dfsLevel func(node *TreeNode, level int)
	dfsLevel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		dfsLevel(node.Left, level+1)
		dfsLevel(node.Right, level+1)
	}
	dfsLevel(root, 0)
	return res
}
