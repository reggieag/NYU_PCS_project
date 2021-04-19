FROM golang:latest AS build
WORKDIR /src
COPY . /src
RUN cd /src && go build -o auth_service

ENTRYPOINT ["./auth_service", "/src/clients.yml"]
