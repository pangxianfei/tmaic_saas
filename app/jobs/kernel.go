package jobs

import "gitee.com/pangxianfei/frame/job"

func Initialize() {
	// initialize topic and channel
	job.RegisterQueue()
}
