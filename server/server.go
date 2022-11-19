package server

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid"
	"github.com/nireo/tmpf/filestore"
)

type Server struct {
	FS *filestore.Filestore
}

func (s *Server) CreateFile(c *fiber.Ctx) error {
	// read multipart file
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	uuid := shortuuid.New()

	// store metadata
	if err = s.FS.Add(uuid, file.Filename); err != nil {
		return err
	}

	return c.SaveFile(file, filepath.Join(s.FS.Dir, uuid+filepath.Ext(file.Filename)))
}

func (s *Server) ServeFile(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	meta, err := s.FS.Get(uuid)
	if err != nil {
		return err
	}

	return c.SendFile(filepath.Join(s.FS.Dir, uuid+filepath.Ext(meta.Filename)), true)
}
