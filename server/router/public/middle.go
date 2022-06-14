package public

import (
	"VueDemo.net/server/router/middle"
	"VueDemo.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func MiddleWare(c *fiber.Ctx) error {
	c.Set("Data-Path", "VueDemo.net/api/public")

	// 授权验证
	err := middle.EncryptAuth(c)
	if err != nil {
		return c.JSON(result.ErrAuth.WithData(err))
	}

	return c.Next()
}
