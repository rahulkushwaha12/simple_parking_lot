#!/bin/bash

echo "executing test cases..."
go test -v ../internal/...
echo "done."