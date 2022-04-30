#!/bin/bash

set -e

# this file handles recompilation of files after a protocol buffer
# file gets modified


# VERSION
GRPC_GATEWAY_VERSION="v1.15.2";
GOGO_PROTOBUF_VERSION="v1.3.2";

# PATHS
GO_PKG_PATH="${GOPATH}/pkg/mod";
GRPC_PATH="${GO_PKG_PATH}/github.com/grpc-ecosystem/grpc-gateway";
GOGO_PROTOBUF_PATH="${GO_PKG_PATH}/github.com/gogo/protobuf"


# Links to versions
## GRPC
[ -d "${GRPC_PATH}" ] && rm -Rf ${GRPC_PATH}; ln -s "${GRPC_PATH}@${GRPC_GATEWAY_VERSION}" ${GRPC_PATH}

## GOGO PROTOBUF
[ -d "${GOGO_PROTOBUF_PATH}" ] && rm -Rf ${GOGO_PROTOBUF_PATH}; ln -s "${GOGO_PROTOBUF_PATH}@${GOGO_PROTOBUF_VERSION}" ${GOGO_PROTOBUF_PATH}


# HACK - generate non-gogo proto for gateway due to bugs
# (custom types, plus gateway json encoder doesn't support time.Time')
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:../ \
    ./protobuf/account.proto
mv account/account.pb.go account/api/

protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:../ \
    ./protobuf/account.proto
mv ./account/account.pb.gw.go ./account/api/

sed -i "s/package account/package main/g" account/api/account.pb.go
sed -i "s/package account/package main/g" account/api/account.pb.gw.go

# Main account package
# account.proto -> account
# model
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --gogo_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:../ \
    ./protobuf/account.proto
# gateway
# swagger
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --swagger_out=logtostderr=true:./account/api/ \
    ./protobuf/account.proto
# Encode swagger
cd ./account/api/
./build.sh
gofmt -s -w rice-box.go
cd ../..

# email.proto -> email
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    --gogo_out=plugins=grpc:.. \
    ./protobuf/email.proto
    
# sms.proto -> sms
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    --gogo_out=plugins=grpc:.. \
    ./protobuf/sms.proto

# bot.proto -> bot
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --gogo_out=plugins=grpc:.. \
    ./protobuf/bot.proto

# HACK- generate non-gogo proto for gateway due to bugs
# (custom types, plus gateway json encoder doesn't support time.Time')
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --go_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:../ \
    ./protobuf/company.proto
mv company/company.pb.go company/api/

protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:../ \
    ./protobuf/company.proto
mv ./company/company.pb.gw.go ./company/api/

sed -i "s/package company/package main/g" company/api/company.pb.go
sed -i "s/package company/package main/g" company/api/company.pb.gw.go

# company.proto -> company
# model
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --gogo_out=Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,plugins=grpc:../ \
    ./protobuf/company.proto
# swagger
protoc \
    -I ./protobuf/ \
    -I ${GOPATH}/pkg/mod \
    -I ${GRPC_PATH}/third_party/googleapis \
    --swagger_out=logtostderr=true:./company/api/ \
    ./protobuf/company.proto
# Encode swagger

cd ./company/api/

./build.sh
gofmt -s -w rice-box.go
cd ../..
