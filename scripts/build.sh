#!/bin/bash

rm -f -r ./dist
mkdir -p ./dist
sh ./scripts/update-swagger.sh
go mod download

if [ "$(expr substr $(uname -s) 1 5)" == "MINGW" ]; then
    go build -o ./dist/$(basename "$PWD").exe ./cmd 
else
    go build -o ./dist/$(basename "$PWD") ./cmd 
fi

cp -r ./assets ./dist
cp -r ./configs ./dist
