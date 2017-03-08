#!/bin/bash -el

cd ./commands
go run updateSchema.go
cd ..
go build -o ./bin/bbservice
chmod 744 ./bin/bbservice
codesign -s $CERT ./bin/bbservice
