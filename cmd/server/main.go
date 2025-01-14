package main

import (
	"fmt"
	"log"

	"github.com/CreedsCode/ipfs-go-manager/internal/api"
	"github.com/CreedsCode/ipfs-go-manager/internal/auth"
	"github.com/CreedsCode/ipfs-go-manager/internal/config"
	"github.com/CreedsCode/ipfs-go-manager/internal/ipfs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	storage := auth.NewInMemoryStorage()
	storage.CreateUser(config.Env.AdminApiKey, 1000, true)
	authMiddleWare := auth.NewAuthMiddleware(storage)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, my beloved ones!")
	})

	node, err := ipfs.ConnectLocalNode(config.Env.IPFSHostAddress)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	api.SetupRoutes(app, node, authMiddleWare)
	api.SetupAdminRoutes(app, authMiddleWare)

	log.Fatal(app.Listen(config.Env.HostAddress))
}
