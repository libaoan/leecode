package main

// hash全遍历， 速度72%， 内存76%
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if !check(board, i, j) {
				return false
			}
		}
	}
	return true
}

func check(board [][]byte, i, j int) bool {
	m := make(map[byte]bool)

	// check row
	for _, v := range board[i] {
		if _, ok := m[v]; ok {
			if v == '.' {
				continue
			}
			return false
		}
		m[v] = true
	}

	// check col
	m = make(map[byte]bool)
	for _, v := range board {
		if _, ok := m[v[j]]; ok {
			if v[j] == '.' {
				continue
			}
			return false
		}
		m[v[j]] = true
	}
	// check matrix
	switch {
	case i >= 0 && i < 3:
		i = 0
	case i >= 3 && i < 6:
		i = 3
	case i >= 6 && i < 9:
		i = 6
	default:
		return false
	}
	switch {
	case j >= 0 && j < 3:
		j = 0
	case j >= 3 && j < 6:
		j = 3
	case j >= 6 && j < 9:
		j = 6
	default:
		return false
	}
	m = make(map[byte]bool)
	for k := 0; k < 3; k++ {
		for h := 0; h < 3; h++ {
			if _, ok := m[board[i+k][j+h]]; ok {
				if board[i+k][j+h] == '.' {
					continue
				}
				return false
			}
			m[board[i+k][j+h]] = true
		}

	}
	return true
}
