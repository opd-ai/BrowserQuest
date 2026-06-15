package main

import (
	"log"

	"github.com/opd-ai/BrowserQuest/pkg/client/game"
)

func main() {
	if err := game.Run("BrowserQuest Go Client (Desktop)"); err != nil {
		log.Fatal(err)
	}
}
