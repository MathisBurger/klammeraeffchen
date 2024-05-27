package pkg

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func HandleJsonUnmarshal(ctx *fiber.Ctx, v interface{}) error {
	ctx.Accepts("application/json")
	return json.Unmarshal(ctx.Body(), v)
}
