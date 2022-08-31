package main

import (
	"image"
	"image/color"

	"overdrive/geometry"
	"overdrive/mesh"
	"overdrive/render"
	"overdrive/utilities"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"fmt"
	// "sync"
	"time"
)

func main() {
	// text1 := canvas.NewText("1", color.White)
	// textFps := canvas.NewText("2", color.White)

	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myCanvas := myWindow.Canvas()

	img := image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	viewport := canvas.NewImageFromImage(img)

	// grid := container.New(layout.NewGridLayout(2), viewport, textFps, content)

	bottom := widget.NewButton("Assets browser", func() {
		fmt.Println("tapped")
	})

	right := canvas.NewText("fps", color.White)
	// middle := canvas.NewText("content", color.White)
	content := container.New(layout.NewBorderLayout(nil, bottom, nil, right),
		bottom, right, viewport)

	myCanvas.SetContent(content)

	go func() {

		cam := render.Camera{
			Position: geometry.NewVector(0, 0, -100),
			Rotation: geometry.ZeroVector()}
		light := render.Light{
			Position:  geometry.NewVector(0, 0, 0),
			Rotation:  geometry.ZeroVector(),
			LightType: render.Point,
			Color:     color.RGBA{0, 255, 255, 255},
			Length:    1000,
		}

		start := time.Now()

		// objects := make([]mesh.Mesh, 10)
		cube := mesh.Cube(geometry.NewVector(0, 0, 0), geometry.ZeroVector(), geometry.NewVector(400, 400, 400))

		suzanne1 := mesh.ReadObjFile()
		suzanne2 := mesh.ReadObjFile()

		suzanne1.Translate(geometry.NewVector(200, 0, 0))
		suzanne2.Translate(geometry.NewVector(-200, 0, 0))

		for {

			img = image.NewRGBA(image.Rect(0, 0, utilities.RESOLUTION_X, utilities.RESOLUTION_Y))

			// for x := 0; x < utilities.RESOLUTION_X; x++ {
			// 	for y := 0; y < utilities.RESOLUTION_Y; y++ {
			// 		img.Set(x, y, color.Black)
			// 	}
			// }

			// for i := range objects {
			// 	objects[i] = suzanne1
			// }

			// wg := sync.WaitGroup{}
			// for i := range objects {
			// 	wg.Add(1)
			// 	go func(i int) {
			// 		objects[i].Draw(img, cam, []render.Light{light})
			// 		wg.Done()
			// 	}(i)
			// }
			// wg.Wait()

			suzanne1.Draw(img, cam, []render.Light{light})
			suzanne2.Draw(img, cam, []render.Light{light})

			// suzanne.Translate(geometry.NewVector(0, 0, -0.1))
			
			// suzanne1.Rotate(geometry.NewVector(0, 0.01, 0))

			cube.Translate(geometry.NewVector(0, 0, 1))
			cube.Rotate(geometry.NewVector(0, 0.01, 0))

			viewport.Image = img
			viewport.Refresh()

			t := time.Since(start).Milliseconds()
			if t == 0 {
				t = 1
			}

			right.Text = fmt.Sprint("fps : ", 1000/t)
			right.Refresh()

			start = time.Now()
			// break
		}
	}()

	myWindow.Resize(fyne.NewSize(utilities.RESOLUTION_X, utilities.RESOLUTION_Y))
	myWindow.ShowAndRun()
}



			//TODO Aberty666