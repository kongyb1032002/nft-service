package cmd

import (
	"log"
	"nft-service/configs"
	"nft-service/servers"
)

func Execute() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	httpServer := servers.NewHTTPServer(cfg)
	log.Printf("NFT Service running at %s", cfg.ServerAddress)
	httpServer.Run()
}
