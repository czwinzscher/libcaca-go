package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

const (
	COLOR_BLACK     = 0x00
	COLOR_BLUE      = 0x01
	COLOR_GREEN     = 0x02
	COLOR_CYAN      = 0x03
	COLOR_RED       = 0x04
	COLOR_MAGENTA   = 0x05
	COLOR_BROWN     = 0x06
	COLOR_LIGHTGRAY = 0x07
	COLOR_DARKGRAY  = 0x08

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

	MAGIC_FULLWIDTH = 0x000ffffe
)

// GetVersion returns a string with the libcaca version information.
func GetVersion() string {
	cStr := C.caca_get_version()
	b := C.GoBytes(unsafe.Pointer(cStr), C.int(C.strlen(cStr)))

	return string(b)
}

// Rand returns a random number between min and max.
func Rand(min int, max int) int {
	return int(C.caca_rand(C.int(min), C.int(max)))
}

// AttrToAnsi gets the ANSI colour pair for a given attribute. The returned value
// is an 8-bit value whose higher 4 bits are the background colour and lower 4
// bits are the foreground colour.
//
// If the attribute has ARGB colours, the nearest colour is used. Special
// attributes such as CACA_DEFAULT and CACA_TRANSPARENT are not handled and are
// both replaced with CACA_LIGHTGRAY for the foreground colour and CACA_BLACK
// for the background colour.
//
// This function never fails. If the attribute value is outside the expected
// 32-bit range, higher order bits are simply ignored.
func AttrToAnsi(attr uint32) uint8 {
	return uint8(C.caca_attr_to_ansi(C.uint32_t(attr)))
}

// AttrToAnsiFg gets the ANSI foreground colour value for a given attribute. The
// returned value is either one of the CACA_RED, CACA_BLACK etc. predefined
// colours, or the special value CACA_DEFAULT meaning the media's default
// foreground value, or the special value CACA_TRANSPARENT.
//
// If the attribute has ARGB colours, the nearest colour is returned.
//
// This function never fails. If the attribute value is outside the expected
// 32-bit range, higher order bits are simply ignored.
func AttrToAnsiFg(attr uint32) uint8 {
	return uint8(C.caca_attr_to_ansi_fg(C.uint32_t(attr)))
}

// AttrToAnsiBg gets ANSI background information from attribute.
//
// Get the ANSI background colour value for a given attribute. The returned value
// is either one of the CACA_RED, CACA_BLACK etc. predefined colours, or the
// special value CACA_DEFAULT meaning the media's default background value, or
// the special value CACA_TRANSPARENT.
//
// If the attribute has ARGB colours, the nearest colour is returned.
//
// This function never fails. If the attribute value is outside the expected
// 32-bit range, higher order bits are simply ignored.
func AttrToAnsiBg(attr uint32) uint8 {
	return uint8(C.caca_attr_to_ansi_bg(C.uint32_t(attr)))
}

// AttrToRGB12Fg gets the 12-bit foreground colour value for a given attribute.
// The returned value is a native-endian encoded integer with each red, green
// and blue values encoded on 8 bits in the following order:
//
//     8-11 most significant bits: red
//     4-7 most significant bits: green
//     least significant bits: blue
//
// This function never fails. If the attribute value is outside the expected
// 32-bit range, higher order bits are simply ignored.
func AttrToRGB12Fg(attr uint32) uint16 {
	return uint16(C.caca_attr_to_rgb12_fg(C.uint32_t(attr)))
}

// AttrToRGB12Bg gets the 12-bit background colour value for a given attribute.
// The returned value is a native-endian encoded integer with each red, green
// and blue values encoded on 8 bits in the following order:
//
//     8-11 most significant bits: red
//     4-7 most significant bits: green
//     least significant bits: blue
//
// This function never fails. If the attribute value is outside the expected
// 32-bit range, higher order bits are simply ignored.
func AttrToRGB12Bg(attr uint32) uint16 {
	return uint16(C.caca_attr_to_rgb12_bg(C.uint32_t(attr)))
}

// UTF8ToUTF32 converts a UTF-8 character read from a string and returns its
// value in the UTF-32 character set.
//
// If a null byte was reached before the expected end of the UTF-8 sequence,
// this function returns zero.
//
// This function never fails, but its behaviour with illegal UTF-8 sequences is
// undefined.
func UTF8ToUTF32(s string) uint32 {
	return uint32(C.caca_utf8_to_utf32(C.CString(s), nil))
}

// UTF32ToUTF8 Convert a UTF-32 character read from a string and returns its value
// in the UTF-8 character set.
//
// This function never fails, but its behaviour with illegal UTF-32 characters
// is undefined.
func UTF32ToUTF8(ch uint32) string {
	cBuf := [7]C.char{}
	C.caca_utf32_to_utf8(&cBuf[0], C.uint32_t(ch))

	goBuf := (*[7]byte)(unsafe.Pointer(&cBuf[0]))[:7:7]

	return string(goBuf)
}

// UTF32ToCP437 converts a UTF-32 character and returns its value in the CP437
// character set, or "?" if the character has no equivalent.
func UTF32ToCP437(ch uint32) uint8 {
	return uint8(C.caca_utf32_to_cp437(C.uint32_t(ch)))
}

// CP437ToUTF32 converts a CP437 character and returns its value in the UTF-32
// character set, or zero if the character is a CP437 control character.
func CP437ToUTF32(ch uint8) uint32 {
	return uint32(C.caca_cp437_to_utf32(C.uint8_t(ch)))
}

// UTF32ToAscii converts a UTF-32 character into an ASCII character. When no
// equivalent exists, a graphically close equivalent is sought.
//
// This function never fails, but its behaviour with illegal UTF-32 characters
// is undefined.
func UTF32ToAscii(ch uint32) rune {
	return rune(C.caca_utf32_to_ascii(C.uint32_t(ch)))
}

// UTF32IsFullwidth checks whether the given UTF-32 character should be printed
// at twice the normal width (fullwidth characters). If the character is unknown
// or if its status cannot be decided, it is treated as a standard-width
// character.
func UTF32IsFullwidth(ch uint32) bool {
	ret := int(C.caca_utf32_is_fullwidth(C.uint32_t(ch)))

	return int(ret) == 1
}
