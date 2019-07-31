FROM golang:latest AS builder

ENV GO111MODULE=on
WORKDIR /go/src/bbs
COPY go.mod . 
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/bbs

FROM scratch

COPY --from=builder /go/bin/bbs /go/bin/bbs
EXPOSE 403
ENTRYPOINT ["/go/bin/bbs"]
