FROM golang:latest

WORKDIR $GOPATH/src/go_bbs
COPY . $GOPATH/src/go_bbs
RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
ENTRYPOINT ["./bbs"]