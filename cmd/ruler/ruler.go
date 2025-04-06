package main

import (
	"strings"

	"github.com/prequel-dev/cre/pkg/logs"
	"github.com/prequel-dev/cre/pkg/ruler"

	"github.com/alecthomas/kong"
	"github.com/rs/zerolog/log"
)

var cli struct {
	Id    ruler.IdCmd    `cmd:"" help:"Generate random id."`
	Build ruler.BuildCmd `cmd:"" help:"Build rules package."`
	Level string         `short:"l" help:"Log level." default:"INFO"`
}

func initLogger() {
	logs.InitLogger(logs.WithPretty(), logs.WithLevel(strings.ToUpper(cli.Level)))
}

func main() {

	ctx := kong.Parse(&cli)

	initLogger()

	log.Info().
		Str("creVersion", ruler.Version).
		Str("gitHash", ruler.Githash).
		Msg("Starting")

	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
