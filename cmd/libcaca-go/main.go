package main

import (
	"github.com/czwinzscher/libcaca-go"
	"log"
)

func main() {
	dp, err := caca.CreateDisplay(nil)
	if err != nil {
		log.Fatal(err)
	}

	cv := dp.GetCanvas()
	dp.SetTitle("teitel")
	cv.SetColorAnsi(caca.COLOR_GREEN, caca.COLOR_BLACK)
	cv.PutStr(0, 0, "hallohl")
	dp.Refresh()
	dp.GetEvent(caca.EVENT_KEY_PRESS, nil, -1)

	defer dp.Free()
}
