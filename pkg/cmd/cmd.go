package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Name string
	Close func()
}

func Run(app *App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println("Get a signal", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			app.Close()
			fmt.Println(app.Name , "app exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
