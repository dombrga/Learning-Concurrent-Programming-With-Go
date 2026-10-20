package main

import (
	"fmt"
	"sync"
	"time"
)

// The caller goroutine must hold the lock before calling Wait(). Otherwise, error sync: unlock of unlocked mutex.
func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int) {
	cond.L.Lock()
	fmt.Println(playerId, ": Connected")
	*playersRemaining--
	cond.L.Unlock()

	if *playersRemaining == 0 {
		cond.Broadcast()
	}

	// Wait until all players (goroutines) have connected.
	for *playersRemaining > 0 {
		fmt.Println(playerId, ": Waiting for more players")
		cond.Wait()
	}

	// cond.L.Unlock()
	fmt.Println("All players connected. Ready player", playerId)
	// Game started
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playersInGame := 4

	for playerId := 0; playerId < 4; playerId++ {
		go playerHandler(cond, &playersInGame, playerId)
		time.Sleep(time.Second * 1)
	}
}
