package handlers

import (
	"github.com/bwmarrin/discordgo"
	"golords/commands"
	"golords/passive"
	"strings"
)

var activeModules = []commands.ActiveModule{}
var passiveModules = []passive.PassiveModule{}

func RegisterActiveModule(handler commands.ActiveModule) {
	activeModules = append(activeModules, handler)
}

func RegisterPassiveModule(handler passive.PassiveModule) {
	passiveModules = append(passiveModules, handler)
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore ourself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Help will not be its own module
	if strings.HasPrefix(m.Content, "!help") {
		embed := &discordgo.MessageEmbed{Color: 0x3498DB}
		embed.Title = "Commands"

		helpMessage := ""
		for _, handler := range activeModules {
			helpMessage += strings.Join(handler.Prefixes(), ", ")
			helpMessage += ": " + handler.Help() + "\n\n"
		}

		embed.Description = helpMessage

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	if strings.HasPrefix(m.Content, "!") {
		m.Content = m.Content[1:]
		for _, handler := range activeModules {
			for _, prefix := range handler.Prefixes() {
				if strings.HasPrefix(m.Content, prefix) {
					handler.Do(s, m)
					break
				}
			}
		}
	}

	for _, handler := range passiveModules {
		handler.Do(s, m)
	}
}
