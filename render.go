package sdl

type RendererFlags uint32

const (
	RendererSoftware RendererFlags = 1 << iota
	RendererAccelerated
	RendererPresentVSync
	RendererTargetTexture
)

type TextureAccess uint32

const (
	TextureAccessStatic TextureAccess = iota
	TextureAccessStreaming
	TextureAccessTarget
)

type TextureModulate uint32

const (
	TextureModulateNone TextureModulate = iota
	TextureModulateColor
	TextureModulateAlpha
)

type RendererFlip uint32

const (
	FlipNone RendererFlip = iota
	FlipHorizontal
	FlipVertical
)
