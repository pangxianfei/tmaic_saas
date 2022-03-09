package commands

import (
	"fmt"
	"gitee.com/pangxianfei/frame/cmd"
)

func init() {
	cmd.Add(&Controller{})
}

type Controller struct {
}

func (mi *Controller) Command() string {
	return "controller:make"
}

func (mi *Controller) Description() string {
	return "ControllerInit the list"
}

func (mi *Controller) Handler(arg *cmd.Arg) error {
	hi, err := arg.Get("hi")
	if err != nil {
		return err
	}

	if hi == nil {
		fmt.Println("ControllerInit the list")
		return nil
	}

	fmt.Println(*hi)
	return nil
}
