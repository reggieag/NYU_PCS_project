#!/bin/bash
docker build --no-cache . --file Dockerfile --tag fuzzer-modules-sql_injection:latest
