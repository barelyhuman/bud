package sig

import (
	"context"
	"os"
	"os/signal"
)

// Trap cancels the context based on a signal
func Trap(ctx context.Context, signals ...os.Signal) (context.Context, context.CancelFunc) {
	ret, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal, len(signals))
	go func() {
		<-ch
		cancel()
		signal.Stop(ch)
	}()
	signal.Notify(ch, signals...)
	return ret, cancel
}
