package main

import (
	"github.com/pangxianfei/framework/cmd"
	command_queue "github.com/pangxianfei/framework/cmd/commands/queue"
	"github.com/pangxianfei/framework/cmd/commands/schedule"
	"github.com/pangxianfei/framework/console"
	"github.com/pangxianfei/framework/graceful"
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/sentry"
	"github.com/urfave/cli"
	"os"
	app_schedule "tmaic/app/console"
	"tmaic/app/console/commands"
	"tmaic/bootstrap"
	"tmaic/database/migrations"
)

func init() {
	bootstrap.Initialize()
	migrations.Initialize()
	command_queue.Initialize()
	commands.Initialize()
	schedule.Initialize()
	app_schedule.Schedule(cmd.NewSchedule())
}

func main() {
	cliServe()
}

func cliServe() {
	app := cli.NewApp()
	app.Name = "artisan"
	app.Usage = "Let's work like an artisan"
	app.Version = "1.0.1"

	app.Commands = cmd.List()

	app.Action = func(c *cli.Context) error {
		console.Println(console.CODE_INFO, "COMMANDS:")

		for _, cate := range app.Categories() {
			categoryName := cate.Name
			if categoryName == "" {
				categoryName = "kernel"
			}
			console.Println(console.CODE_WARNING, "    "+categoryName+":")

			for _, cmds := range cate.Commands {
				console.Println(console.CODE_SUCCESS, "        "+cmds.Name+" "+console.Sprintf(console.CODE_INFO, "%s", cmds.ArgsUsage)+"    "+console.Sprintf(console.CODE_WARNING, "%s", cmds.Usage))
			}
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		sentry.CaptureError(err)
		log.Fatal(err.Error())
	}

	// tmaic-go framework shutdown
	graceful.ShutDown(true)
}
