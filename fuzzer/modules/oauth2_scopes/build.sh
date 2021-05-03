#!/bin/bash

cp -rf ../api ./

docker build . --file Dockerfile --tag fuzzer-modules-oauth2_scopes:latest 
