#!/bin/bash

echo "formatting the code..."
go fmt ../cmd
echo "building binary..."
go build -o parking_lot ../cmd/parking_lot/main.go