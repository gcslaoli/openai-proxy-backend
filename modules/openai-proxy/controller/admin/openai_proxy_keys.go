package admin

import (
	"context"
	"openai-proxy-backend/modules/openai-proxy/service"

	"github.com/cool-team-official/cool-admin-go/cool"

	"github.com/gogf/gf/v2/frame/g"
)

type OpenaiProxyKeysController struct {
	*cool.Controller
}

func init() {
	var openai_proxy_keys_controller = &OpenaiProxyKeysController{
		&cool.Controller{
			Perfix:  "/admin/openai_proxy/keys",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewOpenaiProxyKeysService(),
		},
	}
	// 注册路由
	cool.RegisterController(openai_proxy_keys_controller)
}

// CheckKey 检查Key
type OpenaiProxyKeysCheckKeyReq struct {
	g.Meta `path:"/check_key" method:"GET"`
	Id     uint `json:"id" v:"required#id不能为空"`
}
type OpenaiProxyKeysCheckKeyRes struct {
	*cool.BaseRes
}

func (c *OpenaiProxyKeysController) CheckKey(ctx context.Context, req *OpenaiProxyKeysCheckKeyReq) (res *OpenaiProxyKeysCheckKeyRes, err error) {
	key, err := service.NewOpenaiProxyKeysService().GetKeyById(req.Id)
	if err != nil {
		res = &OpenaiProxyKeysCheckKeyRes{
			BaseRes: cool.Fail("获取Key失败,请检查Key是否存在"),
		}
		return
	}
	err = service.NewOpenaiProxyKeysService().CheckKey(ctx, key)
	if err != nil {
		res = &OpenaiProxyKeysCheckKeyRes{
			BaseRes: cool.Fail("检查Key失败,请检查Key是否正确"),
		}
		return
	}
	res = &OpenaiProxyKeysCheckKeyRes{
		BaseRes: cool.Ok("检查Key成功"),
	}
	return
}
