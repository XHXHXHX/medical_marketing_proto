#!/bin/bash

set -e

PROTOBUF_VERSION="3.14.0"
PROTOBUF_FOLDER=protoc-$PROTOBUF_VERSION
OS=linux

if [[ $(uname) =~ "Darwin" ]]; then
    OS=osx
fi

export PATH=/usr/local/bin:$PATH

if [[ $(protoc --version | cut -f2 -d' ') != "$PROTOBUF_VERSION" ]]; then
    echo "install protoc-buf $PROTOBUF_VERSION"
    wget https://github.com/google/protobuf/releases/download/v$PROTOBUF_VERSION/protoc-$PROTOBUF_VERSION-$OS-x86_64.zip -O protoc.zip
    unzip protoc.zip -d $PROTOBUF_FOLDER
    sudo rm -rf /usr/local/$PROTOBUF_FOLDER
    sudo mv $PROTOBUF_FOLDER /usr/local/
    sudo rm -rf /usr/local/bin/protoc
    sudo ln -s /usr/local/$PROTOBUF_FOLDER/bin/protoc /usr/local/bin/protoc
    rm -rf protoc.zip
fi

echo "protoc version:"
protoc --version
