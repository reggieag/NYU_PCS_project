#!/bin/bash

cp -rf ../api ./

docker build --no-cache . --file Dockerfile --tag fuzzer-modules-sql_injection:latest
