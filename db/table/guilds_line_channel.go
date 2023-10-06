package table

type GuildsLineChannel struct {
	ChannelID      uint64   `db:"channel_id"`
	GuildID        uint64   `db:"guild_id"`
	LineNgChannel  bool     `db:"line_ng_channel"`
	NgMessageTypes []string `db:"ng_message_types"`
	MessageBot     bool     `db:"message_bot"`
	NgUsers        []uint64 `db:"ng_users"`
}
