package game

// SparseArray

type Status int

const (
	Empty Status = iota
	White
	Black
)

type SparseArray struct {
	row, col int
	value    Status
}

type Board struct {
	value [][]SparseArray
	Size  int
}

func NewBoard(size int) *Board {
	board := make([][]SparseArray, size)
	for i := range board {
		board[i] = make([]SparseArray, size)
	}
	return &Board{
		Size:  size,
		value: board,
	}
}

func (b *Board) Get(row, col int) Status {
	if row >= 0 && row < b.Size && col >= 0 && col < b.Size {
		return b.value[row][col].value
	}
	return Empty
}

func (b *Board) Set(row, col int, value Status) {
	if row >= 0 && row < b.Size && col >= 0 && col < b.Size {
		b.value[row][col].value = value
	}
}
