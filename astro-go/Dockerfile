FROM golang:1.8.0

COPY . $GOPATH/src/github.com/astromechio/astro-go
WORKDIR $GOPATH/src/github.com/astromechio/astro-go

RUN go build

ENTRYPOINT ./astro-go