package compat

import (
	"image/color"
)

type RGB565 struct {
	RGB uint16
}

func (c RGB565) RGBA() (r, g, b, a uint32) {
	r = uint32(c.RGB>>11) & 0x1f
	r |= r << 8
	g = uint32(c.RGB>>5) & 0x3f
	g |= g << 8
	b = uint32(c.RGB) & 0x1f
	b |= b << 8
	return r, g, b, 0xffff
}

var (
	RGB565Model      color.Model = color.ModelFunc(rgb565Model)
	ARGB2101010Model color.Model = color.ModelFunc(argb2101010Model)
)

func rgb565Model(c color.Color) color.Color {
	if _, ok := c.(RGB565); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0xffff {
		return RGB565{uint16(r>>11)<<11 | uint16(g>>10)<<5 | uint16(b>>11)}
	}
	if a == 0 {
		return RGB565{0}
	}
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	return RGB565{uint16(r>>11)<<11 | uint16(g>>10)<<5 | uint16(b>>11)}
}

type ARGB2101010 struct {
	ARGB uint32
}

func (c ARGB2101010) RGBA() (r, g, b, a uint32) {
	a = c.ARGB & 0xC0000000
	a >>= 30
	a *= 0x55
	r = c.ARGB & 0x3FF00000
	r = r>>14 | r>>24
	r *= a
	r /= 0xFF
	g = c.ARGB & 0x000FFC00
	g = g>>4 | g>>14
	g *= a
	g /= 0xFF
	b = c.ARGB & 0x000003FF
	b = b<<6 | b>>4
	b *= a
	b /= 0xFF
	a |= a << 8
	return
}

func argb2101010Model(c color.Color) color.Color {
	if _, ok := c.(ARGB2101010); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a == 0xffff {
		return ARGB2101010{0x3<<30 | (r>>6)<<20 | (g>>6)<<10 | (b>>6)<<0}
	}
	if a == 0 {
		return ARGB2101010{0}
	}
	r = (r * 0xffff) / a
	g = (g * 0xffff) / a
	b = (b * 0xffff) / a
	return ARGB2101010{(a>>14)<<30 | (r>>6)<<20 | (g>>6)<<10 | (b>>6)<<0}
}
