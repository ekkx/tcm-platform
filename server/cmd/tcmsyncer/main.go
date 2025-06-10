package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ekkx/tcmrsv-web/server/internal/config"
	"github.com/ekkx/tcmrsv-web/server/internal/jobs/tcmsyncer"
)

func main() {
	var runOnce bool
	flag.BoolVar(&runOnce, "once", false, "Run the sync job once and exit")
	flag.Parse()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed initializing config: %v", err)
	}

	db, err := cfg.Database.Open()
	if err != nil {
		log.Fatalf("failed connecting to database: %v", err)
	}
	defer db.Close()

	job := tcmsyncer.NewSyncReservationsJob(db, []byte(cfg.PasswordAESKey))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	if runOnce {
		if err := job.Execute(ctx); err != nil {
			log.Fatalf("Sync failed: %v", err)
		}
		log.Println("Sync completed successfully")
		return
	}

	go func() {
		<-sigCh
		log.Println("Shutting down...")
		cancel()
	}()

	log.Println("Starting reservation sync scheduler...")
	job.RunAt12PM(ctx)
}
