FROM golang:1.20.1

LABEL maintainer="Douglas <odouglas.dev@gmail.com>"

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT ${PORT}

RUN go build

CMD tail -f /dev/null && ["./gomux-rest-api"]