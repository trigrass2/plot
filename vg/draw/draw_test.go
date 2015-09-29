package draw

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/recorder"
)

func TestCrop(t *testing.T) {
	ls := LineStyle{
		Color: color.NRGBA{0, 20, 0, 123},
		Width: 0.1 * vg.Inch,
	}
	var r1 recorder.Canvas
	c1 := NewCanvas(&r1, 6, 3)
	c11 := Crop(c1, 0, -3, 0, 0)
	c12 := Crop(c1, 3, 0, 0, 0)

	var r2 recorder.Canvas
	c2 := NewCanvas(&r2, 6, 3)
	c21 := Canvas{
		Canvas: c2.Canvas,
		Rectangle: Rectangle{
			Min: Point{0, 0},
			Max: Point{3, 3},
		},
	}
	c22 := Canvas{
		Canvas: c2.Canvas,
		Rectangle: Rectangle{
			Min: Point{3, 0},
			Max: Point{6, 3},
		},
	}
	str := "unexpected result: %+v != %+v"
	if c11.Rectangle != c21.Rectangle {
		t.Errorf(str, c11.Rectangle, c21.Rectangle)
	}
	if c12.Rectangle != c22.Rectangle {
		t.Errorf(str, c11.Rectangle, c21.Rectangle)
	}

	c11.StrokeLine2(ls, c11.Min.X, c11.Min.Y,
		c11.Min.X+3*vg.Inch, c11.Min.Y+3*vg.Inch)
	c12.StrokeLine2(ls, c12.Min.X, c12.Min.Y,
		c12.Min.X+3*vg.Inch, c12.Min.Y+3*vg.Inch)
	c21.StrokeLine2(ls, c21.Min.X, c21.Min.Y,
		c21.Min.X+3*vg.Inch, c21.Min.Y+3*vg.Inch)
	c22.StrokeLine2(ls, c22.Min.X, c22.Min.Y,
		c22.Min.X+3*vg.Inch, c22.Min.Y+3*vg.Inch)

	if !reflect.DeepEqual(r1.Actions, r2.Actions) {
		t.Errorf(str, r1.Actions, r2.Actions)
	}
}

func TestTile(t *testing.T) {
	var r recorder.Canvas
	c := NewCanvas(&r, 13, 7)
	const (
		rows = 2
		cols = 3
		pad  = 1
	)
	tiles := Tiles{
		Rows: rows, Cols: cols,
		PadTop: pad, PadBottom: pad,
		PadRight: pad, PadLeft: pad,
		PadX: pad, PadY: pad,
	}
	rectangles := [][]Rectangle{
		{
			Rectangle{
				Min: Point{1, 4},
				Max: Point{4, 6},
			},
			Rectangle{
				Min: Point{5, 4},
				Max: Point{8, 6},
			},
			Rectangle{
				Min: Point{9, 4},
				Max: Point{12, 6},
			},
		},
		{
			Rectangle{
				Min: Point{1, 1},
				Max: Point{4, 3},
			},
			Rectangle{
				Min: Point{5, 1},
				Max: Point{8, 3},
			},
			Rectangle{
				Min: Point{9, 1},
				Max: Point{12, 3},
			},
		},
	}
	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			str := "row %d col %d unexpected result: %+v != %+v"
			tile := tiles.At(c, i, j)
			if tile.Rectangle != rectangles[j][i] {
				t.Errorf(str, j, i, tile.Rectangle, rectangles[j][i])
			}
		}
	}
}
