package table

type Permission struct {
	guildID                 string   `db:"guild_id"`
	linePermission          int      `db:"line_permission"`
	lineUserIdPermission    []uint64 `db:"line_user_id_permission"`
	lineRoleIdPermission    []uint64 `db:"line_role_id_permission"`
	lineBotPermission       int      `db:"line_bot_permission"`
	lineBotUserIdPermission []uint64 `db:"line_bot_user_id_permission"`
	lineBotRoleIdPermission []uint64 `db:"line_bot_role_id_permission"`
	vcPermission            int      `db:"vc_permission"`
	vcUserIdPermission      []uint64 `db:"vc_user_id_permission"`
	vcRoleIdPermission      []uint64 `db:"vc_role_id_permission"`
	webhookPermission       int      `db:"webhook_permission"`
	webhookUserIdPermission []uint64 `db:"webhook_user_id_permission"`
	webhookRoleIdPermission []uint64 `db:"webhook_role_id_permission"`
}
