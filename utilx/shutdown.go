package utilx

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

type interrupt struct {
	C chan struct{}
}

func HandleKillSig(handler func()) interrupt {
	i := interrupt{
		C: make(chan struct{}),
	}

	sigChannel := make(chan os.Signal, 1)

	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		for signal := range sigChannel {
			log.Info().Msgf("Receive signal %s, Shutting down...", signal)
			handler()
			close(i.C)
		}
	}()
	return i
}
