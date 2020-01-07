package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

const (
	COLOR_BLACK        = 0x00
	COLOR_BLUE         = 0x01
	COLOR_GREEN        = 0x02
	COLOR_CYAN         = 0x03
	COLOR_RED          = 0x04
	COLOR_MAGENTA      = 0x05
	COLOR_BROWN        = 0x06
	COLOR_LIGHTGRAY    = 0x07
	COLOR_DARKGRAY     = 0x08
	COLOR_LIGHTBLUE    = 0x09
	COLOR_LIGHTGREEN   = 0x0a
	COLOR_LIGHTCYAN    = 0x0b
	COLOR_LIGHTRED     = 0x0c
	COLOR_LIGHTMAGENTA = 0x0d
	COLOR_YELLOW       = 0x0e
	COLOR_WHITE        = 0x0f
	COLOR_DEFAULT      = 0x10
	COLOR_TRANSPARENT  = 0x20

	STYLE_BOLD      = 0x01
	STYLE_ITALICS   = 0x02
	STYLE_UNDERLINE = 0x04
	STYLE_BLINK     = 0x08

	EVENT_NONE          = 0x0000
	EVENT_KEY_PRESS     = 0x0001
	EVENT_KEY_RELEASE   = 0x0002
	EVENT_MOUSE_PRESS   = 0x0004
	EVENT_MOUSE_RELEASE = 0x0008
	EVENT_MOUSE_MOTION  = 0x0010
	EVENT_RESIZE        = 0x0020
	EVENT_QUIT          = 0x0040
	EVENT_ANY           = 0xffff

	KEY_UNKNOWN   = 0x00
	KEY_CTRL_A    = 0x01
	KEY_CTRL_B    = 0x02
	KEY_CTRL_C    = 0x03
	KEY_CTRL_D    = 0x04
	KEY_CTRL_E    = 0x05
	KEY_CTRL_F    = 0x06
	KEY_CTRL_G    = 0x07
	KEY_BACKSPACE = 0x08
	KEY_TAB       = 0x09
	KEY_CTRL_J    = 0x0a
	KEY_CTRL_K    = 0x0b
	KEY_CTRL_L    = 0x0c
	KEY_RETURN    = 0x0d
	KEY_CTRL_N    = 0x0e
	KEY_CTRL_O    = 0x0f
	KEY_CTRL_P    = 0x10
	KEY_CTRL_Q    = 0x11
	KEY_CTRL_R    = 0x12
	KEY_PAUSE     = 0x13
	KEY_CTRL_T    = 0x14
	KEY_CTRL_U    = 0x15
	KEY_CTRL_V    = 0x16
	KEY_CTRL_W    = 0x17
	KEY_CTRL_X    = 0x18
	KEY_CTRL_Y    = 0x19
	KEY_CTRL_Z    = 0x1a
	KEY_ESCAPE    = 0x1b
	KEY_DELETE    = 0x7f
	KEY_UP        = 0x111
	KEY_DOWN      = 0x112
	KEY_LEFT      = 0x113
	KEY_RIGHT     = 0x114
	KEY_INSERT    = 0x115
	KEY_HOME      = 0x116
	KEY_END       = 0x117
	KEY_PAGEUP    = 0x118
	KEY_PAGEDOWN  = 0x119
	KEY_F1        = 0x11a
	KEY_F2        = 0x11b
	KEY_F3        = 0x11c
	KEY_F4        = 0x11d
	KEY_F5        = 0x11e
	KEY_F6        = 0x11f
	KEY_F7        = 0x120
	KEY_F8        = 0x121
	KEY_F9        = 0x122
	KEY_F10       = 0x123
	KEY_F11       = 0x124
	KEY_F12       = 0x125
	KEY_F13       = 0x126
	KEY_F14       = 0x127
	KEY_F15       = 0x128
)

func GetVersion() string {
	cStr := C.caca_get_version()
	b := C.GoBytes(unsafe.Pointer(cStr), C.int(C.strlen(cStr)))

	return string(b)
}
