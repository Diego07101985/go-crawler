FROM golang:latest 

ENV SRC_DIR=/go/src/go-crawler/search-api
ADD . $SRC_DIR
RUN cd $SRC_DIR

RUN go get github.com/gin-gonic/gin && go get github.com/go-sql-driver/mysql 
RUN go get  github.com/olivere/elastic
RUN go get  github.com/jinzhu/gorm/dialects/mysql
RUN go get  github.com/teris-io/shortid
RUN go get  github.com/jessevdk/go-assets
RUN go get github.com/jinzhu/inflection
RUN go get github.com/lib/pq
RUN go get github.com/stretchr/testify/assert

WORKDIR  $SRC_DIR
RUN go build -o main


CMD ["./main"]
