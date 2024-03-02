package game

import "sync"

//todo: make the game room

type Game struct {
	Board       *Board
	RedPlayer   Player
	BlackPlayer Player
	mu          sync.Mutex
	Turn        int
}
//
// func NewGame() *Game {
// 	return &Game{
// 		Board: NewBoard(),
// 		Turn:  0,
// 	}
// }
//
// Player drop the chess
func (g *Game) DropChess(player Player, row, col int) bool {
	g.mu.Lock()
	defer g.mu.Unlock()
	//检查落子是否合法
	if row < 0 || row > g.Board.Size || col < 0 || col >= g.Board.Size {
		return false
	}
	//判断是否到该玩家落子
	if player == g.RedPlayer && g.Turn != 0 ||
		player == g.BlackPlayer && g.Turn != 0 {
		return false
	}

	//在棋盘上放置棋子

	// 判断是否有玩家胜利
	return false
}

//game winning check
