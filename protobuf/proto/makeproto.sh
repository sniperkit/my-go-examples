#!/bin/bash
# my-go-examples makeproto.sh

protoc --go_out=. messages.proto
cp messages.pb.go ../.
