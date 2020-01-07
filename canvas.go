package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"unsafe"
)

type Canvas struct {
	Cv *C.struct_caca_canvas
}

func CreateCanvas(width int, height int) (Canvas, error) {
	cPtr := C.caca_create_canvas(C.int(width), C.int(height))

	if cPtr == nil {
		// TODO check errno
		return Canvas{}, errors.New("could not create canvas")
	}

	return Canvas{Cv: cPtr}, nil
}

func (c Canvas) PutStr(x int, y int, str string) {
	C.caca_put_str(c.Cv, C.int(x), C.int(y), C.CString(str))
}

func (c Canvas) SetColorAnsi(fg byte, bg byte) error {
	ret := C.caca_set_color_ansi(c.Cv, C.uint8_t(fg), C.uint8_t(bg))

	if int(ret) == 0 {
		return nil
	}

	// TODO
	return errors.New("fail")
}

func (c Canvas) Free() {
	C.free(unsafe.Pointer(c.Cv))
}
