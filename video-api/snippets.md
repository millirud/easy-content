# overlay

```ffmpeg -y -i video.MOV -i broken-2.png -filter_complex [0]overlay=x=0:y=300[out] -map [out] -map 0:a? test.mp4```

```ffmpeg -i video.MOV -i image.png -filter_complex "[0:v][1:v] overlay=25:25:enable='between(t, 0, 20)'" output.mov```

