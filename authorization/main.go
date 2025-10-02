package main

import "github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/http/grpcserver"

func main() {
	grpcserver.NewGrpc(3001).Run()
}
