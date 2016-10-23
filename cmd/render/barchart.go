package main

type Label struct {
	X      int
	Y      int
	Text   string
	Rotate int
}

type Bar struct {
	Width  int
	Height int
	X      int
	Y      int
	Label  Label
	YLabel Label
}

type Barchart struct {
	Width  int
	Height int
	Bars   []Bar
}
