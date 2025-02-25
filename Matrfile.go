//go:build matr

package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/euforic/matr/matr"
)

// Build will build the requested binary
func Build(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("build", flag.ExitOnError)
	var platform = fs.String("p", "linux", "platform")
	fs.Parse(args)

	if len(args) < 1 {
		// loop through cmd dirctory to get all build paths
		files, err := ioutil.ReadDir("./cmd")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			if f.IsDir() {
				args = append(args, f.Name())
			}
		}
	}

	for _, b := range args {
		err := matr.Sh("GOOS=%s GOARCH=arm64 CGO_ENABLED=0 GOGC=2000 go build -a -ldflags '-s' -installsuffix cgo -o build/%s ./cmd/%s", *platform, b, b).Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}

	return nil
}

func Proto(ctx context.Context, args []string) error {
	// Use the root module (.) as input and filter to the proto path
	err := matr.Sh(`buf generate . --path proto`).Run()
	if err != nil {
		return err
	}

	return nil
}

// Docker will build the docker images for all services
func Docker(ctx context.Context, args []string) error {
	err := matr.Sh(`docker build %s.dockerfile -t api/%s:latest`, args[0], args[0]).Run()
	return err
}

func Gql(ctx context.Context, args []string) error {
	fmt.Println("GQL Gen")
	if err := matr.Sh(`go run github.com/99designs/gqlgen generate`).Run(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// Run will run the requested program (bkgrpc, bkgql or bkctl) if left empty bkgrpc and bkgql will run
func Run(ctx context.Context, args []string) error {
	if len(args) == 0 {
		go func() {
			fmt.Println(matr.Sh(`go run ./cmd/grpc`).Run())
		}()
		go func() {
			fmt.Println(matr.Sh(`go run ./cmd/http`).Run())
		}()

		// keep things going
		select {}
	}

	if err := matr.Sh(`go run ./cmd/` + strings.Join(args, " ")).Run(); err != nil {
		return err
	}

	return nil
}
