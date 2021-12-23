package bench

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

type Metrics struct {
	Ticker   *time.Ticker
	Bounds   *image.Rectangle
	Colorful *bool
	Amount   *int
	Gpu      string
	Tps      *Plot
	Fps      *Plot
	Objects  *Plot
}

func NewMetrics(rate time.Duration, resolution *image.Rectangle, colorful *bool, amount *int) *Metrics {
	return &Metrics{
		Ticker:   time.NewTicker(rate),
		Gpu:      GpuInfo(),
		Tps:      NewPlot(20, 60),
		Fps:      NewPlot(20, 60),
		Objects:  NewPlot(20, 60000),
		Bounds:   resolution,
		Colorful: colorful,
		Amount:   amount,
	}
}

func (m *Metrics) Update(objects float64) {
	select {
	case <-m.Ticker.C:
		m.Objects.Update(objects)
		m.Tps.Update(ebiten.CurrentTPS())
		m.Fps.Update(ebiten.CurrentFPS())
	default:
	}
}

func (m *Metrics) Draw(screen *ebiten.Image) {
	str := fmt.Sprintf(
		"GPU: %s\nTPS: %.2f, FPS: %.2f, Objects: %.f\nColorful: %t, Amount: %d\nResolution: %dx%d",
		m.Gpu, m.Tps.Last(), m.Fps.Last(), m.Objects.Last(),
		*m.Colorful, *m.Amount,
		m.Bounds.Dx(), m.Bounds.Dy(),
	)

	rect := text.BoundString(basicfont.Face7x13, str)
	width, height := float64(rect.Dx()), float64(rect.Dy())

	padding := 20.0
	rectW, rectH := width+padding, height+padding
	plotW, plotH := 100.0, 40.0

	ebitenutil.DrawRect(screen, 0, 0, rectW, rectH, color.RGBA{A: 128})
	text.Draw(screen, str, basicfont.Face7x13, int(padding)/2, 10+int(padding)/2, colornames.White)

	m.Tps.Draw(screen, 0, padding+rectH, plotW, plotH)
	m.Fps.Draw(screen, 0, padding+rectH*2, plotW, plotH)
	m.Objects.Draw(screen, 0, padding+rectH*3, plotW, plotH)
}
