package utils

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SignalHandler() {
	sigCh := make(chan os.Signal, 1)
	doneCh := make(chan os.Signal)

	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for {
			select {
			case <-time.After(60 * time.Second):
				continue
			case sig := <-sigCh:
				switch sig {
				case syscall.SIGQUIT:
					doneCh <- sig
					return
				case syscall.SIGINT:
					doneCh <- sig
					return
				case syscall.SIGTERM:
					doneCh <- sig
					return
				default:
					fmt.Println("default")
					continue
				}
			}
		}
	}()
	zap.L().Sugar().Infof("Exit by Signal: %v", <-doneCh)
}
