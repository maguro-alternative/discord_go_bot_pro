package discord_model

type DiscordUser struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	GlobalName       string `json:"global_name"`
	DisplayName      string `json:"display_name"`
	Avatar           string `json:"avatar"`
	AvatarDecoration string `json:"avatar_decoration"`
	Discriminator    string `json:"discriminator"`
	PublicFlags      int    `json:"public_flags"`
	Flags            int    `json:"flags"`
	Banner           string `json:"banner"`
	BannerColor      int    `json:"banner_color"`
	AccentColor      int    `json:"accent_color"`
	Locale           string `json:"locale"`
	MfaEnabled       bool   `json:"mfa_enabled"`
	PremiumType      int    `json:"premium_type"`
	Email            string `json:"email"`
	Verified         bool   `json:"verified"`
	Bio              string `json:"bio"`
}
