package config

type SchedulerConfig struct {
	CronExpression    string `env:"SCHEDULER_CRON_EXPRESSION" envDefault:"0 0 12 * * *"`
	MaxConcurrentJobs int    `env:"SCHEDULER_MAX_CONCURRENT_JOBS" envDefault:"5"`
}
