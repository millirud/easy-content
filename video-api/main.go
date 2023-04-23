package main

import (
	"fmt"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {

	overlay := ffmpeg.Input("./data/broken-2.png")
	err := ffmpeg.Filter(
		[]*ffmpeg.Stream{
			ffmpeg.Input("./data/video.MOV"),
			overlay,
		}, "overlay", ffmpeg.Args{"10:10"}, ffmpeg.KwArgs{"enable": "gte(t,1)"}).
		Output("./data/out1.mp4").OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		fmt.Println(err)
	}
}
