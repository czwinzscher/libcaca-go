package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// Canvas is a libcaca canvas.
type Canvas struct {
	Cv *C.struct_caca_canvas
}

// CreateCanvas initialises internal libcaca structures and the backend that
// will be used for subsequent graphical operations. It must be the first
// libcaca function to be called in a function.
// Free() should be called at the end of the program to free all allocated
// resources.
//
// Both the cursor and the canvas handle are initialised at the top-left corner.
//
// If an error occurs, nil and the according errno is returned.
func CreateCanvas(width int, height int) (Canvas, error) {
	cPtr, err := C.caca_create_canvas(C.int(width), C.int(height))

	if cPtr == nil {
		return Canvas{}, err
	}

	return Canvas{Cv: cPtr}, nil
}

// SetSize sets the canvas' width and height, in character cells.
//
// The contents of the canvas are preserved to the extent of the new canvas
// size. Newly allocated character cells ate the right and/or at the bottom
// of the canvas are filled with spaces.
//
// If as a result of the resize the cursor coordinates fall outside the
// new canvas boundaries, they are readjusted. For instance, if the current
// X cursor coordinate is 11 and the requested width is 10, the new X cursor
// coordinate will be 10.
//
// It is an error to try to resize the canvas if an output driver has been
// attached to the canvas using CreateDisplay(). You need to remove the
// output driver using display.Free() before you can change the canvas size
// again. However, the caca output driver can cause a canvas resize through
// user interaction. See the Event documentation for more about this.
//
// If an error occurs the according errno is returned.
func (cv Canvas) SetSize(width int, height int) error {
	ret, err := C.caca_set_canvas_size(cv.Cv, C.int(width), C.int(height))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetWidth returns the canvas' width, in character cells.
func (cv Canvas) GetWidth() int {
	return int(C.caca_get_canvas_width(cv.Cv))
}

// GetHeight returns the canvas' height, in character cells.
func (cv Canvas) GetHeight() int {
	return int(C.caca_get_canvas_height(cv.Cv))
}

// GoToXY puts the cursor at the given coordinates. Functions making use of the
// cursor will use the new values. Setting the cursor position outside the
// canvas is legal but the cursor will not be shown.
func (cv Canvas) GoToXY(x int, y int) {
	C.caca_gotoxy(cv.Cv, C.int(x), C.int(y))
}

// WhereX retrieves the X coordinate of the cursor's position.
func (cv Canvas) WhereX() int {
	return int(C.caca_wherex(cv.Cv))
}

// WhereY retrieves the Y coordinate of the cursor's position.
func (cv Canvas) WhereY() int {
	return int(C.caca_wherey(cv.Cv))
}

// PutChar prints an ASCII or Unicode character at the given coordinates, using
// the default foreground and background colour values.
//
// If the coordinates are outside the canvas boundaries, nothing is printed.
// If a fullwidth Unicode character gets overwritten, its remaining visible
// parts are replaced with spaces. If the canvas' boundaries would split the
// fullwidth character in two, a space is printed instead.
//
// The behaviour when printing non-printable characters or invalid UTF-32
// characters is undefined. To print a sequence of bytes forming an UTF-8
// character instead of an UTF-32 character, use the PutStr() function.
//
// This function returns the width of the printed character. If it is a
// fullwidth character, 2 is returned. Otherwise, 1 is returned.
func (cv Canvas) PutChar(x int, y int, ch rune) int {
	return int(C.caca_put_char(cv.Cv, C.int(x), C.int(y), C.uint32_t(ch)))
}

// GetChar gets the ASCII or Unicode value of the character at the given
// coordinates. If the value is less or equal to 127 (0x7f), the character can
// be printed as ASCII. Otherise, it must be handled as a UTF-32 value.
//
// If the coordinates are outside the canvas boundaries, a space (0x20) is
// returned.
//
// A special exception is when MAGIC_FULLWIDTH is returned. This value
// is guaranteed not to be a valid Unicode character, and indicates that the
// character at the left of the requested one is a fullwidth character.
func (cv Canvas) GetChar(x int, y int) rune {
	return rune(C.caca_get_char(cv.Cv, C.int(x), C.int(y)))
}

// PutStr prints an UTF-8 string at the given coordinates, using the default
// foreground and background values. The coordinates may be outside the canvas
// boundaries (eg. a negative Y coordinate) and the string will be cropped
// accordingly if it is too long.
//
// See PutChar() for more information on how fullwidth characters are
// handled when overwriting each other or at the canvas' boundaries.
//
// This function returns the number of cells printed by the string. It is not
// the number of characters printed, because fullwidth characters account for
// two cells.
func (cv Canvas) PutStr(x int, y int, str string) int {
	return int(C.caca_put_str(cv.Cv, C.int(x), C.int(y), C.CString(str)))
}

// Clear clears the canvas using the current foreground and background colours.
func (cv Canvas) Clear() {
	C.caca_clear_canvas(cv.Cv)
}

// SetHandle sets the canvas' handle. Blitting functions will use the handle
// value to put the canvas at the proper coordinates.
func (cv Canvas) SetHandle(x int, y int) {
	C.caca_set_canvas_handle(cv.Cv, C.int(x), C.int(y))
}

// GetHandleX retrieves the X coordinate of the canvas' handle.
func (cv Canvas) GetHandleX() int {
	return int(C.caca_get_canvas_handle_x(cv.Cv))
}

// GetHandleY retrieves the Y coordinate of the canvas' handle.
func (cv Canvas) GetHandleY() int {
	return int(C.caca_get_canvas_handle_y(cv.Cv))
}

// Blit blits a canvas onto another one at the given coordinates. An optional
// mask canvas can be used.
//
// If an error occurs the according errno is returned.
func (cv Canvas) Blit(x int, y int, src Canvas, mask *Canvas) error {
	var ret C.int
	var err error

	if mask == nil {
		ret, err = C.caca_blit(cv.Cv, C.int(x), C.int(y), src.Cv, nil)
	} else {
		ret, err = C.caca_blit(cv.Cv, C.int(x), C.int(y), src.Cv, mask.Cv)
	}

	if int(ret) == -1 {
		return err
	}

	return nil
}

// SetBoundaries sets new boundaries for a canvas. This function can be used to
// crop a canvas, to expand it or for combinations of both actions. All frames
// are affected by this function.
//
// If an error occurs the according errno is returned.
func (cv Canvas) SetBoundaries(x int, y int, w int, h int) error {
	ret, err := C.caca_set_canvas_boundaries(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// DisableDirtyRect disables dirty rectangle handling for all libcaca graphic
// calls. This is handy when the calling application needs to do slow operations
// within a known area. Just call AddDirtyRect() afterwards.
//
// This function is recursive. Dirty rectangles are only reenabled when
// EnableDirtyRect() is called as many times.
func (cv Canvas) DisableDirtyRect() {
	C.caca_disable_dirty_rect(cv.Cv)
}

// EnableDirtyRect enables dirty rectangles.
// This function can only be called after DisableDirtyRect() was called.
func (cv Canvas) EnableDirtyRect() {
	C.caca_enable_dirty_rect(cv.Cv)
}

// GetDirtyRectCount gets the number of dirty rectangles in a canvas. Dirty
// rectangles are areas that contain cells that have changed since the last reset.
//
// The dirty rectangles are used internally by display drivers to optimise
// rendering by avoiding to redraw the whole screen. Once the display driver has
// rendered the canvas, it resets the dirty rectangle list.
//
// Dirty rectangles are guaranteed not to overlap.
func (cv Canvas) GetDirtyRectCount() int {
	return int(C.caca_get_dirty_rect_count(cv.Cv))
}

// GetDirtyRect gets the canvas's given dirty rectangle coordinates. The index
// must be within the dirty rectangle count. See caca_get_dirty_rect_count() for
// how to compute this count.
//
// If an error occurs the according errno is returned.
func (cv Canvas) GetDirtyRect(idx int) (map[string]int, error) {
	var x, y, width, height C.int

	ret, err := C.caca_get_dirty_rect(cv.Cv, C.int(idx), &x, &y, &width, &height)

	if int(ret) == -1 {
		return map[string]int{}, err
	}

	return map[string]int{"x": int(x), "y": int(y), "width": int(width), "height": int(height)}, nil
}

// AddDirtyRect adds an invalidating zone to the canvas's dirty rectangle list.
// For more information about the dirty rectangles, see GetDirtyRect().
//
// This function may be useful to force refresh of a given zone of the canvas
// even if the dirty rectangle tracking indicates that it is unchanged. This may
// happen if the canvas contents were somewhat directly modified.
//
// If an error occurs the according errno is returned.
func (cv Canvas) AddDirtyRect(x int, y int, width int, height int) error {
	ret, err := C.caca_add_dirty_rect(cv.Cv, C.int(x), C.int(y), C.int(width), C.int(height))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// RemovesDirtyRect removes an area from the dirty rectangle list.
//
// Mark a cell area in the canvas as not dirty. For more information about the
// dirty rectangles, see GetDirtyRect().
//
// Values such that xmin > xmax or ymin > ymax indicate that the dirty rectangle
// is empty. They will be silently ignored.
//
// If an error occurs the according errno is returned.
func (cv Canvas) RemoveDirtyRect(x int, y int, width int, height int) error {
	ret, err := C.caca_remove_dirty_rect(cv.Cv, C.int(x), C.int(y), C.int(width), C.int(height))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// ClearDirtyRectList empties the canvas' dirty rectabgle list.
func (cv Canvas) ClearDirtyRectList() {
	C.caca_clear_dirty_rect_list(cv.Cv)
}

// Invert a canvas' colours (black becomes white, red becomes cyan, etc.)
// without changing the characters in it.
func (cv Canvas) Invert() {
	C.caca_invert(cv.Cv)
}

// Flip flips a canvas horizontally, choosing characters that look like the
// mirrored version wherever possible. Some characters will stay unchanged by the
// process, but the operation is guaranteed to be involutive: performing it
// again gives back the original canvas.
func (cv Canvas) Flip() {
	C.caca_flip(cv.Cv)
}

// Flop flips a canvas vertically, choosing characters that look like the mirrored
// version wherever possible. Some characters will stay unchanged by the process,
// but the operation is guaranteed to be involutive: performing it again gives
// back the original canvas.
func (cv Canvas) Flop() {
	C.caca_flop(cv.Cv)
}

// Rotate applies a 180-degree transformation to a canvas, choosing characters
// that look like the upside-down version wherever possible. Some characters will
// stay unchanged by the process, but the operation is guaranteed to be
// involutive: performing it again gives back the original canvas.
func (cv Canvas) Rotate180() {
	C.caca_rotate_180(cv.Cv)
}

// RotateLeft rotates a canvas, 90 degrees counterclockwise.
//
// Apply a 90-degree transformation to a canvas, choosing characters
// that look like the rotated version wherever possible. Characters cells are
// rotated two-by-two. Some characters will stay unchanged by the process, some
// others will be replaced by close equivalents. Fullwidth characters at odd
// horizontal coordinates will be lost. The operation is not guaranteed to be
// reversible at all.
//
// Note that the width of the canvas is divided by two and becomes the new height.
// Height is multiplied by two and becomes the new width. If the original width is
// an odd number, the division is rounded up.
//
// If an error occurs the according errno is returned.
func (cv Canvas) RotateLeft() error {
	ret, err := C.caca_rotate_left(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}

// RotateRight rotates a canvas, 90 degrees clockwise..
//
// Apply a 90-degree transformation to a canvas, choosing characters that look
// like the rotated version wherever possible. Characters cells are rotated
// two-by-two. Some characters will stay unchanged by the process, some others
// will be replaced by close equivalents. Fullwidth characters at odd horizontal
// coordinates will be lost. The operation is not guaranteed to be reversible at
// all.
//
// Note that the width of the canvas is divided by two and becomes the new height.
// Height is multiplied by two and becomes the new width. If the original width is
// an odd number, the division is rounded up.
//
// If an error occurs the according errno is returned.
func (cv Canvas) RotateRight() error {
	ret, err := C.caca_rotate_right(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}

// StretchLeft rotates and stretches a canvas, 90 degrees counterclockwise.
//
// Apply a 90-degree transformation to a canvas, choosing characters that look
// like the rotated version wherever possible. Some characters will stay unchanged
// by the process, some others will be replaced by close equivalents. Fullwidth
// characters will be lost. The operation is not guaranteed to be reversible at
// all.
//
// Note that the width and height of the canvas are swapped, causing its aspect
// ratio to look stretched.
//
// If an error occurs the according errno is returned.
func (cv Canvas) StretchLeft() error {
	ret, err := C.caca_stretch_left(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}

// StretchRight rotates and stretches a canvas, 90 degrees clockwise.
//
// Apply a 270-degree transformation to a canvas, choosing characters that look
// like the rotated version wherever possible. Some characters will stay unchanged
// by the process, some others will be replaced by close equivalents. Fullwidth
// characters will be lost. The operation is not guaranteed to be reversible at
// all.
//
// Note that the width and height of the canvas are swapped, causing its aspect
// ratio to look stretched.
//
// If an error occurs the according errno is returned.
func (cv Canvas) StretchRight() error {
	ret, err := C.caca_stretch_right(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetAttr gets the internal libcaca attribute value of the character at the given
// coordinates. The attribute value has 32 significant bits, organised as follows
// from MSB to LSB:
//
//    3 bits for the background alpha
//    4 bits for the background red component
//    4 bits for the background green component
//    3 bits for the background blue component
//    3 bits for the foreground alpha
//    4 bits for the foreground red component
//    4 bits for the foreground green component
//    3 bits for the foreground blue component
//    4 bits for the bold, italics, underline and blink flags
//
// If the coordinates are outside the canvas boundaries, the current attribute is
// returned.
func (cv Canvas) GetAttr(x int, y int) rune {
	return rune(C.caca_get_attr(cv.Cv, C.int(x), C.int(y)))
}

// SetAttr sets the default character attribute for drawing. Attributes define
// foreground and background colour, transparency, bold, italics and underline
// styles, as well as blink. String functions such as Printf() and graphical
// primitive functions such as DrawLine() will use this attribute.
//
// The value of attr is either:
//
//    - a 32-bit integer as returned by GetAttr(), in which case it also
//      contains colour information,
//    - a combination (bitwise OR) of style values (CACA_UNDERLINE, CACA_BLINK,
//      CACA_BOLD and CACA_ITALICS), in which case setting the attribute does not
//      modify the current colour information.
//
// To retrieve the current attribute value, use GetAttr(-1,-1).
func (cv Canvas) SetAttr(attr rune) {
	C.caca_set_attr(cv.Cv, C.uint32_t(attr))
}

// UnsetAttr unsets flags in the default character attribute for drawing.
// Attributes define foreground and background colour, transparency, bold, italics
// and underline styles, as well as blink. String functions such as Printf() and
// graphical primitive functions such as DrawLine() will use this attribute.
//
// The value of attr is a combination (bitwise OR) of style values
// (CACA_UNDERLINE, CACA_BLINK, CACA_BOLD and CACA_ITALICS). Unsetting these
// attributes does not modify the current colour information.
//
// To retrieve the current attribute value, use caca_get_attr(-1,-1).
func (cv Canvas) UnsetAttr(attr rune) {
	C.caca_unset_attr(cv.Cv, C.uint32_t(attr))
}

// ToggleAttr toggles flags in the default character attribute for drawing.
// Attributes define foreground and background colour, transparency, bold, italics
// and underline styles, as well as blink. String functions such as Printf() and
// graphical primitive functions such as caca_draw_line() will use this attribute.
//
// The value of attr is a combination (bitwise OR) of style values
// (CACA_UNDERLINE, CACA_BLINK, CACA_BOLD and CACA_ITALICS). Toggling these
// attributes does not modify the current colour information.
//
// To retrieve the current attribute value, use caca_get_attr(-1,-1).
func (cv Canvas) ToggleAttr(attr rune) {
	C.caca_toggle_attr(cv.Cv, C.uint32_t(attr))
}

// PutAttr sets the character attribute, without changing the character's value.
// If the character at the given coordinates is a fullwidth character, both cells'
// attributes are replaced.
//
// The value of attr is either:
//
//     - a 32-bit integer as returned by caca_get_attr(), in which case it also
//       contains colour information,
//     - a combination (bitwise OR) of style values (CACA_UNDERLINE, CACA_BLINK,
//       CACA_BOLD and CACA_ITALICS), in which case setting the attribute does not
//       modify the current colour information.
func (cv Canvas) PutAttr(x int, y int, attr rune) {
	C.caca_put_attr(cv.Cv, C.int(x), C.int(y), C.uint32_t(attr))
}

// SetColorAnsi sets the default ANSI colour pair for text drawing. String
// functions such as Printf() and graphical primitive functions such as DrawLine()
// will use these attributes.
//
// Color values are those defined in common.go, such as CACA_RED or
// CACA_TRANSPARENT.
//
// If an error occurs the according errno is returned.
func (cv Canvas) SetColorAnsi(fg byte, bg byte) error {
	ret, err := C.caca_set_color_ansi(cv.Cv, C.uint8_t(fg), C.uint8_t(bg))

	if int(ret) != -1 {
		return nil
	}

	return err
}

// SetColorARGB sets the default ARGB colour pair for text drawing. String
// functions such as Printf() and graphical primitive functions such as
// DrawLine() will use these attributes.
//
// Colors are 16-bit ARGB values, each component being coded on 4 bits. For
// instance, 0xf088 is solid dark cyan (A=15 R=0 G=8 B=8), and 0x8fff is white
// with 50% alpha (A=8 R=15 G=15 B=15).
func (cv Canvas) SetColorARGB(fg int16, bg int16) {
	C.caca_set_color_argb(cv.Cv, C.uint16_t(fg), C.uint16_t(bg))
}

// DrawLine draws a line on the canvas using the given character.
func (cv Canvas) DrawLine(x1 int, y1 int, x2 int, y2 int, ch rune) {
	C.caca_draw_line(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.uint32_t(ch))
}

// DrawPolyline draws a polyline on the canvas using the given character and
// coordinate slice. The first and last points are not connected, hence in order
// to draw a polygon you need to specify the starting point at the end of the list
// as well.
func (cv Canvas) DrawPolyline(xy []IntPair, ch rune) {
	var cx, cy []C.int

	for _, e := range xy {
		cx = append(cx, C.int(e.First))
		cy = append(cy, C.int(e.Second))
	}

	C.caca_draw_polyline(cv.Cv, &cx[0], &cy[0], C.int(len(xy)-1), C.uint32_t(ch))
}

// DrawThinLine draws a thin line on the canvas, using ASCII art.
func (cv Canvas) DrawThinLine(x1 int, y1 int, x2 int, y2 int) {
	C.caca_draw_thin_line(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// DrawThinPolyline draws a thin polyline on the canvas using the given coordinate
// slice and with ASCII art. The first and last points are not connected, so in
// order to draw a polygon you need to specify the starting point at the end of
// the list as well.
func (cv Canvas) DrawThinPolyline(xy []IntPair) {
	var cx, cy []C.int

	for _, e := range xy {
		cx = append(cx, C.int(e.First))
		cy = append(cy, C.int(e.Second))
	}

	C.caca_draw_thin_polyline(cv.Cv, &cx[0], &cy[0], C.int(len(xy)-1))
}

// DrawCircle draws a circle on the canvas using the given character.
func (cv Canvas) DrawCircle(x int, y int, r int, ch rune) {
	C.caca_draw_circle(cv.Cv, C.int(x), C.int(y), C.int(r), C.uint32_t(ch))
}

// DrawEllipse draws an ellipse on the canvas using the given character.
func (cv Canvas) DrawEllipse(xo int, yo int, a int, b int, ch rune) {
	C.caca_draw_ellipse(cv.Cv, C.int(xo), C.int(yo), C.int(a), C.int(b), C.uint32_t(ch))
}

// DrawThinEllipse draws a thin ellipse on the canvas.
func (cv Canvas) DrawThinEllipse(xo int, yo int, a int, b int) {
	C.caca_draw_thin_ellipse(cv.Cv, C.int(xo), C.int(yo), C.int(a), C.int(b))
}

// FillEllipse fills an ellipse on the canvas using the given character.
func (cv Canvas) FillEllipse(xo int, yo int, a int, b int, ch rune) {
	C.caca_fill_ellipse(cv.Cv, C.int(xo), C.int(yo), C.int(a), C.int(b), C.uint32_t(ch))
}

// DrawBox draws a box on the canvas using the given character.
func (cv Canvas) DrawBox(x int, y int, w int, h int, ch rune) {
	C.caca_draw_box(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h), C.uint32_t(ch))
}

// DrawThinBox draws a thin box on the canvas.
func (cv Canvas) DrawThinBox(x int, y int, w int, h int) {
	C.caca_draw_thin_box(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h))
}

// DrawCP437Box draws a box on the canvas using CP437 characters.
func (cv Canvas) DrawCP437Box(x int, y int, w int, h int) {
	C.caca_draw_cp437_box(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h))
}

// FillBox fills a box on the canvas using the given character.
func (cv Canvas) FillBox(x int, y int, w int, h int, ch rune) {
	C.caca_fill_box(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h), C.uint32_t(ch))
}

// DrawTriangle draws a triangle on the canvas using the given character.
func (cv Canvas) DrawTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int, ch rune) {
	C.caca_draw_triangle(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3), C.uint32_t(ch))
}

// DrawThinTriangle draws a thin triangle on the canvas.
func (cv Canvas) DrawThinTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) {
	C.caca_draw_thin_triangle(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3))
}

// FillTriangle fills a triangle on the canvas using the given character.
func (cv Canvas) FillTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int, ch rune) {
	C.caca_fill_triangle(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3), C.uint32_t(ch))
}

// FillTriangleTextured fills a triangle on the canvas using an arbitrary-sized
// texture.
//
// This function fails if one or both the canvas are missing.
func (cv Canvas) FillTriangleTextured(coords [6]int, tex Canvas, uv [6]float64) int {
	var cCoords [6]C.int
	for i, c := range coords {
		cCoords[i] = C.int(c)
	}

	var cUv [6]C.float
	for i, u := range uv {
		cUv[i] = C.float(u)
	}

	return int(C.caca_fill_triangle_textured(cv.Cv, &cCoords[0], tex.Cv, &cUv[0]))
}

// GetFrameCount returns the current canvas' frame count.
func (cv Canvas) GetFrameCount() int {
	return int(C.caca_get_frame_count(cv.Cv))
}

// SetFrame sets the active canvas frame. All subsequent drawing operations will
// be performed on that frame. The current painting context set by SetAttr() is
// inherited.
//
// If the frame index is outside the canvas' frame range, nothing happens.
//
// If an error occurs the according errno is returned.
func (cv Canvas) SetFrame(id int) error {
	ret, err := C.caca_set_frame(cv.Cv, C.int(id))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetFrameName returns the current frame's name. The returned string is valid
// until the frame is deleted or SetFrameName() is called to change the frame
// name again.
func (cv Canvas) GetFrameName() string {
	return C.GoString(C.caca_get_frame_name(cv.Cv))
}

// SetFrameName sets the current frame's name. Upon creation, a frame has a
// default name of "frame#xxxxxxxx" where xxxxxxxx is a self-incrementing
// hexadecimal number.
func (cv Canvas) SetFrameName(name string) error {
	ret, err := C.caca_set_frame_name(cv.Cv, C.CString(name))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// CreateFrame creates a new frame within the canvas. Its contents and
// attributes are copied from the currently active frame.
//
// The frame index indicates where the frame should be inserted. Valid values
// range from 0 to the current canvas frame count. If the frame index is greater
// than or equals the current canvas frame count, the new frame is appended at
// the end of the canvas. If the frame index is less than zero, the new frame is
// inserted at index 0.
//
// The active frame does not change, but its index may be renumbered due to the
// insertion.
//
// If an error occurs the according errno is returned.
func (cv Canvas) CreateFrame(id int) error {
	ret, err := C.caca_create_frame(cv.Cv, C.int(id))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// FreeFrame deletes a frame from a given canvas.
//
// The frame index indicates the frame to delete. Valid values range from 0 to
// the current canvas frame count minus 1. If the frame index is greater than or
// equals the current canvas frame count, the last frame is deleted.
//
// If the active frame is deleted, frame 0 becomes the new active frame.
// Otherwise, the active frame does not change, but its index may be renumbered
// due to the deletion.
//
// If an error occurs the according errno is returned.
func (cv Canvas) FreeFrame(id int) error {
	ret, err := C.caca_free_frame(cv.Cv, C.int(id))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// ImportFromMemory imports a memory buffer into the given libcaca canvas's
// current frame. The current frame is resized accordingly and its contents are
// replaced with the imported data.
//
// Valid values for format are:
//
//     "": attempt to autodetect the file format.
//     "caca": import native libcaca files.
//     "text": import ASCII text files.
//     "ansi": import ANSI files.
//     "utf8": import UTF-8 files with ANSI colour codes.
//     "bin": import BIN files.
//
// The number of bytes read is returned. If the file format is valid, but not
// enough data was available, 0 is returned.
//
// If an error occurs -1 and the according errno is returned.
func (cv Canvas) ImportFromMemory(data []byte, format string) (int, error) {
	l := C.size_t(len(data))
	ret, err := C.caca_import_canvas_from_memory(cv.Cv, unsafe.Pointer(&data[0]), l, C.CString(format))

	if int(ret) == -1 {
		return -1, err
	}

	return int(ret), nil
}

// ImportFromFile imports a file into the given libcaca canvas's current frame.
// The current frame is resized accordingly and its contents are replaced with
// the imported data.
//
// Valid values for format are:
//
//     "": attempt to autodetect the file format.
//     "caca": import native libcaca files.
//     "text": import ASCII text files.
//     "ansi": import ANSI files.
//     "utf8": import UTF-8 files with ANSI colour codes.
//     "bin": import BIN files.
//
// The number of bytes read is returned. If the file format is valid, but not
// enough data was available, 0 is returned.
//
// If an error occurs -1 and the according errno is returned.
func (cv Canvas) ImportFromFile(filename string, format string) (int, error) {
	ret, err := C.caca_import_canvas_from_file(cv.Cv, C.CString(filename), C.CString(format))

	if int(ret) == -1 {
		return -1, err
	}

	return int(ret), nil
}

// ImportAreaFromMemory imports a memory buffer into a canvas area.
//
// Import a memory buffer into the given libcaca canvas's current frame, at the
// specified position. For more information, see ImportFromMemory().
//
// The number of bytes read is returned. If the file format is valid, but not
// enough data was available, 0 is returned.
//
// If an error occurs -1 and the according errno is returned.
func (cv Canvas) ImportAreaFromMemory(x int, y int, data []byte, format string) (int, error) {
	l := C.size_t(len(data))
	ret, err := C.caca_import_area_from_memory(cv.Cv, C.int(x), C.int(y), unsafe.Pointer(&data[0]), l, C.CString(format))

	if int(ret) == -1 {
		return -1, err
	}

	return int(ret), nil
}

// ImportAreaFromFile imports a memory buffer into a canvas area.
//
// Import a file into the given libcaca canvas's current frame, at the
// specified position. For more information, see ImportFromFile().
//
// The number of bytes read is returned. If the file format is valid, but not
// enough data was available, 0 is returned.
//
// If an error occurs -1 and the according errno is returned.
func (cv Canvas) ImportAreaFromFile(x int, y int, filename string, format string) (int, error) {
	ret, err := C.caca_import_area_from_file(cv.Cv, C.int(x), C.int(y), C.CString(filename), C.CString(format))

	if int(ret) == -1 {
		return -1, err
	}

	return int(ret), nil
}

// ExportToMemory This function exports a libcaca canvas into various foreign
// formats such as ANSI art, HTML, IRC colours, etc. The returned pointer should
// be passed to free() to release the allocated storage when it is no longer
// needed.
//
// Valid values for format are:
//
//     "caca": export native libcaca files.
//     "ansi": export ANSI art (CP437 charset with ANSI colour codes).
//     "html": export an HTML page with CSS information.
//     "html3": export an HTML table that should be compatible with most
//              navigators, including textmode ones.
//     "irc": export UTF-8 text with mIRC colour codes.
//     "ps": export a PostScript document.
//     "svg": export an SVG vector image.
//     "tga": export a TGA image.
//     "troff": export a troff source.
//
// If an error occurs an empty byte slice and the according errno is returned.
func (cv Canvas) ExportToMemory(format string) ([]byte, error) {
	var b C.size_t
	ret, err := C.caca_export_canvas_to_memory(cv.Cv, C.CString(format), &b)

	if ret == nil {
		return []byte{}, err
	}

	return C.GoBytes(ret, C.int(b)), nil
}

// ExportAreaToMemory exports a portion of a libcaca canvas into various formats.
// For more information, see ExportToMemory().
//
// If an error occurs an empty byte slice and the according errno is returned.
func (cv Canvas) ExportAreaToMemory(x int, y int, w int, h int, format string) ([]byte, error) {
	var b C.size_t
	ret, err := C.caca_export_area_to_memory(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h), C.CString(format), &b)

	if ret == nil {
		return []byte{}, err
	}

	return C.GoBytes(ret, C.int(b)), nil
}

// SetFigfont loads a figfont and attaches it to a canvas.
func (cv Canvas) SetFigfont(filename string) int {
	return int(C.caca_canvas_set_figfont(cv.Cv, C.CString(filename)))
}

// PutFigchar pastes a character using the current figfont.
func (cv Canvas) PutFigchar(ch rune) int {
	return int(C.caca_put_figchar(cv.Cv, C.uint32_t(ch)))
}

// FlushFiglet flushes the figlet context.
func (cv Canvas) FlushFiglet() int {
	return int(C.caca_flush_figlet(cv.Cv))
}

// Free frees all resources allocated by CreateCanvas(). The canvas pointer
// becomes invalid and must no longer be used unless a new call to
// CreateCanvas() is made.
//
// If an error occurs the according errno is returned.
func (cv Canvas) Free() error {
	ret, err := C.caca_free_canvas(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}
