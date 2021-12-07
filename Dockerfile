FROM golang:1.17.4

RUN apt update && \
  apt-get install -y libmagickwand-dev

WORKDIR /work
