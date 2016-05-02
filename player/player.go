package player

import "github.com/lean-poker/poker-player-go/leanpoker"

type Player interface {
	
	// BetRequest is invoked when the Lean Poker sends the current game state to the Bot.
	BetRequest(state *leanpoker.Game) int
	
	// Showdown is invoked when a Showdown happens on the Lean Poker game server.
	Showdown(state *leanpoker.Game)
	
	// Version is invoked when the Lean Poker game server requests identity of the version of the Bot.
	Version() string
}