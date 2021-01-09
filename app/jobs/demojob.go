package jobs

import (
	"gitee.com/pangxianfei/frame/helpers/debug"
	"gitee.com/pangxianfei/frame/job"
	"github.com/golang/protobuf/proto"

	pbs "tmaic/app/jobs/protocol_jobs"
)

func init() {
	job.Add(&DemoJob{})
}

type DemoJob struct {
	job.Job
}

// Retries 失败重启次数
func (e *DemoJob) Retries() uint32 {
	return 3
}

// Name 列队名称  Topics 名
func (e *DemoJob) Name() string {
	return "demo-job"
}

// ParamProto proto 类名参数 实例
func (e *DemoJob) ParamProto() proto.Message {
	return &pbs.DemoJob{}
}

// Handle 执行
func (e *DemoJob) Handle(paramPtr proto.Message) error {
	obj := paramPtr.(*pbs.DemoJob)
	debug.Dump(obj)
	debug.Dump("DemoJob Handle")
	return nil
}
