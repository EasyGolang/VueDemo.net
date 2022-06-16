package public

import (
	"VueDemo.net/server/router/api/account"
	"VueDemo.net/server/router/api/bing"
	"github.com/gofiber/fiber/v2"
)

/*
/api/public
*/

func Router(router fiber.Router) {
	r := router.Group("/public", MiddleWare)

	r.Get("/url", bing.GetUrl)
	r.Get("/bz", bing.GetBZ)

	r.Post("/login", account.Login)
}
