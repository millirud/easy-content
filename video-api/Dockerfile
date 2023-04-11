FROM golang:1.20.3-buster as dev

WORKDIR /home/app

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y ffmpeg 

