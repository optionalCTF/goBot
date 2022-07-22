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
	fmt.Println("[+] Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.ID == "208996860583477249" {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, "🍆")
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if m.Content == "" {
		return
	}

	// Fields is awful if you require arguments that may have spaces e.g usernames
	// This will error if content is empty. Which it will be if user enters a sticker in a bot-readable channel.............
	cmdlet := strings.Fields(m.Content)

	if m.Content[0:1] != config.BotPrefix {
		return
	}

	fmt.Printf("[~] %s ran command %s\n", m.Author.String(), m.Content)

	if cmdlet[0] == "!ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// OSRS hiscore lookup... This is an awful implementation but eh... who cares
	} else if cmdlet[0] == "!skill" {
		username := strings.Join(cmdlet[2:], " ")
		h := hiscores.NewHiscores()

		stats, err := h.GetPlayerSkillRank(username, cmdlet[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		lvl, _ := h.GetPlayerSkillLevel(username, cmdlet[1])

		level := strconv.Itoa(int(lvl))
		xp := strconv.Itoa(int(stats))

		response := "Username: " + username + "\nCurrent " + cmdlet[1] + " level: " + level + "\nCurrent " + cmdlet[1] + " XP: " + xp
		_, err = s.ChannelMessageSend(m.ChannelID, response)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
