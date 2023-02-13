package ui

import "github.com/go-gl/glfw/v3.3/glfw"


type Column struct {
	Properties *Properties
	Children   []UIElement
}

func (column Column) Draw(screen []byte, window *glfw.Window) {
	
	column.Properties.Draw(screen, window)

	for child := range column.Children {
		
		column.Children[child].SetProperties(
			Size{
				Scale:  column.Properties.Size.Scale,
				Width:  column.Properties.Size.Width,
				Height: column.Properties.Size.Height / len(column.Children),
			},
			Point{
				X: column.Properties.Center.X,
				Y: column.Properties.Center.Y - column.Properties.MaxSize.Height/2 + (2*(len(column.Children) - child - 1)+1)*column.Properties.MaxSize.Height/(len(column.Children)*2),
			},
		)
		column.Children[child].Draw(screen, window)
	}
}


func (column Column) SetProperties(size Size, center Point) {
	column.Properties.MaxSize = size
	column.Properties.Center = center
}

func (column Column) Debug() {
	println(column.Properties.Center.Y)
}
