package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// Event is a libcaca event structure.
type Event struct {
	Ev *C.struct_caca_event
}

// GetType returns the type of the event. This function may always be called on
// an event after display.GetEvent() was called, and its return value indicates
// which other functions may be called:
//
//     CACA_EVENT_NONE: no other function may be called.
//     CACA_EVENT_KEY_PRESS, CACA_EVENT_KEY_RELEASE: GetKeyCh(), GetKeyUTF32()
//                                                   and GetKeyUTF8() may be
//                                                   called.
//     CACA_EVENT_MOUSE_PRESS, CACA_EVENT_MOUSE_RELEASE: GetMouseButton() may be
//                                                       called.
//     CACA_EVENT_MOUSE_MOTION: GetMouseX() and GetMouseY() may be called.
//     CACA_EVENT_RESIZE: GetResizeHeight() and GetResizeWidth() may be called.
//     CACA_EVENT_QUIT: no other function may be called.
func (e Event) GetType() int {
	return int(C.caca_get_event_type(e.Ev))
}

// GetKeyCh returns either the ASCII value for the event's key, or if the key is
// not an ASCII character, an appropriate caca_key value.
//
// This function never fails, but must only be called with a valid event of type
// CACA_EVENT_KEY_PRESS or CACA_EVENT_KEY_RELEASE, or the results will be
// undefined. See GetType() for more information.
func (e Event) GetKeyCh() int {
	return int(C.caca_get_event_key_ch(e.Ev))
}

// GetKeyUTF32 returns the UTF-32/UCS-4 value for the event's key if it resolves
// to a printable character.
//
// This function never fails, but must only be called with a valid event of type
// CACA_EVENT_KEY_PRESS or CACA_EVENT_KEY_RELEASE, or the results will be
// undefined. See GetType() for more information.
func (e Event) GetKeyUTF32() uint32 {
	return uint32(C.caca_get_event_key_utf32(e.Ev))
}

// GetKeyUTF8 returns the UTF-8 value for an event's key if it resolves to a
// printable character. Up to 6 UTF-8 bytes and a null termination are returned.
//
// This function never fails, but must only be called with a valid event of type
// CACA_EVENT_KEY_PRESS or CACA_EVENT_KEY_RELEASE, or the results will be
// undefined. See GetType() for more information.
func (e Event) GetKeyUTF8() string {
	cBuf := [7]C.char{}
	C.caca_get_event_key_utf8(e.Ev, &cBuf[0])

	goBuf := (*[7]byte)(unsafe.Pointer(&cBuf[0]))[:7:7]

	return string(goBuf)
}

// GetMouseButton returns the mouse button index for the event.
//
// This function never fails, but must only be called with a valid event of type
// CACA_EVENT_MOUSE_PRESS or CACA_EVENT_MOUSE_RELEASE, or the results will be
// undefined. See GetType() for more information.
//
// This function returns 1 for the left mouse button, 2 for the right mouse
// button, and 3 for the middle mouse button.
func (e Event) GetMouseButton() int {
	return int(C.caca_get_event_mouse_button(e.Ev))
}

// GetMouseButtonX returns the X coordinate for a mouse motion event.
//
// This function never fails, but must only be called with a valid event of type
// CACA_EVENT_MOUSE_MOTION, or the results will be undefined. See GetType() for
// more information.
func (e Event) GetMouseButtonX() int {
	return int(C.caca_get_event_mouse_x(e.Ev))
}

// GetMouseButtonY returns the Y coordinate for a mouse motion event.
//
// This function never fails, but must only be called with a valid event of type
// CACA_EVENT_MOUSE_MOTION, or the results will be undefined. See GetType() for
// more information.
func (e Event) GetMouseButtonY() int {
	return int(C.caca_get_event_mouse_y(e.Ev))
}

// GetResizeWidth returns the width value for a display resize event.
//
// This function never fails, but must only be called with a valid event of type
// CACA_EVENT_RESIZE, or the results will be undefined. See GetType() for more
// information.
func (e Event) GetResizeWidth() int {
	return int(C.caca_get_event_resize_width(e.Ev))
}

// GetResizeHeight returns the height value for a display resize event.
//
// This function never fails, but must only be called with a valid event of type
// CACA_EVENT_RESIZE, or the results will be undefined. See GetType() for more
// information.
func (e Event) GetResizeHeight() int {
	return int(C.caca_get_event_resize_height(e.Ev))
}
