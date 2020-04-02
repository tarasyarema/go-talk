#!/bin/bash

gcc -o exe_c main.c -lpthread
echo "> gcc -o exe_c main.c -lpthread"

go build -o exe_go main.go
echo "> go build -o exe_go main.go"