package game

import (
	"testing"
)

func TestBoard(t *testing.T) {
	// 创建一个新的棋盘
	board := NewBoard(10)

	// 在棋盘上放置一些棋子
	board.Set(2, 3, 2)
	board.Set(4, 5, 1)

	// 测试获取放置的棋子
	if board.Get(2, 3) != 2 {
		t.Errorf("Expected value at (2, 3) to be Black, got %d", board.Get(2, 3))
	}

	if board.Get(4, 5) != 1 {
		t.Errorf("Expected value at (4, 5) to be White, got %d", board.Get(4, 5))
	}

	// 测试获取未放置棋子的位置
	if board.Get(0, 0) != 0 {
		t.Errorf("Expected value at (0, 0) to be Empty, got %d", board.Get(0, 0))
	}

	// 测试边界情况
	if board.Get(-1, 0) != 0 {
		t.Errorf("Expected value at (-1, 0) to be Empty, got %d", board.Get(-1, 0))
	}

	if board.Get(11, 11) != 0 {
		t.Errorf("Expected value at (11, 11) to be Empty, got %d", board.Get(11, 11))
	}
}

