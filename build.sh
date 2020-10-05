#!/bin/bash
go build main.go
docker build -t percolator:v0.0.1 .
