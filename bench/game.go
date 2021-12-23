package bench

import (
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
)

type Game struct {
	Sprite   *ebiten.Image   // Image for bunnies
	Bounds   image.Rectangle // Physical window size
	Bunnies  []*Bunny        // List of bunnies
	Amount   *int            // How much to add
	Metrics  *Metrics        // Current TPS, FPS, object count and plots
	Colorful *bool           // Add some serious load
	Gpu      string          // Current gpu
}

func NewGame(amount int, colorful bool) *Game {
	g := &Game{
		Sprite:   LoadSprite(),
		Amount:   &amount,
		Colorful: &colorful,
		Metrics:  NewMetrics(500*time.Millisecond, &colorful, &amount),
	}
	g.AddBunnies()

	return g
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.AddBunnies()
	}

	if _, offset := ebiten.Wheel(); offset != 0 {
		*g.Amount += int(offset * 10)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		*g.Colorful = !*g.Colorful
	}

	for _, b := range g.Bunnies {
		b.Update(g.Bounds)
	}

	g.Metrics.Update(float64(len(g.Bunnies)))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Whitesmoke)

	for _, b := range g.Bunnies {
		b.Draw(screen)
	}

	g.Metrics.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.Bounds = image.Rect(0, 0, width, height)

	return width, height
}

func (g *Game) AddBunnies() {
	for i := 0; i < *g.Amount; i++ {
		b := NewBunny(
			g.Sprite,
			float64(len(g.Bunnies)%2),
			RangeFloat(0, 2*math.Pi),
			g.Colorful,
		)

		g.Bunnies = append(g.Bunnies, b)
	}
}
