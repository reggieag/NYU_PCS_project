# Overview
We'll add stuff here later...

## RESTler
RESTler is built in a docker container and can be accessed using the `restler.sh` script.

To build the container.
```
cd restler-fuzzer
docker build -t microsoft/restler-fuzzer .
```

Example run RESTler from the command line.
```
wget https://petstore.swagger.io/v2/swagger.json
./restler.sh compile swagger.json
```
