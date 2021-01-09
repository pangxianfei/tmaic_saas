package middleware

import (
	"gitee.com/pangxianfei/frame/request"
)

// Example 测试中间件
func Example() request.HandlerFunc {
	return func(c request.Context) {

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

	}
}
