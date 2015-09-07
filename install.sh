#!/bin/sh

echo "Installing gnatsd"
go get github.com/nats-io/gnatsd
echo "Installing forego"
go get -u github.com/ddollar/forego
