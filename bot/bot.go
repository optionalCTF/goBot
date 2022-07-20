package bot

import (
	"fmt"
	"gobot/config"

	"github.com/bwmarrin/discordgo"
)

var BotId string
var discBot *discordgo.Session

func Start() {
	discBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	discBot.AddHandler(messageHandler)
	discBot.Identify.Intents = discordgo.IntentsGuildMessages

	err = discBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running !")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content[0:1] != config.BotPrefix {
		return
	} else if m.Content == "!ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
