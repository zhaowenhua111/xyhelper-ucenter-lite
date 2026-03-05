package app

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var TestController = testController{}

type testController struct{}

// 三方登录测试 请求
type TestThirdLoginReq struct {
	g.Meta    `path:"/test" method:"post" tags:"登录管理" summary:"login页面登录请求"`
	UserToken string `json:"userToken" v:"required#userToken不能为空" dc:"登录userToken"`
	CarID     string `json:"carID" dc:"车辆id"`
}

// 三方登录测试 响应
type TestThirdLoginRes struct {
	g.Meta `mime:"application/json"`
}

// 三方登录测试接口
func (a *testController) TestThirdLogin(ctx context.Context, req *TestThirdLoginReq) (res *TestThirdLoginRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	// 登录成功
	r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  "登录成功",
	})
	// 登录无权限
	// r.Response.WriteJsonExit(g.Map{
	// 	"code": 1,
	// 	"msg":  "无权限",
	// })
	// 错误处理
	// r.Response.WriteJsonExit(g.Map{
	// 	"code": -1,
	// 	"msg":  "登录错误",
	// })
	return
}
