package services

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"studae-questions/controller/exception"
	"studae-questions/discord"
	"studae-questions/model"
	"studae-questions/repository"
)

type BotService interface {
	SendQuestion(question model.Question) *exception.BusinessException
	findChannelByName(channel model.Channel) (*discordgo.Channel, *exception.BusinessException)
}

type BotServiceImpl struct {
	DiscordApplication discord.Application
}

func (bot BotServiceImpl) SendQuestion(question model.Question) *exception.BusinessException {
	session := bot.DiscordApplication.GetDiscordApplication()

	channel, err := bot.findChannelByName(question.Channel)

	if err != nil {
		return err
	}

	go func() {
		open, _ := question.Image.Open()

		if err != nil {
			fmt.Printf("An error ocurred while trying to open file.\n")
		}

		_, _ = session.ChannelMessageSendComplex(channel.ID, &discordgo.MessageSend{
			Content: question.Text,
			Files: []*discordgo.File{
				{
					Name:        "test image.png",
					Reader:      open,
					ContentType: "image/png",
				},
			},
		})
	}()

	return nil
}

func (bot BotServiceImpl) findChannelByName(channel model.Channel) (*discordgo.Channel, *exception.BusinessException) {
	channelName := channel.Name

	channelId, found := repository.GetAllowedChannelToSendQuestionByName(channelName)

	if !found {
		fmt.Printf("Channel %s not allowed to send questions", channelName)
		return &discordgo.Channel{},
			&exception.BusinessException{
				Message:    fmt.Sprintf("Channel %s is not allowed to send questions.", channelName),
				StatusCode: 401,
			}
	}

	guildChannel, err := bot.DiscordApplication.GetChannel(channelId)

	if err != nil {
		fmt.Printf("Channel %s not found", channelName)
		return &discordgo.Channel{},
			&exception.BusinessException{
				Message:    fmt.Sprintf("Channel %s could not be found.", channelName),
				StatusCode: 404,
			}
	}

	return guildChannel, nil
}
