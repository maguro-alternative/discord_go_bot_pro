// messageCreate.go
package botHandler

import (
	"fmt"
	"context"

	//"github.com/maguro-alternative/discord_go_bot/db"
	"github.com/maguro-alternative/discord_go_bot/db/table"

	"github.com/bwmarrin/discordgo"
)

type PGTable struct {
	SchemaName string `db:"schemaname"`
	TableName  string `db:"tablename"`
	TableOwner string `db:"tableowner"`
}

func (h *botHandlerDB) OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// メッセージが作成されたときに実行する処理
	//u := m.Author

	var pgTables []table.PGTable

	ctx := context.Background()

	fmt.Println(m.Content)
	fmt.Printf("(%%v)  %v\n", h.db)
	err := h.db.DBPing(ctx)
	if err != nil {
		fmt.Println(err)
	}
	err = h.db.CheckTables(ctx, pgTables)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pgTables)

	if m.Author.Bot == false {
		s.ChannelMessageSend(m.ChannelID, m.Content)
	}
}
