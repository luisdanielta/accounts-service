#!/bin/bash

while [ $# -gt 0 ]; do
    case "$1" in
        -o|--option)
            Option="$2"
            shift 2
            ;;
        *)
            break
            ;;
    esac
done

if [ -z "$Option" ]; then
    Option="go"
fi

source ./env.sh

if [ "$Option" = "go" ]; then
    go run app.go
    exit 0
elif [ "$Option" = "dev" ]; then
    echo "mode dev"
    air -c .air.toml
    exit 0
elif [ "$Option" = "build" ]; then
    echo "mode build"
    path="./build"

    if [ ! -d "$path" ]; then
        mkdir "build"
    fi

    go build -o "$path/app" app.go

    "$path/app"

    exit 0
fi
