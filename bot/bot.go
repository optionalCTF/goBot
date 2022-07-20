package bot

import (
	"fmt"
	"gobot/config"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	hiscores "github.com/joey-colon/osrs-hiscores"
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

	// Fields is awful if you require arguments that may have spaces e.g usernames
	cmdlet := strings.Fields(m.Content)
	if m.Content[0:1] != config.BotPrefix {
		return
	} else if cmdlet[0] == "!ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else if cmdlet[0] == "!skill" {
		h := hiscores.NewHiscores()

		stats, err := h.GetPlayerSkillRank(cmdlet[1], cmdlet[2])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		lvl, _ := h.GetPlayerSkillLevel(cmdlet[1], cmdlet[2])

		level := strconv.Itoa(int(lvl))
		xp := strconv.Itoa(int(stats))

		response := "Username: " + cmdlet[1] + "\nCurrent " + cmdlet[2] + " level: " + level + "\nCurrent " + cmdlet[2] + " XP: " + xp
		_, err = s.ChannelMessageSend(m.ChannelID, response)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
