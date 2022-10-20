package leetcode

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

func permute2(nums []int) (res [][]int) {
	//回溯backtrack
	if len(nums) == 0 {
		return res
	}
	length := len(nums)
	set := make([]int, 0, length)
	var backtack func(i int)
	backtack = func(i int) {
		if i >= length {
			res = append(res, append([]int(nil), set...))
			return
		}
		set = append(set, nums[i])
		for j := i; j < length; j++ {
			nums[i], nums[j] = nums[j], nums[i]
			backtack(i + 1)
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	backtack(0)
	return res
}

func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	var backtrace func(int)
	backtrace = func(start int) {
		if start == len(nums)-1 {
			ret = append(ret, append([]int(nil), nums...))
			return
		}
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backtrace(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}
	}

	backtrace(0)

	return ret
}

func permute23(nums []int) [][]int {
	ret := make([][]int, 0)
	var perm []int
	var backtrace func(int)
	backtrace = func(start int) {
		if start == len(nums)-1 {
			ret = append(ret, append([]int(nil), perm...))
			return
		}
		for _, v := range nums {
			perm = append(perm, v)
			backtrace(start + 1)
			perm = perm[:len(perm)-1]
		}
	}

	backtrace(0)

	return ret
}

/*
class Solution {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList<List<Integer>>();

        List<Integer> output = new ArrayList<Integer>();
        for (int num : nums) {
            output.add(num);
        }

        int n = nums.length;
        backtrack(n, output, res, 0);
        return res;
    }

    public void backtrack(int n, List<Integer> output, List<List<Integer>> res, int first) {
        // 所有数都填完了
        if (first == n) {
            res.add(new ArrayList<Integer>(output));
        }
        for (int i = first; i < n; i++) {
            // 动态维护数组
            Collections.swap(output, first, i);
            // 继续递归填下一个数
            backtrack(n, output, res, first + 1);
            // 撤销操作
            Collections.swap(output, first, i);
        }
    }
}
*/
