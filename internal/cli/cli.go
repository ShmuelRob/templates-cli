package cli

import (
	"github.com/ShmuelRob/templates-cli/internal/templates"
	"github.com/urfave/cli/v2"
)

// NewApp creates a new CLI application
func NewApp() *cli.App {
	app := &cli.App{
		Name:    "pytgen",
		Usage:   "Generate Python project templates",
		Version: "0.1.0",
		Commands: []*cli.Command{
			{
				Name:  "etl",
				Usage: "Generate a Python ETL project template",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "name",
						Aliases: []string{"n"},
						Usage:   "Project name",
						Value:   "python-etl-project",
					},
					&cli.StringFlag{
						Name:    "extract",
						Aliases: []string{"e"},
						Usage:   "Extract method (file, api, database)",
						Value:   "file",
					},
					&cli.StringFlag{
						Name:    "transform",
						Aliases: []string{"t"},
						Usage:   "Transform method (basic, advanced)",
						Value:   "basic",
					},
					&cli.StringFlag{
						Name:    "load",
						Aliases: []string{"l"},
						Usage:   "Load destination (file, database, api)",
						Value:   "file",
					},
					&cli.BoolFlag{
						Name:  "venv",
						Usage: "Initialize virtual environment",
						Value: false,
					},
				},
				Action: templates.GenerateETLTemplate,
			},
			// Add a new interactive command
			{
				Name:    "interactive",
				Aliases: []string{"i"},
				Usage:   "Launch interactive project generator",
				Action:  templates.InteractiveGenerator,
			},
		},
	}

	return app
}
