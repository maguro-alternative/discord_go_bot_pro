package table

import (
	"time"

	"github.com/google/uuid"
)

type Webhook struct {
	uuid             uuid.UUID `db:"uuid"`
	GuildID          uint64    `db:"guild_id"`
	WebhookID        uint64    `db:"webhook_id"`
	SubScriptionID   string    `db:"subscription_id"`
	SubScriptionType string    `db:"subscription_type"`
	MentionRoleIDs   []uint64  `db:"mention_role_ids"`
	MentionUserIDs   []uint64  `db:"mention_user_ids"`
	NgOrWords        []string  `db:"ng_or_words"`
	NgAndWords       []string  `db:"ng_and_words"`
	SearchOrWords    []string  `db:"search_or_words"`
	SearchAndWords   []string  `db:"search_and_words"`
	MentionOrWords   []string  `db:"mention_or_words"`
	MentionAndWords  []string  `db:"mention_and_words"`
	CreateAt         time.Time `db:"create_at"`
}
