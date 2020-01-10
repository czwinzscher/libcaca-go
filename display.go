package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type Display struct {
	Dp *C.struct_caca_display
}

type Event struct {
	Ev *C.struct_caca_event
}

func CreateDisplay(cv *Canvas) (Display, error) {
	var cPtr *C.struct_caca_display
	var err error

	if cv != nil {
		cPtr, err = C.caca_create_display(cv.Cv)
	} else {
		cPtr, err = C.caca_create_display(nil)
	}

	if cPtr == nil {
		return Display{}, err
	}

	return Display{Dp: cPtr}, nil
}

func (d Display) GetCanvas() Canvas {
	cPtr := C.caca_get_canvas(d.Dp)

	return Canvas{Cv: cPtr}
}

func (d Display) SetTitle(title string) error {
	ret, err := C.caca_set_display_title(d.Dp, C.CString(title))

	if int(ret) != -1 {
		return nil
	}

	return err
}

func (d Display) Refresh() {
	C.caca_refresh_display(d.Dp)
}

func (d Display) GetEvent(eventMask int, ev *Event, timeout int) {
	if ev == nil {
		C.caca_get_event(d.Dp, C.int(eventMask), nil, C.int(timeout))
	} else {
		C.caca_get_event(d.Dp, C.int(eventMask), ev.Ev, C.int(timeout))
	}
}

func (d Display) Free() {
	C.free(unsafe.Pointer(d.Dp))
}
