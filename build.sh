#!/bin/bash

cd bin

export CGO_ENABLED=1

go build -a ../src/worker.go
go build -a ../src/producer.go
