// messageCreate.go
package botHandler

import (
	"fmt"

    "github.com/maguro-alternative/discord_go_bot/db"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // メッセージが作成されたときに実行する処理
	//u := m.Author

    fmt.Println(m.Content)
    err := db.PingDB()
    if err != nil {
		fmt.Println(err)
	}
    table, err := db.TablesCheck()
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(table)

    if(m.Author.Bot == false){
        s.ChannelMessageSend(m.ChannelID, m.Content)
    }
}