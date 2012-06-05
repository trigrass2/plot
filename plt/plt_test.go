package plt

import (
	"code.google.com/p/plotinum/vg"
	"code.google.com/p/plotinum/vg/veceps"
	"code.google.com/p/plotinum/vg/vecimg"
	"testing"
)

func TestDrawImage(t *testing.T) {
	img, err := vecimg.New(vg.Inches(4), vg.Inches(4))
	if err != nil {
		t.Fatal(err)
	}

	da := &DrawArea{
		Canvas: img,
		rect: rect{Min: point{0, 0},
			Size: point{vg.Inches(4), vg.Inches(4)},
		},
	}
	draw(da)
	err = img.SavePNG("test.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDrawEps(t *testing.T) {
	eps := veceps.New(vg.Inches(4), vg.Inches(4), "test")
	da := &DrawArea{
		Canvas: eps,
		rect: rect{Min: point{0, 0},
			Size: point{vg.Inches(4), vg.Inches(4)},
		},
	}
	draw(da)
	err := eps.Save("test.eps")
	if err != nil {
		t.Fatal(err)
	}
}

// draw draws a simple test plot
func draw(da *DrawArea) {
	plot := NewPlot()
	plot.Title = "This is a plot"
	plot.XAxis.Min = 100000
	plot.XAxis.Max = 100001
	plot.XAxis.Label = "X-Axis gq"
	plot.YAxis.Min = 10
	plot.YAxis.Max = 20
	plot.YAxis.Ticks.TickMarker = ConstantTicks([]Tick{{10, "ten"}, {15, ""}, {20, "twenty"}})
	plot.YAxis.Label = "Y-Axis gq"
	plot.Draw(da)
}