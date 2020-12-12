package jobs

import (
	"github.com/golang/protobuf/proto"
	"github.com/pangxianfei/framework/helpers/debug"
	"github.com/pangxianfei/framework/job"
	pbs "tmaic/app/jobs/protocol_jobs"
)

func init() {
	job.Add(&DemoJob{})
}

type DemoJob struct {
	job.Job
}

//失败重启次数
func (e *DemoJob) Retries() uint32 {
	return 3
}

//列队名称  Topics 名
func (e *DemoJob) Name() string {
	return "demo-job"
}

//proto 类名参数 实例
func (e *DemoJob) ParamProto() proto.Message {
	return &pbs.DemoJob{}
}

//执行
func (e *DemoJob) Handle(paramPtr proto.Message) error {
	obj := paramPtr.(*pbs.DemoJob)
	debug.Dump(obj)
	return nil
}
