FROM golang:latest

ENV GO111MODULE=on
WORKDIR /go/src/bbs

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build .

EXPOSE 403
ENTRYPOINT ["./bbs"]
