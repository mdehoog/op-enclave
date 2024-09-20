package main

import (
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/mdehoog/op-nitro/enclave"
	"github.com/mdlayher/vsock"
)

func main() {
	log.SetDefault(log.NewLogger(log.LogfmtHandlerWithLevel(os.Stdout, log.LevelInfo)))

	listener, err := vsock.Listen(1234, &vsock.Config{})
	if err != nil {
		log.Crit("Error opening vsock listener", "error", err)
	}

	s := rpc.NewServer()
	serv, err := enclave.NewServer()
	if err != nil {
		log.Crit("Error creating API server", "error", err)
	}
	err = s.RegisterName(enclave.Namespace, serv)
	if err != nil {
		log.Crit("Error registering API", "error", err)
	}

	err = s.ServeListener(listener)
	if err != nil {
		log.Crit("Error starting server", "error", err)
	}
}
