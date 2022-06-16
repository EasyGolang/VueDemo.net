package bing

import (
	"VueDemo.net/server/global/dbType"
	"VueDemo.net/server/router/result"
	"github.com/gofiber/fiber/v2"
)

func GetUrl(c *fiber.Ctx) error {
	return c.JSON(result.Succeed.WithData(dbType.BingImg))
}
