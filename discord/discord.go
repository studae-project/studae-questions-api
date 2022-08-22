package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"sync"
)

var channelCache = map[string]*discordgo.Channel{}

type ApplicationImpl struct {
	session *discordgo.Session
}

func (d *ApplicationImpl) GetDiscordApplication() *discordgo.Session {
	return d.session
}

func (d *ApplicationImpl) startApplication() {
	instance, err := discordgo.New(fmt.Sprintf("Bot %s", os.Getenv("TOKEN")))

	if err != nil {
		fmt.Printf("an error ocurred while trying to instance Discord BOT")
	}

	err = instance.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	d.session = instance
}

func (d *ApplicationImpl) GetChannel(channelId string) (*discordgo.Channel, error) {
	session := d.session

	if val, ok := channelCache[channelId]; ok {
		return val, nil
	}

	foundChannel, err := session.Channel(channelId)

	if err != nil {
		return nil, err
	}

	channelCache[channelId] = foundChannel
	return foundChannel, nil
}

type Application interface {
	startApplication()
	GetDiscordApplication() *discordgo.Session
	GetChannel(channelId string) (*discordgo.Channel, error)
}

var discord *ApplicationImpl
var onceDiscord sync.Once

func StartDiscordApplication() {
	onceDiscord.Do(func() {
		discord = &ApplicationImpl{}
		discord.startApplication()
	})
}

func GetDiscordSession() *ApplicationImpl {
	return discord
}
