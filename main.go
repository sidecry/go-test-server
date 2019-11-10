package main

import (
	"net/http"
	"os"

	"github.com/grimoh/test-server/server"
	"github.com/urfave/cli"
)

const (
	name    = "test-server"
	version = "1.0.1"
	usage   = ""
	addr    = ":8080"
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Usage = usage
	app.Action = mainAction
	app.Run(os.Args)
}

func mainAction(c *cli.Context) error {
	s := server.New()
	h := s.Handler()
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	if err := srv.ListenAndServe(); err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	return nil
}
