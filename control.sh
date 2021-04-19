#!/bin/bash
function ping_api {
    for i in {1..5..2}
    do
        response=$(curl -s -o /dev/null -w "%{http_code}" http://127.0.0.1:8080/ping)
        echo "Ping response: $response"
        if [ "$response" = "200" ]; then
            echo "API has started...Ping is responding"
            exit 0
        fi
        sleep $i
    done
    echo "API not responding after 3 attempts. Stopping..."
    exit 1
}


if [ "$1" = "start" ]; then
    docker-compose up --detach
    ping_api
elif [ "$1" = "stop" ]; then
    docker-compose down
else
    echo "No command found"
fi
