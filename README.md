# Overview
Fuzzy Wuzzy is an API fuzzing framework. 
Given a schema and its accompanying service, 
Fuzzy Wuzzy will run various modules against the service.


Individual modules are implemented in Docker containers, making Fuzzy Wuzzy easily extendible.

This repository contains several components:
1. Example API Service
2. Postgres Database Schema backing service
3. Example OAuth2 server
4. Fuzzy Wuzzy Fuzzer

## Requirements

All components in this repo will require `docker` to be installed on the host machine. 
Fuzzy Wuzzy itself requires `go 1.16.x` to build. Some of the helper scripts will require `bash` or some
equivalent to run, but they can easily be run manually.

## Folder Organization

[service](./service) contains the example service implementation, as well as the OpenAPI spec file that the service
implements


[database](./database) contains the schema for the example service database.

[auth_service](./auth_service) contains a the OAuth2 server implementation

[fuzzer](./fuzzer) contains the `Fuzzy Wuzzy` fuzzer itself.

## Quickstart

To start or stop the example services, the [control script](./control.sh) can be used.
Simply run `bash control.sh start` and `bash control.sh stop` to start or stop the services, respectively.
The OAuth2 server will be available at `port 9096` and the actual API service will be available at `port 8080`.
The database is exposed on `port 5324`.


Once `Fuzzy Wuzzy` has been built, it can be run as any other executable. 
Please see [the Fuzzy Wuzzy README](./fuzzer/README.md) to build and run `Fuzzy Wuzzy`.
