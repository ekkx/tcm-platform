package config

type StripeConfig struct {
	SecretKey     string `env:"STRIPE_SECRET_KEY"`
	WebhookSecret string `env:"STRIPE_WEBHOOK_SECRET"`
	PriceLite     string `env:"STRIPE_PRICE_LITE"`
	PriceStandard string `env:"STRIPE_PRICE_STANDARD"`
	PricePro      string `env:"STRIPE_PRICE_PRO"`
	BaseURL       string `env:"STRIPE_BASE_URL" envDefault:"http://localhost:5173"`
}

func (c StripeConfig) SuccessURL() string {
	return c.BaseURL + "/profile?checkout=success"
}

func (c StripeConfig) CancelURL() string {
	return c.BaseURL + "/profile?checkout=cancel"
}

func (c StripeConfig) PortalReturnURL() string {
	return c.BaseURL + "/profile"
}
