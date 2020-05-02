package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Name  string
	Close func()
}

func Run(name string, closeFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println("Get a signal", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			fmt.Println(name, "app exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
