package main

import (
	"amonhen/cmd"
	"context"
	"github.com/charmbracelet/log"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"os/signal"
)

func main() {
	log.Info("🚀 Launching Amon Hen...")
	cmd.Initialize()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	log.Info("🔥 Starting Amon Hen services...")
	go cmd.Execute()

	<-ctx.Done()

	log.Warn("⬇️ Shutting down Amon Hen...")
}
