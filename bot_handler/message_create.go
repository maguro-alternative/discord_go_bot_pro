// messageCreate.go
package botHandler

import (
	"fmt"

	//"github.com/maguro-alternative/discord_go_bot/db"

	"github.com/bwmarrin/discordgo"
)

func (h *botHandlerDB) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// メッセージが作成されたときに実行する処理
	//u := m.Author

	fmt.Println(m.Content)
	fmt.Printf("(%%v)  %v\n", h.db)
	err := h.db.DBPing()
	if err != nil {
		fmt.Println(err)
	}
	table, err := h.db.CheckTables()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(table)

	if m.Author.Bot == false {
		s.ChannelMessageSend(m.ChannelID, m.Content)
	}
}
