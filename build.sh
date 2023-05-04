#!/bin/bash
echo "Cleaning previous builds..."
rm handler
echo "Building for Linux AMD64"
GOOS=linux GOARCH=amd64 go build handlers/handler.go
echo "New build created"
ls -lh handler