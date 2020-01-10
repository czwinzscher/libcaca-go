package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type Canvas struct {
	Cv *C.struct_caca_canvas
}

func CreateCanvas(width int, height int) (Canvas, error) {
	cPtr, err := C.caca_create_canvas(C.int(width), C.int(height))

	if cPtr == nil {
		return Canvas{}, err
	}

	return Canvas{Cv: cPtr}, nil
}

func (c Canvas) PutStr(x int, y int, str string) {
	C.caca_put_str(c.Cv, C.int(x), C.int(y), C.CString(str))
}

func (c Canvas) SetColorAnsi(fg byte, bg byte) error {
	ret, err := C.caca_set_color_ansi(c.Cv, C.uint8_t(fg), C.uint8_t(bg))

	if int(ret) != -1 {
		return nil
	}

	return err
}

func (c Canvas) ImportFromFile(filename string, format string) error {
	ret, err := C.caca_import_canvas_from_file(c.Cv, C.CString(filename), C.CString(format))

	if int(ret) != -1 {
		return nil
	}

	return err
}

func (c Canvas) Free() {
	C.free(unsafe.Pointer(c.Cv))
}
