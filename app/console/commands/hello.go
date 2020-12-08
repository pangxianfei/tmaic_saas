package commands

import (
	"fmt"

	"github.com/pangxianfei/framework/cmd"
)

func init() {
	cmd.Add(&HelloWorld{})
}

type HelloWorld struct {
}

func (hw *HelloWorld) Command() string {
	return "Demo {hi}"
}

func (hw *HelloWorld) Description() string {
	return "Demo"
}

func (hw *HelloWorld) Handler(arg *cmd.Arg) error {
	hi, err := arg.Get("hi")
	if err != nil {
		return err
	}

	if hi == nil {
		fmt.Println("Hello World")
		return nil
	}

	fmt.Println(*hi)
	return nil
}
