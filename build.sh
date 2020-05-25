#!/bin/bash
GOOS=linux go build cmd/iot-lambda/main.go && zip function.zip main