package caca

// #cgo LDFLAGS: -lcaca
// #include <caca.h>
// #include <stdlib.h>
import "C"

// Display is a libcaca display context.
type Display struct {
	Dp *C.struct_caca_display
}

// CreateDisplay creates a graphical context using device-dependent features
// (ncurses for terminals, an X11 window, a DOS command window...) that attaches
// to a libcaca canvas. Everything that gets drawn in the libcaca canvas can
// then be displayed by the libcaca driver.
//
// If no caca canvas is provided, a new one is created. Its handle can be
// retrieved using GetCanvas() and it is automatically destroyed when
// Free() is called.
//
// See also CreateDisplayWithDriver().
//
// If an error occurs the according errno is returned.
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

// CreateDisplayWithDriver creates a graphical context using device-dependent
// features (ncurses for terminals, an X11 window, a DOS command window...)
// that attaches to a libcaca canvas. Everything that gets drawn in the
// libcaca canvas can then be displayed by the libcaca driver.
//
// If no caca canvas is provided, a new one is created. Its handle can be
// retrieved using GetCanvas() and it is automatically destroyed when
// Free() is called.
//
// If no driver name is provided, libcaca will try to autodetect the best
// output driver it can.
//
// See also CreateDisplay().
//
// If an error occurs the according errno is returned.
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

// GetDriver returns the display's current output driver.
func (d Display) GetDriver() string {
	return C.GoString(C.caca_get_display_driver(d.Dp))
}

// SetDriver dynamically changes the display's output driver.
//
// Returns 0 in case of success, -1 if an error occurred.
func (d Display) SetDriver(driver string) int {
	return int(C.caca_set_display_driver(d.Dp, C.CString(driver)))
}

// GetCanvas returns the canvas that was either attached or created by
// CreateDisplay().
func (d Display) GetCanvas() Canvas {
	cPtr := C.caca_get_canvas(d.Dp)

	return Canvas{Cv: cPtr}
}

// Refresh flushes all graphical operations and prints them to the display
// device. Nothing will show on the screen until this function is called.
//
// If SetTime() was called with a non-zero value, Refresh() will use that
// value to achieve constant framerate: if two consecutive calls to Refresh()
// are within a time range shorter than the value set with SetTime(), the
// second call will be delayed before performing the screen refresh.
func (d Display) Refresh() {
	C.caca_refresh_display(d.Dp)
}

// SetTime sets the refresh delay in microseconds. The refresh delay is used by
// Refresh() to achieve constant framerate. See the Refresh() documentation
// for more details.
//
// If the argument is zero, constant framerate is disabled. This is the default
// behaviour.
//
// If an error occurs the according errno is returned.
func (d Display) SetTime(usec int) error {
	ret, err := C.caca_set_display_time(d.Dp, C.int(usec))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetTime returns the average rendering time, which is the average measured
// time between two Refresh() calls, in microseconds. If constant framerate was
// activated by calling SetTime(), the average rendering time will be close to
// the requested delay even if the real rendering time was shorter.
func (d Display) GetTime() int {
	return int(C.caca_get_display_time(d.Dp))
}

// SetTitle tries to change libcaca's window title if it runs in a window.
// This works with the ncurses, S-Lang, OpenGL, X11 and Win32 drivers.
//
// If an error occurs the according errno is returned.
func (d Display) SetTitle(title string) error {
	ret, err := C.caca_set_display_title(d.Dp, C.CString(title))

	if int(ret) != -1 {
		return nil
	}

	return err
}

// SetMouse shows or hides the mouse pointer. This function works with the
// ncurses, S-Lang and X11 drivers.
//
// If an error occurs the according errno is returned.
func (d Display) SetMouse(flag int) error {
	ret, err := C.caca_set_mouse(d.Dp, C.int(flag))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// SetCursor shows or hides the cursor, for devices that support such a feature.
//
// 0 hides the cursor, 1 shows the system's default cursor (usually a white
// rectangle). Other values are reserved for future use.
//
// If an error occurs the according errno is returned.
func (d Display) SetCursor(flag int) error {
	ret, err := C.caca_set_cursor(d.Dp, C.int(flag))

	if int(ret) == -1 {
		return err
	}

	return nil
}

// GetEvent polls the event queue for mouse or keyboard events matching the
// event mask and returns the first matching event. Non-matching events are
// discarded. If eventMask is zero, the function returns immediately.
//
// The timeout value tells how long this function needs to wait for an event.
// A value of zero returns immediately and the function returns zero if no more
// events are pending in the queue. A negative value causes the function to wait
// indefinitely until a matching event is received.
//
// If not nil, ev will be filled with information about the event received. If
// nil, the function will return but no information about the event will be
// sent.
func (d Display) GetEvent(eventMask int, ev *Event, timeout int) {
	if ev == nil {
		C.caca_get_event(d.Dp, C.int(eventMask), nil, C.int(timeout))
	} else {
		C.caca_get_event(d.Dp, C.int(eventMask), ev.Ev, C.int(timeout))
	}
}

// GetMouseX returns the X coordinate of the mouse position last time it was
// detected. This function is not reliable if the ncurses or S-Lang drivers are
// being used, because mouse position is only detected when the mouse is
// clicked. Other drivers such as X11 work well.
func (d Display) GetMouseX() int {
	return int(C.caca_get_mouse_x(d.Dp))
}

// GetMouseY returns the Y coordinate of the mouse position last time it was
// detected. This function is not reliable if the ncurses or S-Lang drivers are
// being used, because mouse position is only detected when the mouse is
// clicked. Other drivers such as X11 work well.
func (d Display) GetMouseY() int {
	return int(C.caca_get_mouse_y(d.Dp))
}

// Free detaches a graphical context from its caca backend and destroys it.
// The libcaca canvas continues to exist and other graphical contexts can be
// attached to it afterwards.
//
// If the caca canvas was automatically created by CreateDisplay(), it is
// automatically destroyed and any handle to it becomes invalid.
func (d Display) Free() {
	C.caca_free_display(d.Dp)
}
