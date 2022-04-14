package main

import (
	"github.com/kailash-bhanushali/backend-golang/internal/config"
	"github.com/kailash-bhanushali/backend-golang/pkg/server"
)

func main() {
	server.NewServer(config.NewServerConfig())
}
