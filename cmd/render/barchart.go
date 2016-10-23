package main

type Label struct {
	X    int
	Y    int
	Text string
}

type Bar struct {
	Width  int
	Height int
	X      int
	Y      int
	Label  Label
}

type Barchart struct {
	Width  int
	Height int
	Bars   []Bar
}
