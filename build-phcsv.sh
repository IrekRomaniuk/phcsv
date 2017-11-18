#!/usr/bin/bash
env GOOS=linux GOARCH=386 go build -o bin/phcsv_lin
env GOOS=windows GOARCH=386 go build -o bin/phcsv_win