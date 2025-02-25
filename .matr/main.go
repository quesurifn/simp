//go:build matr

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/euforic/matr/matr"
)

func main() {
	// Create new Matr instance
	m := matr.New()
	
	// Build will build the requested binary
	m.Handle(&matr.Task{
		Name: "build",
		Summary: "Build will build the requested binary",
		Doc: `Build will build the requested binary`,
		Handler: Build,
	})
	
	// Proto
	m.Handle(&matr.Task{
		Name: "proto",
		Summary: "",
		Doc: ``,
		Handler: Proto,
	})
	
	// Docker will build the docker images for all services
	m.Handle(&matr.Task{
		Name: "docker",
		Summary: "Docker will build the docker images for all services",
		Doc: `Docker will build the docker images for all services`,
		Handler: Docker,
	})
	
	// Gql
	m.Handle(&matr.Task{
		Name: "gql",
		Summary: "",
		Doc: ``,
		Handler: Gql,
	})
	
	// Run will run the requested program (bkgrpc, bkgql or bkctl) if left empty bkgrpc and bkgql will run
	m.Handle(&matr.Task{
		Name: "run",
		Summary: "Run will run the requested program (bkgrpc, bkgql or bkctl) if left empty bkgrpc and bkgql will run",
		Doc: `Run will run the requested program (bkgrpc, bkgql or bkctl) if left empty bkgrpc and bkgql will run`,
		Handler: Run,
	})

	// Setup context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Setup signal handling for SIGINT and SIGTERM
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Run Matr in a separate goroutine
	errChan := make(chan error)
	go func() {
		errChan <- m.Run(ctx, os.Args[1:]...)
	}()

	// Wait for Matr to finish, a timeout, or a signal
	select {
	case err := <-errChan:
		if err != nil {
			os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
		}
	case <-ctx.Done():
		os.Stderr.WriteString("ERROR: Context timed out\n")
	case <-sig:
		cancel()
		os.Stderr.WriteString("INFO: Received signal, shutting down\n")
	}
}