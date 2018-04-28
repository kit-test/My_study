#!/bin/bash

if [ ! -d "autogen" ]; then
    mkdir autogen
fi

if [ ! -d "autogen/McDullProto" ]; then
    mkdir autogen/McDullProto
fi

protoc -I../common/McDullProto/ ../common/McDullProto/*.proto --gogoslick_out=./autogen/McDullProto/

