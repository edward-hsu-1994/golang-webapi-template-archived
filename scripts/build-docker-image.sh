#!/bin/bash

docker build . -f ./build/Dockerfile -t $(basename "$PWD"):dev
