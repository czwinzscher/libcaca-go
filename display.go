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

func CreateDisplayWithDriver(cv *Canvas, driver string) (Display, error) {
	var cPtr *C.struct_caca_display
	var err error

	if cv != nil {
		cPtr, err = C.caca_create_display_with_driver(cv.Cv, C.CString(driver))
	} else {
		cPtr, err = C.caca_create_display_with_driver(nil, C.CString(driver))
	}

	if cPtr == nil {
		return Display{}, err
	}

	return Display{Dp: cPtr}, nil
}

func (d Display) GetDriver() string {
	return C.GoString(C.caca_get_display_driver(d.Dp))
}

func (d Display) SetDriver(driver string) int {
	return int(C.caca_set_display_driver(d.Dp, C.CString(driver)))
}

func (d Display) GetCanvas() Canvas {
	cPtr := C.caca_get_canvas(d.Dp)

	return Canvas{Cv: cPtr}
}

func (d Display) Refresh() {
	C.caca_refresh_display(d.Dp)
}

func (d Display) SetTime(usec int) error {
	ret, err := C.caca_set_display_time(d.Dp, C.int(usec))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (d Display) GetTime() int {
	return int(C.caca_get_display_time(d.Dp))
}

func (d Display) SetTitle(title string) error {
	ret, err := C.caca_set_display_title(d.Dp, C.CString(title))

	if int(ret) != -1 {
		return nil
	}

	return err
}

func (d Display) SetMouse(flag int) error {
	ret, err := C.caca_set_mouse(d.Dp, C.int(flag))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (d Display) SetCursor(flag int) error {
	ret, err := C.caca_set_cursor(d.Dp, C.int(flag))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (d Display) GetEvent(eventMask int, ev *Event, timeout int) {
	if ev == nil {
		C.caca_get_event(d.Dp, C.int(eventMask), nil, C.int(timeout))
	} else {
		C.caca_get_event(d.Dp, C.int(eventMask), ev.Ev, C.int(timeout))
	}
}

func (d Display) GetMouseX() int {
	return int(C.caca_get_mouse_x(d.Dp))
}

func (d Display) GetMouseY() int {
	return int(C.caca_get_mouse_y(d.Dp))
}

func (d Display) Free() {
	C.free(unsafe.Pointer(d.Dp))
}
