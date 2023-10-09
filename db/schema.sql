CREATE TABLE IF NOT EXISTS guild_set_permissions (
    guild_id NUMERIC PRIMARY KEY,
    line_permission NUMERIC NOT NULL DEFAULT 8,
    line_user_id_permission NUMERIC[] NOT NULL DEFAULT '{}',
    line_role_id_permission NUMERIC[] NOT NULL DEFAULT '{}',
    line_bot_permission NUMERIC NOT NULL DEFAULT 8,
    line_bot_user_id_permission NUMERIC[] NOT NULL DEFAULT '{}',
    line_bot_role_id_permission NUMERIC[] NOT NULL DEFAULT '{}',
    vc_permission NUMERIC NOT NULL DEFAULT 8,
    vc_user_id_permission NUMERIC[] NOT NULL DEFAULT '{}',
    vc_role_id_permission NUMERIC[] NOT NULL DEFAULT '{}',
    webhook_permission NUMERIC NOT NULL DEFAULT 8,
    webhook_user_id_permission NUMERIC[] NOT NULL DEFAULT '{}',
    webhook_role_id_permission NUMERIC[] NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS line_bot (
    guild_id NUMERIC PRIMARY KEY,
    line_notify_token BYTEA NOT NULL DEFAULT '',
    line_bot_token BYTEA NOT NULL DEFAULT '',
    line_bot_secret BYTEA NOT NULL DEFAULT '',
    line_group_id BYTEA NOT NULL DEFAULT '',
    line_client_id BYTEA NOT NULL DEFAULT '',
    line_client_secret BYTEA NOT NULL DEFAULT '',
    default_channel_id NUMERIC NOT NULL DEFAULT 0,
    debug_mode BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS guilds_line_channel (
    channel_id NUMERIC PRIMARY KEY,
    guild_id NUMERIC NOT NULL,
    line_ng_channel BOOLEAN NOT NULL DEFAULT FALSE,
    ng_message_types VARCHAR(50)[] NOT NULL DEFAULT '{}',
    message_bot BOOLEAN NOT NULL DEFAULT True,
    ng_users NUMERIC[] NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS guilds_vc_signal (
    vc_id NUMERIC PRIMARY KEY,
    guild_id NUMERIC NOT NULL,
    send_signal BOOLEAN NOT NULL DEFAULT True,
    send_channel_id NUMERIC NOT NULL DEFAULT 0,
    join_bot BOOLEAN NOT NULL DEFAULT True,
    everyone_mention BOOLEAN NOT NULL DEFAULT True,
    mention_role_ids NUMERIC[] NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS webhook (
    uuid UUID PRIMARY KEY,
    guild_id NUMERIC NOT NULL,
    webhook_id NUMERIC NOT NULL,
    subscription_id VARCHAR(50) NOT NULL,
    subscription_type VARCHAR(50) NOT NULL,
    mention_role_ids NUMERIC[] NOT NULL DEFAULT '{}',
    mention_user_ids NUMERIC[] NOT NULL DEFAULT '{}',
    ng_or_words VARCHAR(50)[] NOT NULL DEFAULT '{}',
    ng_and_words VARCHAR(50)[] NOT NULL DEFAULT '{}',
    search_or_words VARCHAR(50)[] NOT NULL DEFAULT '{}',
    search_and_words VARCHAR(50)[] NOT NULL DEFAULT '{}',
    mention_or_words VARCHAR(50)[] NOT NULL DEFAULT '{}',
    mention_and_words VARCHAR(50)[] NOT NULL DEFAULT '{}',
    create_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);