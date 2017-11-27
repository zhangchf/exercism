package connect

import "errors"

var ErrInvalidBoard = errors.New("Invalid Board")

func ResultOf(board []string) (string, error) {
	boardHeight := len(board)
	if boardHeight == 0 {
		return "", ErrInvalidBoard
	}
	boardWidth := len(board[0])
	for _, line := range board {
		if len(line) != boardWidth {
			return "", ErrInvalidBoard
		}
	}

	// Check top player
	player := uint8('O')
	stack := Stack{}
	for i, j := 0, 0; j < boardWidth; j++ {
		if board[i][j] == player {
			stack.Push(Element{i, j})
		}
	}
	foundPlayer, found := checkPlayer(player, board, &stack, func(element Element) bool {
		if element.row == boardHeight-1 {
			return true
		}
		return false
	})
	if found {
		return foundPlayer, nil
	}

	// Check left player
	player = uint8('X')
	stack = Stack{}
	for i, j := 0, 0; i < boardHeight; i++ {
		if board[i][j] == player {
			stack.Push(Element{i, j})
		}
	}
	foundPlayer, found = checkPlayer(player, board, &stack, func(element Element) bool {
		if element.col == boardWidth-1 {
			return true
		}
		return false
	})
	if found {
		return foundPlayer, nil
	}

	return "", nil

}

func checkPlayer(player uint8, board []string, stack *Stack, foundFunc func(Element) bool) (string, bool) {
	boardHeight, boardWidth := len(board), len(board[0])

	boardMark := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		boardMark[i] = make([]bool, len(board[0]))
	}

	for stack.size > 0 {
		lastElem := stack.Pop()
		boardMark[lastElem.row][lastElem.col] = true
		if foundFunc(lastElem) {
			return string(player), true
		} else {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					row, col := lastElem.row+i, lastElem.col+j
					if !(i == 0 && j == 0) && row >= 0 && row < boardHeight && col >= 0 && col < boardWidth && board[row][col] == player && !boardMark[row][col] {
						stack.Push(Element{row, col})
					}
				}
			}
		}
	}
	return "", false

}

type Element struct {
	row, col int
}
type Stack struct {
	data []Element
	size int
}

func (stack *Stack) Push(elem Element) {
	stack.data = append(stack.data, elem)
	stack.size++
}

func (stack *Stack) Pop() Element {
	if stack.size > 0 {
		result := stack.data[len(stack.data)-1]
		stack.data = stack.data[:len(stack.data)-1]
		stack.size--
		return result
	}
	return Element{}
}
