package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/prequel-dev/cre/pkg/logs"
	"github.com/prequel-dev/cre/pkg/ruler"

	"github.com/alecthomas/kong"
)

var cli struct {
	IdCmd struct {
		Num int ` optional:"" arg:"" default:"1" name:"num" help:"Number to generate."`
	} `cmd:"" name:"id" short:"i" help:"Generate random id."`
	BuildCmd struct {
		InPath  string `name:"path" short:"p" help:"Path to read rules" default:"rules"`
		OutPath string `name:"out" short:"o" help:"Optional path to write files; default curdir"`
		Version string `name:"vers" short:"v" help:"Optional semantic version override"`
	} `cmd:"" name:"build" short:"b" help:"Build rules package."`
	VersionCmd struct {
	} `cmd:"" name:"version" short:"v" help:"Print version information." default:1`
	Version bool   `name:"version" short:"v" help:"Version"`
	Level   string `short:"l" help:"Log level." default:"info"`
}

func initLogger() {
	logs.InitLogger(logs.WithPretty(), logs.WithLevel(strings.ToUpper(cli.Level)))
}

func printVersion() {
	fmt.Fprintf(os.Stdout, "creVersion: %s\n", ruler.Version)
	fmt.Fprintf(os.Stdout, "gitHash: %s\n", ruler.Githash)
}

func main() {

	ctx := kong.Parse(&cli)

	initLogger()

	switch ctx.Command() {
	case "version":
		printVersion()
	case "id":
		ruler.RunId(cli.IdCmd.Num)
	case "build":
		ruler.RunBuild(cli.BuildCmd.InPath, cli.BuildCmd.OutPath, cli.BuildCmd.Version)
	}
}
