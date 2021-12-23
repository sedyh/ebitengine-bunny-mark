package bench

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bunny struct {
	Sprite     *ebiten.Image
	Colorful   *bool
	Hue        float64
	Gravity    float64
	PosX, PosY float64
	VelX, VelY float64
}

func NewBunny(sprite *ebiten.Image, shift, hue float64, colorful *bool) *Bunny {
	return &Bunny{
		Sprite:   sprite,
		Colorful: colorful,
		Hue:      hue,
		PosX:     shift,
		Gravity:  0.00095,
		VelX:     RangeFloat(0, 0.005),
		VelY:     RangeFloat(0.0025, 0.005),
	}
}

func (b *Bunny) Update(bounds image.Rectangle) {
	b.PosX += b.VelX
	b.PosY += b.VelY
	b.VelY += b.Gravity

	sw, sh := float64(bounds.Dx()), float64(bounds.Dy())
	iw, ih := float64(b.Sprite.Bounds().Dx()), float64(b.Sprite.Bounds().Dy())
	relW, relH := iw/sw, ih/sh

	if b.PosX+relW > 1 {
		b.VelX *= -1
		b.PosX = 1 - relW
	}
	if b.PosX < 0 {
		b.VelX *= -1
		b.PosX = 0
	}
	if b.PosY+relH > 1 {
		b.VelY *= -0.85
		b.PosY = 1 - relH
		if Chance(0.5) {
			b.VelY -= RangeFloat(0, 0.009)
		}
	}
	if b.PosY < 0 {
		b.VelY = 0
		b.PosY = 0
	}
}

func (b *Bunny) Draw(screen *ebiten.Image) {
	sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.PosX*sw, b.PosY*sh)
	if *b.Colorful {
		op.ColorM.RotateHue(b.Hue)
	}
	screen.DrawImage(b.Sprite, op)
}
