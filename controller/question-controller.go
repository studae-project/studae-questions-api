package controller

import (
	"github.com/gofiber/fiber/v2"
	"studae-questions/controller/exception"
	requests "studae-questions/controller/request"
	"studae-questions/model"
	"studae-questions/services"
)

type QuestionController struct {
	BotService *services.BotService
}

func (controller *QuestionController) POST(ctx *fiber.Ctx) {
	form, err := ctx.MultipartForm()

	file, err := ctx.FormFile("file")

	if err != nil {
		print(file.Filename)
	}

	if err != nil {
		exceptionMessage := err.Error()
		ctx.Status(500).JSON(&exception.ErrorResponse{
			Message:        "Error parsing multipart form.",
			DetailedReason: &exceptionMessage,
		})
	}

	botService := *controller.BotService
	questionRequest := requests.ParseCreateQuestionRequest(form)

	businessException := botService.SendQuestion(model.Question{
		Text:    questionRequest.Text,
		Channel: model.Channel{Name: questionRequest.Channel},
		Image:   questionRequest.Image.Header,
	})

	if businessException != nil {
		handleError(businessException.Message, businessException.StatusCode, nil, ctx)
	}

	ctx.Status(201)
}

func handleError(message string, statuscode int, detailedReason *string, ctx *fiber.Ctx) {
	ctx.Status(statuscode).JSON(exception.ErrorResponse{
		Message:        message,
		DetailedReason: detailedReason,
	})
}
