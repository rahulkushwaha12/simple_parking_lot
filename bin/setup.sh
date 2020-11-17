#!/bin/bash

echo "formatting the code..."
go fmt ../cmd/...
go fmt ../internal/...
echo "done."

sh ./run_functional_tests.sh

echo "building binary..."
go build -o parking_lot ../cmd/parking_lot/main.go
echo "done."