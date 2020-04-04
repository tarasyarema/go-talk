#!/bin/bash

gcc -o exe_c main.c -lpthread
echo "> gcc -o exe_c main.c -pthread"

g++ -std=c++11 -o exe_cpp main.cpp -pthread
echo "> g++ -std=c++11 -o exe_cpp main.cpp -pthread"

go build -o exe_go main.go
echo "> go build -o exe_go main.go"