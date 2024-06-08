package main

import (
	_ "fmt"
	_ "io"
	"log"
	_ "math/rand"
	_ "strings"
	_ "time"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	_ "github.com/hajimehoshi/ebiten/v2/inpututil"
	_ "github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	title string = "Space_Defender"
)

var (
// fontSize        int = 32
// mplusNormalFont font.Face

// bkg    = color.Black
// white = color.White
// red    = color.RGBA{252, 69, 3, 255}
// yellow = color.RGBA{184, 245, 17, 255}
// green  = color.RGBA{37, 199, 22, 255}
// defeat = false
)

func main() {
	ebiten.SetWindowSize(screenHeight, screenWidth)
	ebiten.SetWindowTitle(title)
	updatebleGameObjects := []gameObjects{}
	updatebleGameObjects = append(updatebleGameObjects, player1)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
