package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func responsText(s *discordgo.Session, i *discordgo.InteractionCreate, contentText string) error {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: contentText,
		},
	})
	if err != nil {
		fmt.Printf("error responding to record command: %v\n", err)
		return err
	}
	return nil
}
