#!/bin/sh

# Create initial paths
mkdir -p example/foo example/bar

# Build and start go program with debug enabled
go build -tags debug -o notify .
./notify &

sleep 1

for i in {1..10}; do
    touch example/foo/a.txt
    touch example/bar/b.txt
    rm -rf example/foo example/bar
    sleep 1
    mkdir -p example/foo example/bar
done