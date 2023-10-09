package botHandler

import (
	"fmt"

	//"github.com/maguro-alternative/discord_go_bot/db"

	"github.com/bwmarrin/discordgo"
)

func (h *botHandlerDB) OnVoiceStateUpdate(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	fmt.Print("hoge")
	err := h.db.DBPing()
	if err != nil {
		fmt.Println(err)
	}
    fmt.Printf("%+v", vs.VoiceState)
}