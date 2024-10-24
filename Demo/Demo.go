package main

import (
	"fmt"

	"github.com/feiya1314/Auth"
	"github.com/feiya1314/Common"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("App:" + Auth.Check("yf"))
	fmt.Println("App:" + Common.Greet("jack"))
	// 定义fiber应用
	app := fiber.New()

	// 定义http route
	app.Get("/", func(c *fiber.Ctx) error {
		// 以字符串的形式返回hello world
		return c.SendString("Hello, World 👋!")
	})

	// 监听3000端口
	app.Listen(":3000")
}
