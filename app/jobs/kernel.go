package jobs

import "gitee.com/pangxianfei/frame/job"

func Initialize() {
	// 初始化主题和频道
	job.RegisterQueue()
}
