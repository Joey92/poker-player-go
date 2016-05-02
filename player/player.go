package player

import "github.com/lean-poker/poker-player-go/leanpoker"

type Player interface {
	BetRequest(state *leanpoker.Game) int
	Showdown(state *leanpoker.Game)
	Version() string
}