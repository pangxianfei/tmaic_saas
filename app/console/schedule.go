package console

import "gitee.com/pangxianfei/frame/cmd"

func Schedule(schedule *cmd.Schedule) {
	schedule.Command("say:hello-world hi,tmaic").EverySecond()
}
