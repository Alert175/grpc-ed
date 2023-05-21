package main

import "grpc-ed/api/grpc/user_v1"

func main() {
	user_v1.StartGrpcUserServer()
}
