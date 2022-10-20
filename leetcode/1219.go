package leetcode

/*
输入：grid = [[0,6,0],[5,8,7],[0,9,0]]
输出：24
解释：
[[0,6,0],
 [5,8,7],
 [0,9,0]]
一种收集最多黄金的路线是：9 -> 8 -> 7。
*/

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getMaximumGold(grid [][]int) (ans int) {
	var dfs func(x, y, gold int)
	dfs = func(x, y, gold int) {
		gold += grid[x][y]
		if gold > ans {
			ans = gold
		}

		rec := grid[x][y]
		// 挖光grid不参与回溯
		grid[x][y] = 0
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[nx]) && grid[nx][ny] > 0 {
				dfs(nx, ny, gold)
			}
		}
		grid[x][y] = rec
	}

	for i, row := range grid {
		for j, gold := range row {
			if gold > 0 {
				// 递归计算ans
				dfs(i, j, 0)
			}
		}
	}
	return
}
