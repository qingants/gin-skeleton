#!/usr/bin/env bash

sh ./build.sh

docker build -t ibroomcorn/skeleton .
docker push ibroomcorn/skeleton:latest

rm ./skeleton