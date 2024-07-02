FROM golang:1.22 AS build

WORKDIR /app
COPY main.go ./

EXPOSE 8090

RUN go build main.go
CMD ./main
