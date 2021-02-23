package main

import (
	"errors"
	"fmt"
	"github.com/czwinzscher/libcaca-go"
	"os"
)

func showCanvas() (errs []error) {
	canvasWidth := 50
	canvasHeight := 10

	cv, err := caca.CreateCanvas(canvasWidth, canvasHeight)
	if err != nil {
		return []error{errors.New("error while creating canvas: " + err.Error())}
	}

	defer func() {
		if err = cv.Free(); err != nil {
			errs = append(errs, errors.New("error while releasing memory: "+err.Error()))
		}
	}()

	err = cv.SetColorAnsi(caca.ColorGreen, caca.ColorBlack)
	if err != nil {
		return []error{errors.New("error while setting colors: " + err.Error())}
	}

	cv.PutStr(0, 0, "hello world")

	dp, err := caca.CreateDisplay(&cv)
	if err != nil {
		return []error{errors.New("error while creating display: " + err.Error())}
	}

	err = dp.SetTitle("title")
	if err != nil {
		return []error{errors.New("error while setting title: " + err.Error())}
	}

	dp.Refresh()
	dp.GetEvent(caca.EventKeyPress, nil, -1)

	dp.Free()

	return errs
}

func main() {
	errs := showCanvas()

	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintln(os.Stderr, "[ERR] "+e.Error())
		}

		os.Exit(1)
	}
}
