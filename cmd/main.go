package main

import (
	"context"

	"github.com/2Cheetah/campaign-metadata/internal/app"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()
	app := app.NewApp(ctx)
	app.MustRun()
}
