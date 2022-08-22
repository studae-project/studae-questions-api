package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"studae-questions/controller"
	"studae-questions/discord"
	"studae-questions/services"
)

func main() {
	app := fiber.New()
	discord.StartDiscordApplication()

	var botServiceImpl services.BotService = services.BotServiceImpl{
		DiscordApplication: discord.GetDiscordSession(),
	}

	questionController := controller.QuestionController{
		BotService: &botServiceImpl,
	}

	app.Post("/question", func(ctx *fiber.Ctx) error {
		questionController.POST(ctx)
		return nil
	})

	app.Listen(getAppPort())
}

func getAppPort() string {
	applicationPort, found := os.LookupEnv("PORT")

	if !found {
		return ":8090"
	}

	return fmt.Sprintf(":%s", applicationPort)
}
