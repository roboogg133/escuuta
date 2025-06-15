package main

import (
	"escuuta/srcYoutube"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

var widgetClick widget.Clickable
var config, _ = InternalFile("config.json")

func run(window *app.Window) error {
	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:

			// videoTitle, _ := srcYoutube.GetTitle("https://www.youtube.com/watch?v=wdSPqru3NDo")

			title := material.H1(theme, "oie")
			butao := material.Button(theme, &widgetClick, "baixa")

			// Change the color of the label.
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon
			butao.Color = maroon

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)
			butao.Layout(gtx)
			widgetClick.Clicked(gtx)

			// Pass the drawing operations to the GPU.
			if widgetClick.Pressed() {
				log.Println("furunfo: Button pressed!")

				audio, _ := InternalFile("audio.webm")

				err := srcYoutube.Download(config, "https://www.youtube.com/watch?v=uUZBAt00V6s", audio)
				if err != nil {
					log.Panic("furunfo: ", err)
				}
			}
			e.Frame(gtx.Ops)
			log.Println("furunfo: All loaded!")

		}
	}
}

func InternalFile(name string) (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, name), nil
}
