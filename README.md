poker-player-go
===============

Go client skeleton for Lean Poker. For more information visit: http://leanpoker.org 

## Usage instructions

### Ingredients

See [Deploying Go Apps on Heroku](https://devcenter.heroku.com/articles/deploying-go#prerequisites)

### Workflow

## Prerequisites:

Install Godep
```go get github.com/tools/godep```

## Create your project
## Get the lean-poker library 
```go get github.com/lean-poker/poker-player-go```
## Start using the library
```Go
package main

import (
    poker "github.com/lean-poker/poker-player-go"
    "github.com/lean-poker/poker-player-go/leanpoker"
)


type Player struct{}

func (p *Player) BetRequest(state *leanpoker.Game) int {
    return 0;
}

func (p *Player) Showdown(state *leanpoker.Game) {
    
}

func (p *Player) Version() string {
    return "Default folding Go player"
}

func main() {
    p := &Player{}
    
    poker.Start(p)
}
```

## Save dependencies in Godep. Mind the 3 dots!
```
Godep save ./...
```

## Push to Heroku!