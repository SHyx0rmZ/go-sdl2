package ttf

//#include <SDL2/SDL_ttf.h>
import "C"
import "unsafe"

type FontStyle int

const (
	StyleNormal        FontStyle = C.TTF_STYLE_NORMAL
	StyleBold          FontStyle = C.TTF_STYLE_BOLD
	StyleItalic        FontStyle = C.TTF_STYLE_ITALIC
	StyleUnderline     FontStyle = C.TTF_STYLE_UNDERLINE
	StyleStrikethrough FontStyle = C.TTF_STYLE_STRIKETHROUGH
)

type FontHinting int

const (
	HintingNormal FontHinting = C.TTF_HINTING_NORMAL
	HintingLight  FontHinting = C.TTF_HINTING_LIGHT
	HintingMono   FontHinting = C.TTF_HINTING_MONO
	HintingNone   FontHinting = C.TTF_HINTING_NONE
)

func (f *Font) Style() FontStyle {
	return (FontStyle)(C.TTF_GetFontStyle((*C.struct__TTF_Font)(f)))
}

func (f *Font) WithStyle(style FontStyle) {
	C.TTF_SetFontStyle((*C.struct__TTF_Font)(f), C.int(style))
}

func (f *Font) Outline() int {
	return int(C.TTF_GetFontOutline((*C.struct__TTF_Font)(f)))
}

func (f *Font) WithOutline(outline int) {
	C.TTF_SetFontOutline((*C.struct__TTF_Font)(f), C.int(outline))
}

func (f *Font) Hinting() FontHinting {
	return (FontHinting)(C.TTF_GetFontHinting((*C.struct__TTF_Font)(f)))
}

func (f *Font) WithHinting(hinting FontHinting) {
	C.TTF_SetFontHinting((*C.struct__TTF_Font)(f), C.int(hinting))
}

func (f *Font) Kerning() int {
	return int(C.TTF_GetFontKerning((*C.struct__TTF_Font)(f)))
}

func (f *Font) WithKerning(kerning int) {
	C.TTF_SetFontKerning((*C.struct__TTF_Font)(f), C.int(kerning))
}

func (f *Font) Height() int {
	return int(C.TTF_FontHeight((*C.struct__TTF_Font)(f)))
}

func (f *Font) Ascent() int {
	return int(C.TTF_FontAscent((*C.struct__TTF_Font)(f)))
}

func (f *Font) Descent() int {
	return int(C.TTF_FontDescent((*C.struct__TTF_Font)(f)))
}

func (f *Font) LineSkip() int {
	return int(C.TTF_FontLineSkip((*C.struct__TTF_Font)(f)))
}

func (f *Font) Faces() int {
	return int(C.TTF_FontFaces((*C.struct__TTF_Font)(f)))
}

func (f *Font) FaceIsFixedWidth() bool {
	return C.TTF_FontFaceIsFixedWidth((*C.struct__TTF_Font)(f)) != 0
}

func (f *Font) FaceFamilyName() string {
	nativeName := C.TTF_FontFaceFamilyName((*C.struct__TTF_Font)(f))
	if nativeName == nil {
		return ""
	}
	return C.GoString(nativeName)
}

func (f *Font) FaceStyleName() string {
	nativeName := C.TTF_FontFaceStyleName((*C.struct__TTF_Font)(f))
	if nativeName == nil {
		return ""
	}
	return C.GoString(nativeName)
}

func (f *Font) GlyphIsProvided(glyph rune) int {
	return int(C.TTF_GlyphIsProvided((*C.struct__TTF_Font)(f), C.Uint16(glyph)))
}

func (f *Font) GlyphMetrics(glyph rune) (minX, maxX, minY, maxY, advance int, err error) {
	var nativeMinX, nativeMaxX, nativeMinY, nativeMaxY, nativeAdvance C.int

	if C.TTF_GlyphMetrics(
		(*C.struct__TTF_Font)(f),
		C.Uint16(glyph),
		&nativeMinX,
		&nativeMaxX,
		&nativeMinY,
		&nativeMaxY,
		&nativeAdvance,
	) != 0 {
		return 0, 0, 0, 0, 0, GetError()
	}

	return int(nativeMinX), int(nativeMaxX), int(nativeMinY), int(nativeMaxY), int(nativeAdvance), nil
}

func (f *Font) SizeText(text string) (w, h int, err error) {
	var nativeW, nativeH C.int
	nativeText := C.CString(text)
	defer C.free(unsafe.Pointer(nativeText))

	if C.TTF_SizeText((*C.struct__TTF_Font)(f), nativeText, &nativeW, &nativeH) != 0 {
		return 0, 0, GetError()
	}

	return int(nativeW), int(nativeH), nil
}

func (f *Font) SizeUTF8(text string) (w, h int, err error) {
	var nativeW, nativeH C.int
	nativeText := C.CString(text)
	defer C.free(unsafe.Pointer(nativeText))

	if C.TTF_SizeUTF8((*C.struct__TTF_Font)(f), nativeText, &nativeW, &nativeH) != 0 {
		return 0, 0, GetError()
	}

	return int(nativeW), int(nativeH), nil
}

func (f *Font) SizeUnicode(text string) (w, h int, err error) {
	var nativeW, nativeH C.int
	var nativeText []C.Uint16

	for _, ch := range text {
		nativeText = append(nativeText, C.Uint16(ch))
	}

	nativeText = append(nativeText, 0)

	if C.TTF_SizeUNICODE((*C.struct__TTF_Font)(f), (*C.Uint16)(unsafe.Pointer(&nativeText[0])), &nativeW, &nativeH) != 0 {
		return 0, 0, GetError()
	}

	return int(nativeW), int(nativeH), nil
}
