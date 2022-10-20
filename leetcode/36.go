package leetcode

func isValidSudoku(board [][]byte) bool {
	var (
		row [9][10]int
		col [9][10]int
		box [9][10]int
	)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur := board[i][j] - '0'
			if row[i][cur] == 1 || col[j][cur] == 1 || box[j/3+(i/3)*3][cur] == 1 {
				return false
			}
			row[i][cur] = 1
			col[j][cur] = 1
			box[j/3+(i/3)*3][cur] = 1
		}
	}
	return true
}
