package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
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
	if cv != nil {
		cPtr = C.caca_create_display(cv.Cv)
	} else {
		cPtr = C.caca_create_display(nil)
	}

	if cPtr == nil {
		// TODO check errno
		return Display{}, errors.New("could not create display")
	}

	return Display{Dp: cPtr}, nil
}

func (d Display) GetCanvas() Canvas {
	cPtr := C.caca_get_canvas(d.Dp)

	return Canvas{Cv: cPtr}
}

func (d Display) SetTitle(title string) {
	C.caca_set_display_title(d.Dp, C.CString(title))
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
