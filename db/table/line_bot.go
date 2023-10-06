package table

type LineBot struct {
	GuildID          uint64 `db:"guild_id"`
	LineNotifyToken  byte   `db:"line_notify_token"`
	LineBotToken     byte   `db:"line_bot_token"`
	LineBotSecret    byte   `db:"line_bot_secret"`
	LineGroupID      byte   `db:"line_group_id"`
	LineClientID     byte   `db:"line_client_id"`
	LineClientSecret byte   `db:"line_client_secret"`
	DefaultChannelID byte   `db:"default_channel_id"`
	DebugMode        bool   `db:"debug_mode"`
}
