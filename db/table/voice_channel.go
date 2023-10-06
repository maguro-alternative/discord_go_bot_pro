package table

type VcChannel struct {
	VcID            uint64   `db:"vc_id"`
	GuildID         uint64   `db:"guild_id"`
	SendSignal      bool     `db:"send_signal"`
	SendChannelID   uint64   `db:"send_channel_id"`
	JoinBot         bool     `db:"join_bot"`
	EveryoneMention bool     `db:"everyone_mention"`
	MentionRoleIDs  []uint64 `db:"mention_role_ids"`
}
