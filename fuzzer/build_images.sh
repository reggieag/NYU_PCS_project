#!/bin/bash
(cd ./modules/api && ./build.sh)
(cd ./modules/sql_injection && ./build.sh)
(cd ./modules/oauth2_scopes && ./build.sh)
