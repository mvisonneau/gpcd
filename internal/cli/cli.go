package cli

import (
	"log"
	"runtime"
	"time"

	"github.com/mvisonneau/gpcd/internal/cmd"
	"github.com/urfave/cli/v2"
)

// Run handles the instanciation of the CLI application.
func Run(version string, args []string) {
	err := NewApp(version, time.Now()).Run(args)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// NewApp configures the CLI application.
func NewApp(version string, start time.Time) (app *cli.App) {
	app = cli.NewApp()
	app.Name = "gpcd"
	app.Version = version
	app.Usage = "download bulk medias from GoPro Cloud"
	app.EnableBashCompletion = true

	app.Flags = cli.FlagsByName{
		&cli.StringFlag{
			Name:    "api-endpoint",
			EnvVars: []string{"GPCD_API_ENDPOINT"},
			Usage:   "Go Pro Cloud API endpoint",
			Value:   "https://api.gopro.com/media/",
		},
		&cli.StringFlag{
			Name:    "local-path",
			EnvVars: []string{"GPCD_LOCAL_PATH"},
			Usage:   "where the medias should be downloaded",
			Value:   "./medias",
		},
		&cli.StringFlag{
			Name:    "bearer-token",
			EnvVars: []string{"GPCD_BEARER_TOKEN"},
			Usage:   "Used to authenticate over your account",
		},
		&cli.StringFlag{
			Name:    "user-agent",
			EnvVars: []string{"GPCD_USER_AGENT"},
			Value:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0",
		},
		&cli.TimestampFlag{
			Name:   "from",
			Usage:  "filter for medias captured after this date",
			Layout: time.RFC3339,
		},
		&cli.TimestampFlag{
			Name:   "to",
			Usage:  "filter for medias captured before this date",
			Layout: time.RFC3339,
		},
		&cli.IntFlag{
			Name:  "max-concurrent-downloads",
			Value: runtime.NumCPU(),
		},
	}

	app.Commands = cli.CommandsByName{
		{
			Name:   "list",
			Usage:  "list available medias",
			Action: cmd.List,
		},
		{
			Name:   "download",
			Usage:  "download medias",
			Action: cmd.Download,
		},
	}

	return
}
