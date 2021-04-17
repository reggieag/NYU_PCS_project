FROM golang:latest AS build
WORKDIR /src
COPY . /src
RUN cd /src && go build -o openapi

ENTRYPOINT ["./openapi"]
