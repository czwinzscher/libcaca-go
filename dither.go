package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// Dither is a dither structure.
type Dither struct {
	Di *C.struct_caca_dither
}

// CreateDither creates a dither structure from its coordinates
// (depth, width, height and pitch) and pixel mask values. If the depth is 8
// bits per pixel, the mask values are ignored and the colour palette should be
// set using the SetPalette() function. For depths greater than 8 bits per
// pixel, a zero alpha mask causes the alpha values to be ignored.
//
// If an error occurs the according errno is returned.
func CreateDither(bpp int, w int, h int, pitch int, rmask uint32, gmask uint32, bmask uint32, amask uint32) (Dither, error) {
	cPtr, err := C.caca_create_dither(C.int(bpp), C.int(w), C.int(h), C.int(pitch), C.uint32_t(rmask), C.uint32_t(gmask), C.uint32_t(bmask), C.uint32_t(amask))

	if cPtr == nil {
		return Dither{}, err
	}

	return Dither{Di: cPtr}, nil
}

// SetPalette sets the palette of an 8 bits per pixel bitmap. Values should be
// between 0 and 4095 (0xfff).
//
// If an error occurs the according errno is returned.
func (di Dither) SetPalette(red [256]uint32, green [256]uint32, blue [256]uint32, alpha [256]uint32) error {
	ret, err := C.caca_set_dither_palette(di.Di, (*C.uint32_t)(&red[0]), (*C.uint32_t)(&green[0]), (*C.uint32_t)(&blue[0]), (*C.uint32_t)(&alpha[0]))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// SetBrightness sets the brightness of the dither.
//
// If an error occurs the according errno is returned.
func (di Dither) SetBrightness(brightness float64) error {
	ret, err := C.caca_set_dither_brightness(di.Di, C.float(brightness))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetBrightness returns the brightness of the dither.
func (di Dither) GetBrightness() float64 {
	return float64(C.caca_get_dither_brightness(di.Di))
}

// SetGamma sets the gamma of the dither object. A negative value causes colour
// inversion.
//
// If an error occurs the according errno is returned.
func (di Dither) SetGamma(gamma float64) error {
	ret, err := C.caca_set_dither_gamma(di.Di, C.float(gamma))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetGamma returns the gamma of the dither object.
func (di Dither) GetGamma() float64 {
	return float64(C.caca_get_dither_gamma(di.Di))
}

// SetContrast sets the contrast of the dither object.
//
// If an error occurs the according errno is returned.
func (di Dither) SetContrast(gamma float64) error {
	ret, err := C.caca_set_dither_contrast(di.Di, C.float(gamma))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetContrast returns the contrast of the dither object.
func (di Dither) GetContrast() float64 {
	return float64(C.caca_get_dither_contrast(di.Di))
}

// SetAntialias tells the renderer whether to antialias the dither.
// Antialiasing smoothens the rendered image and avoids the commonly seen
// staircase effect.
//
//     "none": no antialiasing.
//     "prefilter" or "default": simple prefilter antialiasing. This is the
//                               default value.
//
// If an error occurs the according errno is returned.
func (di Dither) SetAntialias(str string) error {
	ret, err := C.caca_set_dither_antialias(di.Di, C.CString(str))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetAntialias returns the antialiasing method of the dither object.
func (di Dither) GetAntialias() string {
	return C.GoString(C.caca_get_dither_antialias(di.Di))
}

// SetColor tells the renderer which colours should be used to render the
// bitmap. Valid values for str are:
//
//     "mono": use light gray on a black background.
//     "gray": use white and two shades of gray on a black background.
//     "8": use the 8 ANSI colours on a black background.
//     "16": use the 16 ANSI colours on a black background.
//     "fullgray": use black, white and two shades of gray for both the
//                 characters and the background.
//     "full8": use the 8 ANSI colours for both the characters and the
//              background.
//     "full16" or "default": use the 16 ANSI colours for both the characters
//                            and the background. This is the default value.
//
// If an error occurs the according errno is returned.
func (di Dither) SetColor(str string) error {
	ret, err := C.caca_set_dither_color(di.Di, C.CString(str))

	if ret == -1 {
		return err
	}

	return nil
}

// GetColor returns the current colour mode of the dither object.
func (di Dither) GetColor() string {
	return C.GoString(C.caca_get_dither_color(di.Di))
}

// SetCharset tells the renderer which characters should be used to render the
// dither. Valid values for str are:
//
//     "ascii" or "default": use only ASCII characters. This is the default
//                           value.
//     "shades": use Unicode characters "U+2591 LIGHT SHADE",
//               "U+2592 MEDIUM SHADE" and "U+2593 DARK SHADE". These
//               characters are also present in the CP437 codepage available
//               on DOS and VGA.
//     "blocks": use Unicode quarter-cell block combinations. These characters
//               are only found in the Unicode set.
//
// If an error occurs the according errno is returned.
func (di Dither) SetCharset(str string) error {
	ret, err := C.caca_set_dither_charset(di.Di, C.CString(str))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetCharset returns the current character set of the dither object.
func (di Dither) GetCharset() string {
	return C.GoString(C.caca_get_dither_charset(di.Di))
}

// SetAlgorithm tells the renderer which dithering algorithm should be used.
// Dithering is necessary because the picture being rendered has usually far
// more colours than the available palette. Valid values for str are:
//
//     "none": no dithering is used, the nearest matching colour is used.
//     "ordered2": use a 2x2 Bayer matrix for dithering.
//     "ordered4": use a 4x4 Bayer matrix for dithering.
//     "ordered8": use a 8x8 Bayer matrix for dithering.
//     "random": use random dithering.
//     "fstein": use Floyd-Steinberg dithering. This is the default value.
//
// If an error occurs the according errno is returned.
func (di Dither) SetAlgorithm(str string) error {
	ret, err := C.caca_set_dither_algorithm(di.Di, C.CString(str))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetAlgorithm returns the current dithering algorithm of the dither object.
func (di Dither) GetAlgorithm() string {
	return C.GoString(C.caca_get_dither_algorithm(di.Di))
}

// Bitmap dithers a bitmap at the given coordinates. The dither can be of any
// size and will be stretched to the text area.
func (di Dither) Bitmap(cv Canvas, x int, y int, w int, h int, pixels []byte) {
	C.caca_dither_bitmap(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h), di.Di, unsafe.Pointer(&pixels[0]))
}

// Free frees the memory allocated by CreateDither().
func (di Dither) Free() {
	C.caca_free_dither(di.Di)
}
