package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/ekkx/tcm-platform/internal/config"
)

type NotifyLevel int

const (
	NotifyLevelDebug NotifyLevel = iota
	NotifyLevelInfo
	NotifyLevelError
)

func parseNotifyLevel(s string) NotifyLevel {
	switch strings.ToLower(s) {
	case "debug":
		return NotifyLevelDebug
	case "error":
		return NotifyLevelError
	default:
		return NotifyLevelInfo
	}
}

type NotifyField struct {
	Name  string
	Value string
}

type NotifyMessage struct {
	Level  NotifyLevel
	Title  string
	Body   string
	Fields []NotifyField
}

type Notifier interface {
	Notify(msg NotifyMessage)
}

// NoopNotifier は Discord が未設定の場合に使用するなにもしない実装
type NoopNotifier struct{}

func (n *NoopNotifier) Notify(_ NotifyMessage) {}

// DiscordNotifier は Discord Webhook に通知を送る実装
type DiscordNotifier struct {
	webhookURL string
	minLevel   NotifyLevel
}

func NewDiscordNotifier(cfg config.DiscordConfig) Notifier {
	if cfg.WebhookURL == "" {
		return &NoopNotifier{}
	}
	return &DiscordNotifier{
		webhookURL: cfg.WebhookURL,
		minLevel:   parseNotifyLevel(cfg.NotifyLevel),
	}
}

func (d *DiscordNotifier) Notify(msg NotifyMessage) {
	if msg.Level < d.minLevel {
		return
	}
	if err := d.send(msg); err != nil {
		slog.Error("discord notification failed", slog.String("error", err.Error()))
	}
}

type discordEmbed struct {
	Title       string         `json:"title"`
	Description string         `json:"description,omitempty"`
	Color       int            `json:"color"`
	Fields      []discordField `json:"fields,omitempty"`
	Timestamp   string         `json:"timestamp"`
}

type discordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type discordPayload struct {
	Embeds []discordEmbed `json:"embeds"`
}

func (d *DiscordNotifier) send(msg NotifyMessage) error {
	color := map[NotifyLevel]int{
		NotifyLevelDebug: 9807270,  // gray
		NotifyLevelInfo:  3447003,  // blue
		NotifyLevelError: 15158332, // red
	}[msg.Level]

	fields := make([]discordField, 0, len(msg.Fields))
	for _, f := range msg.Fields {
		fields = append(fields, discordField{Name: f.Name, Value: f.Value, Inline: false})
	}

	payload := discordPayload{
		Embeds: []discordEmbed{
			{
				Title:       msg.Title,
				Description: msg.Body,
				Color:       color,
				Fields:      fields,
				Timestamp:   time.Now().UTC().Format(time.RFC3339),
			},
		},
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal discord payload: %w", err)
	}

	resp, err := http.Post(d.webhookURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("post discord webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook returned status %d", resp.StatusCode)
	}
	return nil
}
