#! /bin/bash

go build .
# GIN_MODE=release ./start -TEST=false -ENV=development
GIN_MODE=release ./start -TEST=false -ENV=local