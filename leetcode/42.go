package leetcode

//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

func trap4(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

func trap(height []int) (ans int) {
	var stack []int
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			curWidth := i - left - 1
			curHeight := min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}

func trap3(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
}

func trap2(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	var res int
	for i := 1; i < len(height); i++ {
		if height[i] < height[i-1] {
			bonus, right := calBonus(height, i)
			res += bonus
			i = right
		}
	}
	return res
}

func calBonus(height []int, left int) (int, int) {
	var right int
	i := left
	for i < len(height) && height[i] <= height[i-1] {
		i++
	}

	for i < len(height)-1 && height[i] <= height[i+1] {
		i++
	}
	right = i
	var bonus int
	if i >= len(height)-1 {
		return bonus, right
	}
	tall := min(height[left-1], height[right])

	for i = left; i < right; i++ {
		t := tall - height[i]
		if t > 0 {
			bonus += t
		}
	}
	return bonus, right
}
