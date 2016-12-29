#!/bin/bash

go build -o yimp .
docker build -t localhost:5000/yimp .
