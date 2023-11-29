package commands

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/bwmarrin/discordgo"
	botRouter "github.com/maguro-alternative/discord_go_bot/bot_handler/bot_router"
)

func PingCommand(db *sqlx.DB) *botRouter.Command {
	/*
		pingコマンドの定義

		コマンド名: ping
		説明: Pong!
		オプション: なし
	*/
	exec := NewSqlDB(db)
	return &botRouter.Command{
		Name:        "ping",
		Description: "Pong!",
		Options: []*discordgo.ApplicationCommandOption{},
		Executor: exec.handlePing,
	}
}

func (h *commandHandlerDB) handlePing(s *discordgo.Session, i *discordgo.InteractionCreate) {
	/*
		pingコマンドの実行

		コマンドの実行結果を返す
	*/
	if i.Interaction.ApplicationCommandData().Name != "ping"{
		return
	}
	if i.Interaction.GuildID != i.GuildID {
		return
	}
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong",
		},
	})
	if err != nil {
		fmt.Printf("error responding to ping command: %v\n", err)
	}
}
