package main

type SlideWindow struct {
	// 滑动窗口不需要记住窗口内的元素,通过map里的个数判断是否满足条件即可
	cnt map[int]int
	sum int
}

func InitWindow(nums []int) *SlideWindow {
	res := &SlideWindow{cnt: map[int]int{}}
	for _, v := range nums {
		res.cnt[v]++
		res.sum += v
	}
	return res
}

func (w *SlideWindow) Add(num int) {
	w.cnt[num]++
	w.sum += num
}

func (w *SlideWindow) Del(num int) {
	w.cnt[num]--
	if w.cnt[num] == 0 {
		delete(w.cnt, num)
	}
	w.sum -= num
}

func maximumSubarraySum(nums []int, k int) int64 {
	ans := 0
	window := InitWindow(nums[:k-1])
	for i := k - 1; i < len(nums); i++ {
		// 移入元素
		window.Add(nums[i])
		// 判断满足条件并更新ans
		if len(window.cnt) == k && ans < window.sum {
			ans = window.sum
		}
		// 移出元素
		window.Del(nums[i-k+1])
	}
	return int64(ans)
}
