#!/bin/bash
if [ "$1" = "start" ]; then
    docker-compose up --detach
elif [ "$1" = "stop" ]; then
    docker-compose down
else
    echo "No command found"
fi
