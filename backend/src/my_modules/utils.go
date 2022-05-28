package my_modules

import (
	"context"
	"time"
)

func SetTimeOut(callback func(), wait time.Duration) context.CancelFunc {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done(): //context cancelled
		case <-time.After(wait): //timeout
			callback()
		}
	}()
	return cancel
}
