package main

import (
	"fmt"
	"log"

	"github.com/gabyross/superMercado/cmd/server"
)

func main() {
	serverConfig := getServerConfig()

	app := server.NewServerChi(serverConfig)

	log.Printf("Starting server on http://localhost%s", serverConfig.ServerAddress)

	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}

}

func getServerConfig() *server.ConfigServerChi {
	serverConfig := &server.ConfigServerChi{
		ServerAddress:  ":8080",
		LoaderFilePath: "./../products.json",
	}
	return serverConfig
}
