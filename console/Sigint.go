package console

import (
	"os"
	"os/signal"
	"syscall"
)

/*
ListenForSIGINT listens for SIGINT and SIGTERM and calls a provided
handler function when captured. This is useful for closing down
console applications and web servers
*/
func ListenForSIGINT(handler func()) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-done
		handler()
		return
	}()
}
