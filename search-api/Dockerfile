FROM golang:1.10.0

ENV SRC_DIR=/go/src/go-elasticsearch-example/search-api
ADD . $SRC_DIR
RUN cd $SRC_DIR

RUN go install -v ./...
WORKDIR  $SRC_DIR
RUN go build -o main


CMD ["./main"]
