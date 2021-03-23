FROM golang:latest AS build
WORKDIR /src
COPY . /src
RUN go build -o openapi /src/main.go 

ENTRYPOINT ["./openapi"]
