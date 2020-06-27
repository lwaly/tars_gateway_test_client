#! /bin/bash

cd $GOPATH/src/github.com/lwaly/tars_gateway_test_client/user_proto
protoc --go_out=plugins=tarsrpc:. *.proto