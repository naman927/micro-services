FROM golang:1.17.12-alpine

WORKDIR $GOPATH/src/github.com/store

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["store"]