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

func (cv Canvas) SetSize(width int, height int) error {
	ret, err := C.caca_set_canvas_size(cv.Cv, C.int(width), C.int(height))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) GetWidth() int {
	return int(C.caca_get_canvas_width(cv.Cv))
}

func (cv Canvas) GetHeight() int {
	return int(C.caca_get_canvas_height(cv.Cv))
}

func (cv Canvas) GoToXY(x int, y int) {
	C.caca_gotoxy(cv.Cv, C.int(x), C.int(y))
}

func (cv Canvas) WhereX() int {
	return int(C.caca_wherex(cv.Cv))
}

func (cv Canvas) WhereY() int {
	return int(C.caca_wherey(cv.Cv))
}

func (cv Canvas) PutChar(x int, y int, ch rune) {
	C.caca_put_char(cv.Cv, C.int(x), C.int(y), C.uint32_t(ch))
}

func (cv Canvas) GetChar(x int, y int) rune {
	return rune(C.caca_get_char(cv.Cv, C.int(x), C.int(y)))
}

func (cv Canvas) PutStr(x int, y int, str string) {
	C.caca_put_str(cv.Cv, C.int(x), C.int(y), C.CString(str))
}

func (cv Canvas) Clear() {
	C.caca_clear_canvas(cv.Cv)
}

func (cv Canvas) SetHandle(x int, y int) {
	C.caca_set_canvas_handle(cv.Cv, C.int(x), C.int(y))
}

func (cv Canvas) GetHandleX() int {
	return int(C.caca_get_canvas_handle_x(cv.Cv))
}

func (cv Canvas) GetHandleY() int {
	return int(C.caca_get_canvas_handle_y(cv.Cv))
}

func (cv *Canvas) Blit(x int, y int, src Canvas, mask *Canvas) error {
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

func (cv Canvas) SetBoundaries(x int, y int, w int, h int) error {
	ret, err := C.caca_set_canvas_boundaries(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) DisableDirtyRect() {
	C.caca_disable_dirty_rect(cv.Cv)
}

func (cv Canvas) EnableDirtyRect() {
	C.caca_enable_dirty_rect(cv.Cv)
}

func (cv Canvas) GetDirtyRectCount() int {
	return int(C.caca_get_dirty_rect_count(cv.Cv))
}

func (cv Canvas) GetDirtyRect(idx int) (map[string]int, error) {
	var x, y, width, height C.int

	ret, err := C.caca_get_dirty_rect(cv.Cv, C.int(idx), &x, &y, &width, &height)

	if int(ret) == -1 {
		return map[string]int{}, err
	}

	return map[string]int{"x": int(x), "y": int(y), "width": int(width), "height": int(height)}, nil
}

func (cv Canvas) AddDirtyRect(x int, y int, width int, height int) error {
	ret, err := C.caca_add_dirty_rect(cv.Cv, C.int(x), C.int(y), C.int(width), C.int(height))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) RemoveDirtyRect(x int, y int, width int, height int) error {
	ret, err := C.caca_remove_dirty_rect(cv.Cv, C.int(x), C.int(y), C.int(width), C.int(height))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) ClearDirtyRectList() {
	C.caca_clear_dirty_rect_list(cv.Cv)
}

func (cv Canvas) Invert() {
	C.caca_invert(cv.Cv)
}

func (cv Canvas) Flip() {
	C.caca_flip(cv.Cv)
}

func (cv Canvas) Flop() {
	C.caca_flop(cv.Cv)
}

func (cv Canvas) Rotate180() {
	C.caca_rotate_180(cv.Cv)
}

func (cv Canvas) RotateLeft() error {
	ret, err := C.caca_rotate_left(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) RotateRight() error {
	ret, err := C.caca_rotate_right(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) StretchLeft() error {
	ret, err := C.caca_stretch_left(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) StretchRight() error {
	ret, err := C.caca_stretch_right(cv.Cv)

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) GetAttr(x int, y int) rune {
	return rune(C.caca_get_attr(cv.Cv, C.int(x), C.int(y)))
}

func (cv Canvas) SetAttr(attr rune) {
	C.caca_set_attr(cv.Cv, C.uint32_t(attr))
}

func (cv Canvas) UnsetAttr(attr rune) {
	C.caca_unset_attr(cv.Cv, C.uint32_t(attr))
}

func (cv Canvas) ToggleAttr(attr rune) {
	C.caca_toggle_attr(cv.Cv, C.uint32_t(attr))
}

func (cv Canvas) PutAttr(x int, y int, attr rune) {
	C.caca_put_attr(cv.Cv, C.int(x), C.int(y), C.uint32_t(attr))
}

func (cv Canvas) SetColorAnsi(fg byte, bg byte) error {
	ret, err := C.caca_set_color_ansi(cv.Cv, C.uint8_t(fg), C.uint8_t(bg))

	if int(ret) != -1 {
		return nil
	}

	return err
}

func (cv Canvas) SetColorArgb(fg int16, bg int16) {
	C.caca_set_color_argb(cv.Cv, C.uint16_t(fg), C.uint16_t(bg))
}

func (cv Canvas) DrawLine(x1 int, y1 int, x2 int, y2 int, ch rune) {
	C.caca_draw_line(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.uint32_t(ch))
}

// DrawPolyline draws a polyline on the canvas using the given character and
// the given coordinates.
func (cv Canvas) DrawPolyline(xy []IntPair, ch rune) {
	var cx, cy []C.int

	for _, e := range xy {
		cx = append(cx, C.int(e.First))
		cy = append(cy, C.int(e.Second))
	}

	C.caca_draw_polyline(cv.Cv, &cx[0], &cy[0], C.int(len(xy)-1), C.uint32_t(ch))
}

// DrawThinLine draws a thin line on the canvas from (x1, y1) to (x2, y2).
func (cv Canvas) DrawThinLine(x1 int, y1 int, x2 int, y2 int) {
	C.caca_draw_thin_line(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

// DrawThinPolyline draws a thin polyline on the canvas using the given coordinates.
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

func (cv Canvas) DrawEllipse(xo int, yo int, a int, b int, ch rune) {
	C.caca_draw_ellipse(cv.Cv, C.int(xo), C.int(yo), C.int(a), C.int(b), C.uint32_t(ch))
}

func (cv Canvas) DrawThinEllipse(xo int, yo int, a int, b int) {
	C.caca_draw_thin_ellipse(cv.Cv, C.int(xo), C.int(yo), C.int(a), C.int(b))
}

func (cv Canvas) FillEllipse(xo int, yo int, a int, b int, ch rune) {
	C.caca_fill_ellipse(cv.Cv, C.int(xo), C.int(yo), C.int(a), C.int(b), C.uint32_t(ch))
}

func (cv Canvas) DrawBox(x int, y int, w int, h int, ch rune) {
	C.caca_draw_box(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h), C.uint32_t(ch))
}

func (cv Canvas) DrawThinBox(x int, y int, w int, h int) {
	C.caca_draw_thin_box(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (cv Canvas) DrawCP437Box(x int, y int, w int, h int) {
	C.caca_draw_cp437_box(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h))
}

func (cv Canvas) FillBox(x int, y int, w int, h int, ch rune) {
	C.caca_fill_box(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h), C.uint32_t(ch))
}

func (cv Canvas) DrawTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int, ch rune) {
	C.caca_draw_triangle(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3), C.uint32_t(ch))
}

func (cv Canvas) DrawThinTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int) {
	C.caca_draw_thin_triangle(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3))
}

func (cv Canvas) FillTriangle(x1 int, y1 int, x2 int, y2 int, x3 int, y3 int, ch rune) {
	C.caca_fill_triangle(cv.Cv, C.int(x1), C.int(y1), C.int(x2), C.int(y2), C.int(x3), C.int(y3), C.uint32_t(ch))
}

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

func (cv Canvas) GetFrameCount() int {
	return int(C.caca_get_frame_count(cv.Cv))
}

func (cv Canvas) SetFrame(id int) error {
	ret, err := C.caca_set_frame(cv.Cv, C.int(id))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) GetFrameName() string {
	return C.GoString(C.caca_get_frame_name(cv.Cv))
}

func (cv Canvas) SetFrameName(name string) error {
	ret, err := C.caca_set_frame_name(cv.Cv, C.CString(name))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) CreateFrame(id int) error {
	ret, err := C.caca_create_frame(cv.Cv, C.int(id))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) FreeFrame(id int) error {
	ret, err := C.caca_free_frame(cv.Cv, C.int(id))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) ImportFromMemory(data []byte, length uint, format string) error {
	ret, err := C.caca_import_canvas_from_memory(cv.Cv, unsafe.Pointer(&data[0]), C.size_t(length), C.CString(format))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) ImportFromFile(filename string, format string) error {
	ret, err := C.caca_import_canvas_from_file(cv.Cv, C.CString(filename), C.CString(format))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) ImportAreaFromMemory(x int, y int, data []byte, length uint, format string) error {
	ret, err := C.caca_import_area_from_memory(cv.Cv, C.int(x), C.int(y), unsafe.Pointer(&data[0]), C.size_t(length), C.CString(format))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) ImportAreaFromFile(x int, y int, filename string, format string) error {
	ret, err := C.caca_import_area_from_file(cv.Cv, C.int(x), C.int(y), C.CString(filename), C.CString(format))

	if int(ret) == -1 {
		return err
	}

	return nil
}

func (cv Canvas) ExportToMemory(format string) ([]byte, error) {
	var b C.size_t
	ret, err := C.caca_export_canvas_to_memory(cv.Cv, C.CString(format), &b)

	if ret == nil {
		return []byte{}, err
	}

	return C.GoBytes(ret, C.int(b)), nil
}

func (cv Canvas) ExportAreaToMemory(x int, y int, w int, h int, format string) ([]byte, error) {
	var b C.size_t
	ret, err := C.caca_export_area_to_memory(cv.Cv, C.int(x), C.int(y), C.int(w), C.int(h), C.CString(format), &b)

	if ret == nil {
		return []byte{}, err
	}

	return C.GoBytes(ret, C.int(b)), nil
}

func (cv Canvas) SetFigfont(filename string) int {
	return int(C.caca_canvas_set_figfont(cv.Cv, C.CString(filename)))
}

func (cv Canvas) PutFigchar(ch rune) int {
	return int(C.caca_put_figchar(cv.Cv, C.uint32_t(ch)))
}

func (cv Canvas) FlushFiglet() int {
	return int(C.caca_flush_figlet(cv.Cv))
}

// func (cv Canvas) Render(font Font, buf []byte, width int, height int, pitch int) error {
// 	ret, err := C.caca_render_canvas(cv.Cv, C.CString(font), unsafe.Pointer(&buf[0]), C.int(width), C.int(height), C.int(pitch))

// 	if int(ret) == -1 {
// 		return err
// 	}

// 	return nil
// }

func (cv Canvas) Free() {
	C.free(unsafe.Pointer(cv.Cv))
}
