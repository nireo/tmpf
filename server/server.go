package server

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid"
	"github.com/nireo/tmpf/filestore"
)

type Server struct {
	fs *filestore.Filestore
}

func (s *Server) CreateFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	uuid := shortuuid.New()

	return c.SaveFile(file, filepath.Join(s.fs.Dir, uuid+filepath.Ext(file.Filename)))
}

func (s *Server) ServeFile(c *fiber.Ctx) error {
	return nil
}
