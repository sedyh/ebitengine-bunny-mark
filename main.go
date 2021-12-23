package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/sedyh/ebiten-bunny-mark/bench"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	ebiten.SetWindowSize(800, 600)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(bench.NewGame(100, false)); err != nil {
		log.Fatal(err)
	}
}
