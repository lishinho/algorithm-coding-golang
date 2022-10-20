package leetcode

import (
	"strconv"
)

/*
输入: num1 = "2", num2 = "3"
输出: "6"
"498828660196"
"840477629533"
"419254329864656431168468"
*/

func multiply1(num1 string, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	return strconv.Itoa(n1 * n2)
}

func multiply2(num1 string, num2 string) string {
	num1Byte, num2Byte := []byte(num1), []byte(num2)
	var res []byte
	l := max(len(num1), len(num2))
	var plus int
	for i := l; i >= 0; i-- {
		var n1, n2 int
		var a, b, c, d int
		// operate num1Byte num2Byte
		if len(num1Byte) > i {
			a, _ = strconv.Atoi(string(num1Byte[i]))
			if i > 0 {
				c, _ = strconv.Atoi(string(num1Byte[i-1]))
			}
		} else {
			a = 1
			c = 1
		}
		if len(num2Byte) > i {
			b, _ = strconv.Atoi(string(num2Byte[i]))
			if i > 0 {
				d, _ = strconv.Atoi(string(num2Byte[i-1]))
			}
		} else {
			b = 1
			d = 1
		}
		if i == 0 {
			n1 = a
			n2 = b
		} else {
			n1 = a * d
			n2 = b * c
		}
		tmp := (n1 + n2 + plus) % 10
		plus = (n1 + n2 + plus) / 10
		res = append(res, []byte(strconv.Itoa(tmp))[0])
	}
	return string(res)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := "0"
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		curr := ""
		add := 0
		for j := n - 1; j > i; j-- {
			curr += "0"
		}
		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			product := x*y + add
			curr = strconv.Itoa(product%10) + curr
			add = product / 10
		}
		for ; add != 0; add /= 10 {
			curr = strconv.Itoa(add%10) + curr
		}
		ans = addStrings(ans, curr)
	}
	return ans
}

func addStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add := 0
	ans := ""
	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// 左-中-右
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}
