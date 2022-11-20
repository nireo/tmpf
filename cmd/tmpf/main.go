package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/nireo/tmpf/filestore"
	"github.com/nireo/tmpf/server"
)

func main() {
	port := flag.Int("port", 8080, "The port to host the server on.")
	flag.Parse()

	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	fs, err := filestore.New("./")
	if err != nil {
		log.Fatal(err)
	}
	server := &server.Server{FS: fs}

	app.Get("/", server.RenderIndex)
	app.Get("/:uuid", server.ServeFile)
	app.Post("/", server.CreateFile)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", *port)))
}
