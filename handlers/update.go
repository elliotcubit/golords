package handlers

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

// TODO clean this mess up
func OnMessageUpdate(s *discordgo.Session, mup *discordgo.MessageUpdate) {
	// Ignore ourself, and do nothing if previous message wasn't cached
	// MessageUpdate.Author is nil when requests made via webhook.
	if mup.Author == nil || mup.Author.ID == s.State.User.ID || mup.BeforeUpdate == nil {
		return
	}

	content := strings.ToLower(mup.Content)

	buyWords := []string{
		"buy",
		"bought",
		"purchase",
		"get a new",
		"house",
		"pay",
		"spend",
		"buying",
		"irl money",
		"for me",
		"for us",
		"for eric",
		"gil",
		"csgo crate",
		"get me",
		"give me",
		"cash",
	}

	shouldTrigger := false

	for _, word := range buyWords {
		shouldTrigger = shouldTrigger || strings.Contains(content, word)
	}

	shouldTrigger = shouldTrigger && (strings.Contains(content, "ian") || strings.Contains(content, "ina"))

	if shouldTrigger && !(strings.Contains(content, "don't") || strings.Contains(content, "dont")) {
		s.ChannelMessageSend(mup.ChannelID, "<@208773246009475072>, don't buy that thing!!!!!")
		return
	}
}
