package cmd

import (
	"awesomeProject1/src/infrastructure"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"time"
)

func StartHttp(ctx context.Context, containerDI *infrastructure.ContainerDI) {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				if err := app.Shutdown(); err != nil {
					panic(err)
				}
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "*",
	}))

	app.Get("/check", containerDI.CheckHanler.Check)
	app.Get("/user/list", containerDI.HandlerUser.List)
	app.Post("/user/create", containerDI.HandlerUser.Create)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
