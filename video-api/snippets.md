# overlay

```ffmpeg -y -i video.MOV -i broken-2.png -filter_complex [0]overlay=x=0:y=300[out] -map [out] -map 0:a? test.mp4```

```ffmpeg -i video.MOV -i image.png -filter_complex "[0:v][1:v] overlay=25:25:enable='between(t, 0, 20)'" output.mov```

ffmpeg -i ./data/video.MOV -lavfi '[0:v]scale=ih*16/9:-1,boxblur=luma_radius=min(h\,w)/20:luma_power=1:chroma_radius=min(cw\,ch)/20:chroma_power=1[bg];[bg][0:v]overlay=(W-w)/2:(H-h)/2,crop=h=iw*9/16' -vb 800k ./data/blur.mp4

ffmpeg -i ./data/oldFilm1080.mp4 -i ./data/video.MOV -filter_complex "[0]format=rgba,colorchannelmixer=aa=0.25[fg];[1][fg]overlay[out]" -map [out] -pix_fmt yuv420p -c:v libx264 -crf 18 ./data/touchdown-vintage.mp4