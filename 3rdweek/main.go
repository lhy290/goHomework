package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

var (
	ErrInterrupt      = errors.New("iterrupt by user")
	ErrServerInternal = errors.New("internal error")
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	server := http.Server{
		Addr:    ":5612",
		Handler: mux,
	}

	eg, egctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		return server.ListenAndServe()
	})

	eg.Go(func() error {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		select {
		case <-sig:
			ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(5*time.Second))
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				return errors.Wrap(err, "shutdown server")
			}
			return ErrInterrupt
		case <-egctx.Done():
			return ErrServerInternal
		}
	})

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}
