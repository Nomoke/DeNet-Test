package app

import (
	"context"
	"log"

	"denet/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg

	ctx := context.Background()
	_ = ctx
}
