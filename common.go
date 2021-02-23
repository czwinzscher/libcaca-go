package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

const (
	ColorBlack     = 0x00
	ColorBlue      = 0x01
	ColorGreen     = 0x02
	ColorCyan      = 0x03
	ColorRed       = 0x04
	ColorMagenta   = 0x05
	ColorBrown     = 0x06
	ColorLightgray = 0x07
	ColorDarkgray  = 0x08

	ColorLightblue    = 0x09
	ColorLightgreen   = 0x0a
	ColorLightcyan    = 0x0b
	ColorLightred     = 0x0c
	ColorLightmagenta = 0x0d
	ColorYellow       = 0x0e
	ColorWhite        = 0x0f
	ColorDefault      = 0x10
	ColorTransparent  = 0x20

	StyleBold      = 0x01
	StyleItalics   = 0x02
	StyleUnderline = 0x04
	StyleBlink     = 0x08

	EventNone         = 0x0000
	EventKeyPress     = 0x0001
	EventKeyRelease   = 0x0002
	EventMousePress   = 0x0004
	EventMouseRelease = 0x0008
	EventMouseMotion  = 0x0010
	EventResize       = 0x0020
	EventQuit         = 0x0040
	EventAny          = 0xffff

	KeyUnknown   = 0x00
	KeyCtrlA     = 0x01
	KeyCtrlB     = 0x02
	KeyCtrlC     = 0x03
	KeyCtrlD     = 0x04
	KeyCtrlE     = 0x05
	KeyCtrlF     = 0x06
	KeyCtrlG     = 0x07
	KeyBackspace = 0x08
	KeyTab       = 0x09
	KeyCtrlJ     = 0x0a
	KeyCtrlK     = 0x0b
	KeyCtrlL     = 0x0c
	KeyReturn    = 0x0d
	KeyCtrlN     = 0x0e
	KeyCtrlO     = 0x0f
	KeyCtrlP     = 0x10
	KeyCtrlQ     = 0x11
	KeyCtrlR     = 0x12
	KeyPause     = 0x13
	KeyCtrlT     = 0x14
	KeyCtrlU     = 0x15
	KeyCtrlV     = 0x16
	KeyCtrlW     = 0x17
	KeyCtrlX     = 0x18
	KeyCtrlY     = 0x19
	KeyCtrlZ     = 0x1a
	KeyEscape    = 0x1b
	KeyDelete    = 0x7f
	KeyUp        = 0x111
	KeyDown      = 0x112
	KeyLeft      = 0x113
	KeyRight     = 0x114
	KeyInsert    = 0x115
	KeyHome      = 0x116
	KeyEnd       = 0x117
	KeyPageup    = 0x118
	KeyPagedown  = 0x119
	KeyF1        = 0x11a
	KeyF2        = 0x11b
	KeyF3        = 0x11c
	KeyF4        = 0x11d
	KeyF5        = 0x11e
	KeyF6        = 0x11f
	KeyF7        = 0x120
	KeyF8        = 0x121
	KeyF9        = 0x122
	KeyF10       = 0x123
	KeyF11       = 0x124
	KeyF12       = 0x125
	KeyF13       = 0x126
	KeyF14       = 0x127
	KeyF15       = 0x128

	MagicFullwidth = 0x000ffffe
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

// UTF32ToASCII converts a UTF-32 character into an ASCII character. When no
// equivalent exists, a graphically close equivalent is sought.
//
// This function never fails, but its behaviour with illegal UTF-32 characters
// is undefined.
func UTF32ToASCII(ch uint32) rune {
	return rune(C.caca_utf32_to_ascii(C.uint32_t(ch)))
}

// UTF32IsFullwidth checks whether the given UTF-32 character should be printed
// at twice the normal width (fullwidth characters). If the character is unknown
// or if its status cannot be decided, it is treated as a standard-width
// character.
func UTF32IsFullwidth(ch uint32) bool {
	ret := int(C.caca_utf32_is_fullwidth(C.uint32_t(ch)))

	return ret == 1
}
