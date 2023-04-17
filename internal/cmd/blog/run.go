package blog

import (
	"context"
	"magnum/ent"
	"magnum/internal/app"
	"magnum/internal/initialize"
	"magnum/internal/router"
	"os"
	"os/signal"
	"time"
)

func Run() {
	app.DB = initialize.DB()
	defer func(DB *ent.Client) {
		_ = DB.Close()
	}(app.DB)
	app.Router = initialize.Router()
	router.Setup(app.Router)
	go func() {
		if err := app.Router.Start(":9009"); err != nil {
			app.Router.Logger.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	app.Router.Logger.Info("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Router.Shutdown(ctx); err != nil {
		app.Router.Logger.Fatal("Server Shutdown:", err)
	}
	app.Router.Logger.Info("Server exiting")
}
