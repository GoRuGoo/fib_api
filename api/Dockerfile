FROM golang:latest

RUN mkdir /go/src/api
WORKDIR /go/src/api

COPY . /go/src/api
RUN go mod download

RUN go build -o main .

EXPOSE 8080
CMD ["./main"]
