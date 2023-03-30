package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"git.brilliantmonkey.net/frigate/frigate-clips/config"
	"golang.org/x/sync/errgroup"
)

type App struct {
	cancel      context.CancelFunc
	errGroup    *errgroup.Group
	errGroupCtx context.Context
}

func waitForExit(cancel context.CancelFunc) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit

	log.Println("Terminating application...")
	cancel()
}

func NewApp() *App {
	ctx, cancel := context.WithCancel(context.Background())
	g, gCtx := errgroup.WithContext(ctx)
	return &App{
		cancel:      cancel,
		errGroup:    g,
		errGroupCtx: gCtx,
	}
}

func (app *App) LoadConfig(pathEnv string, out interface{}) {
	if err := config.Load(pathEnv, out); err != nil {
		panic(err)
	}
}

func (app *App) Go(process func() error) {
	app.errGroup.Go(process)
}

func (app *App) Start(stopCallback func() error) (err error) {
	go waitForExit(app.cancel)

	app.errGroup.Go(func() error {
		<-app.errGroupCtx.Done()
		return stopCallback()
	})
	err = app.errGroup.Wait()
	if err != nil {
		log.Printf("exit reason: %s", err.Error())
	}
	return
}
