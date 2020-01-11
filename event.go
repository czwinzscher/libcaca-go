package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

type Event struct {
	Ev *C.struct_caca_event
}

func (e Event) GetType() int {
	return int(C.caca_get_event_type(e.Ev))
}

func (e Event) GetKeyCh() int {
	return int(C.caca_get_event_key_ch(e.Ev))
}

func (e Event) GetKeyUTF32() uint32 {
	return uint32(C.caca_get_event_key_utf32(e.Ev))
}

func (e Event) GetKeyUTF8() string {
	cBuf := [7]C.char{}
	C.caca_get_event_key_utf8(e.Ev, &cBuf[0])

	goBuf := (*[7]byte)(unsafe.Pointer(&cBuf[0]))[:7:7]

	return string(goBuf)
}

func (e Event) GetMouseButton() int {
	return int(C.caca_get_event_mouse_button(e.Ev))
}

func (e Event) GetMouseButtonX() int {
	return int(C.caca_get_event_mouse_x(e.Ev))
}

func (e Event) GetMouseButtonY() int {
	return int(C.caca_get_event_mouse_y(e.Ev))
}

func (e Event) GetResizeWidth() int {
	return int(C.caca_get_event_resize_width(e.Ev))
}

func (e Event) GetResizeHeight() int {
	return int(C.caca_get_event_resize_height(e.Ev))
}
