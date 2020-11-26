package jobs

import "github.com/pangxianfei/framework/job"

func Initialize() {
	//初始化主题和频道
	job.RegisterQueue()
}
