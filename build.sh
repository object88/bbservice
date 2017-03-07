#!/bin/bash -el

cd ./commands
go run updateSchema.go
cd ..
go build -o bin/bbservice
codesign -s $CERT bin/service
