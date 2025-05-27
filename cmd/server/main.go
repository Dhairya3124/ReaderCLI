package main

import (
	"github.com/Dhairya3124/ReaderCLI/internal/api"
	"github.com/Dhairya3124/ReaderCLI/internal/store"
	"go.uber.org/zap"
)

func main() {

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()
	store := new(store.Store)
	if _, err := store.Init(); err != nil {
		logger.Fatalf("unable to init store: %v", err)

	}
	config := api.Config{Addr: ":3000"}
	app := api.NewServer(*store, config, logger)
	mux := app.Mount()
	logger.Fatal(app.Run(mux))

}
