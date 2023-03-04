package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type config struct {
	OpenAIHost string `json:"openai_host"` // OpenAI Host
}

var Config = &config{}

func init() {
	ctx := gctx.GetInitCtx()
	Config.OpenAIHost = "https://api.openai.com/v1"
	host, err := g.Cfg().GetWithEnv(ctx, "OPENAIHOST")
	if err == nil && host.String() != "" {
		Config.OpenAIHost = host.String()
	}
	g.Log().Infof(ctx, "OpenAIHost: %s", Config.OpenAIHost)
}
