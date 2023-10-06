package table

type Permission struct {
	GuildID                 uint64   `db:"guild_id"`
	LinePermission          uint64   `db:"line_permission"`
	LineUserIdPermission    []uint64 `db:"line_user_id_permission"`
	LineRoleIdPermission    []uint64 `db:"line_role_id_permission"`
	LineBotPermission       uint64   `db:"line_bot_permission"`
	LineBotUserIdPermission []uint64 `db:"line_bot_user_id_permission"`
	LineBotRoleIdPermission []uint64 `db:"line_bot_role_id_permission"`
	VcPermission            uint64   `db:"vc_permission"`
	VcUserIdPermission      []uint64 `db:"vc_user_id_permission"`
	VcRoleIdPermission      []uint64 `db:"vc_role_id_permission"`
	WebhookPermission       uint64   `db:"webhook_permission"`
	WebhookUserIdPermission []uint64 `db:"webhook_user_id_permission"`
	WebhookRoleIdPermission []uint64 `db:"webhook_role_id_permission"`
}
