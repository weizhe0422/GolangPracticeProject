#!/bin/bash

protoc gRPC_Greet/greetpb/greet.proto --go_out=plugins=grpc:. 