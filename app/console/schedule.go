package console

import "github.com/pangxianfei/framework/cmd"

func Schedule(schedule *cmd.Schedule) {
	schedule.Command("say:hello-world hi,tmaic").EverySecond()
}
