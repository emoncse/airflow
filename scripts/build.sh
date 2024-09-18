#!/bin/bash
# Build the Go project and place binary in bin directory
go build -o ../bin/airflow ../cmd/airflow/main.go
