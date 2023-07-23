#!/bin/sh
export BACKEND_ADDR=backend

# start grpcui
grpcui -plaintext $BACKEND_ADDR:50051
