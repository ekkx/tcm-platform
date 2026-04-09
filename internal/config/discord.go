package config

type DiscordConfig struct {
	WebhookURL  string `env:"DISCORD_WEBHOOK_URL"`
	NotifyLevel string `env:"DISCORD_NOTIFY_LEVEL" envDefault:"info"`
}
