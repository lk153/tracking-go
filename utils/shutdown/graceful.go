package shutdown

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

//ServerAbstract ...
type ServerAbstract interface {
	Start() error
	Close() error
}

//NewCtx ...
func NewCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sysnalChan := make(chan os.Signal, 1)
		signal.Notify(sysnalChan, syscall.SIGINT, syscall.SIGTERM)
		<-sysnalChan
		cancel()
	}()

	return ctx
}

//BlockListen ...
func BlockListen(ctx context.Context, r ServerAbstract) error {
	errChan := make(chan error, 1)
	go func() {
		fmt.Println("Start", reflect.TypeOf(r).Elem().Name())
		if e := r.Start(); e != nil {
			errChan <- e
		}
	}()

	defer close(errChan)

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return r.Close()
	}
}
