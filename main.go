package main

import "gopkg.in/alecthomas/kingpin.v1"

var (
	Version = "0.0.1"
	app     = kingpin.New("goed", "A code editor")
	test    = kingpin.Flag("testterm", "Pints colors to the terminal to test it.").Bool()
	colors  = kingpin.Flag("c", "Number of colors(0,2,16,256). 0 means Detect.").Default("0").Int()

	Ed Editor
)

func main() {
	kingpin.Version("0.0.1")

	kingpin.Parse()
	if *test {
		testTerm()
		return
	}
	if *colors == 0 {
		*colors = detectColors()
	}
	if *colors != 256 && *colors != 16 {
		*colors = 2
	}

	Ed = Editor{}
	Ed.Start()
}