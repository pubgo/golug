package healthy_type

import (
	"github.com/gofiber/fiber/v2"
)

type Handler func(req *fiber.Ctx) error
