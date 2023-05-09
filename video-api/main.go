package main

import (
	"fmt"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type myWriter struct {
	data string
}

func (w *myWriter) Write(p []byte) (n int, err error) {
	w.data += string(p)

	println("AAAAAA", string(p))

	return len(p), nil
}

func main() {

	inputFile := "./data/video.MOV"

	a, err := ffmpeg.Probe(inputFile)
	if err != nil {
		panic(err)
	}

	println("probe", a)

	// overlay := ffmpeg.Input("./data/broken-2.png")
	// err := ffmpeg.Filter(
	// 	[]*ffmpeg.Stream{
	// 		ffmpeg.Input("./data/video.MOV"),
	// 		overlay,
	// 	}, "overlay", ffmpeg.Args{"10:10"}, ffmpeg.KwArgs{"enable": "gte(t,1)"}).
	// 	Output("./data/out1.mp4").
	// 	OverWriteOutput().
	// 	ErrorToStdOut().Run()

	w := &myWriter{}
	//e := &myWriter{}

	stream := ffmpeg.
		Input(inputFile, ffmpeg.KwArgs{
			"ss": "00:00:20.000",
			"to": "00:00:30.000",
		}).
		Output("./data/out1.MOV", ffmpeg.KwArgs{
			// "c:v":    "vp9",
			// "preset": "medium",
		}).
		GlobalArgs("-progress", "pipe:1").
		OverWriteOutput().
		WithOutput(w)

	fmt.Println("My cmd", "ffmpeg", strings.Join(stream.GetArgs(), " "))

	err = stream.Run()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("end")
}
