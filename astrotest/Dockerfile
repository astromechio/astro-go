FROM golang:1.8.0

COPY . $GOPATH/src/github.com/astromechio/astrotest
WORKDIR $GOPATH/src/github.com/astromechio/astrotest

RUN go build

ENTRYPOINT ./astrotest