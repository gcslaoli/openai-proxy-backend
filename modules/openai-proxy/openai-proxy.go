package openai_proxy

import (
	_ "openai-proxy-backend/modules/openai-proxy/controller"
	_ "openai-proxy-backend/modules/openai-proxy/middleware"
	"openai-proxy-backend/modules/openai-proxy/service"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {

	s := g.Server()

	s.BindHandler("/v1/chat/completions", service.ChatCompletionsProxy)
}
