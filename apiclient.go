package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var mainFlags = []cli.Flag{
	cli.BoolFlag{"form, f", "Request as application/x-www-form-urlencoded"},
}

func main() {
	app := cli.NewApp()
	app.Name = "API Client"
	app.Usage = "simple rest api client"
	app.Flags = mainFlags
	app.Commands = Commands
	app.Run(os.Args)
}
