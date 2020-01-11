package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type Dither struct {
	Di *C.struct_caca_dither
}

func CreateDither(bpp int, w int, h int, pitch int, rmask uint32, gmask uint32, bmask uint32, amask uint32) (Dither, error) {
	cPtr, err := C.caca_create_dither(C.int(bpp), C.int(w), C.int(h), C.int(pitch), C.uint32_t(rmask), C.uint32_t(gmask), C.uint32_t(bmask), C.uint32_t(amask))

	if cPtr == nil {
		return Dither{}, err
	}

	return Dither{Di: cPtr}, nil
}

func (di Dither) SetPalette(red [256]uint32, green [256]uint32, blue [256]uint32, alpha [256]uint32) error {
	ret, err := C.caca_set_dither_palette(di.Di, (*C.uint32_t)(&red[0]), (*C.uint32_t)(&green[0]), (*C.uint32_t)(&blue[0]), (*C.uint32_t)(&alpha[0]))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (di Dither) SetBrightness(brightness float64) error {
	ret, err := C.caca_set_dither_brightness(di.Di, C.float(brightness))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (di Dither) GetBrightness() float64 {
	return float64(C.caca_get_dither_brightness(di.Di))
}

func (di Dither) SetGamma(gamma float64) error {
	ret, err := C.caca_set_dither_gamma(di.Di, C.float(gamma))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (di Dither) GetGamma() float64 {
	return float64(C.caca_get_dither_gamma(di.Di))
}

func (di Dither) SetContrast(gamma float64) error {
	ret, err := C.caca_set_dither_contrast(di.Di, C.float(gamma))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (di Dither) GetContrast() float64 {
	return float64(C.caca_get_dither_contrast(di.Di))
}

func (di Dither) SetAntialias(str string) error {
	ret, err := C.caca_set_dither_antialias(di.Di, C.CString(str))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (di Dither) GetAntialias() string {
	return C.GoString(C.caca_get_dither_antialias(di.Di))
}

func (di Dither) SetColor(str string) error {
	ret, err := C.caca_set_dither_color(di.Di, C.CString(str))

	if ret == -1 {
		return err
	}

	return nil
}

func (di Dither) GetColor() string {
	return C.GoString(C.caca_get_dither_color(di.Di))
}

func (di Dither) SetCharset(str string) error {
	ret, err := C.caca_set_dither_charset(di.Di, C.CString(str))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (di Dither) GetCharset() string {
	return C.GoString(C.caca_get_dither_charset(di.Di))
}

func (di Dither) SetAlgorithm(str string) error {
	ret, err := C.caca_set_dither_algorithm(di.Di, C.CString(str))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (di Dither) GetAlgorithm() string {
	return C.GoString(C.caca_get_dither_algorithm(di.Di))
}

func (di Dither) Bitmap(cv Canvas, x int, y int, w int, h int, pixels []byte) {
	C.caca_dither_bitmap(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h), di.Di, unsafe.Pointer(&pixels[0]))
}

func (di Dither) Free() {
	C.free(unsafe.Pointer(di.Di))
}
