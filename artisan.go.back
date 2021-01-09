package main

import (
	"os"

	"gitee.com/pangxianfei/frame/graceful"
	"gitee.com/pangxianfei/frame/helpers/log"
	"gitee.com/pangxianfei/frame/sentry"

	"tmaic/bootstrap"

	"github.com/urfave/cli"

	"gitee.com/pangxianfei/frame/console"

	"tmaic/database/migrations"

	"gitee.com/pangxianfei/frame/cmd"
	command_queue "gitee.com/pangxianfei/frame/cmd/commands/queue"
	"gitee.com/pangxianfei/frame/cmd/commands/schedule"

	app_schedule "tmaic/app/console"

	"tmaic/app/console/commands"
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
	app.Version = "0.5.5"

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

	graceful.ShutDown(true)
}
