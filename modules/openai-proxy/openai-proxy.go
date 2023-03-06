package openai_proxy

import (
	"errors"
	"net/http"
	"net/http/httputil"
	"net/url"
	"openai-proxy-backend/modules/openai-proxy/config"
	_ "openai-proxy-backend/modules/openai-proxy/controller"
	_ "openai-proxy-backend/modules/openai-proxy/middleware"
	"openai-proxy-backend/modules/openai-proxy/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	ctx := gctx.GetInitCtx()
	OpenaiProxyKeysService := service.NewOpenaiProxyKeysService()
	s := g.Server()
	s.BindHandler("/v1/chat/completions", func(r *ghttp.Request) {
		key, err := OpenaiProxyKeysService.RondomGetKey()
		if err != nil {
			r.Response.WriteJson(g.Map{
				"code":    1001,
				"message": "无可用OpenAI Key",
			})
			return
		}
		u, _ := url.Parse(config.Config.OpenAIHost)
		g.Dump(u)
		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.URL.Scheme = u.Scheme
			req.URL.Host = u.Host
		},
			ModifyResponse: func(resp *http.Response) error {
				if resp.StatusCode != http.StatusOK {
					token := resp.Request.Header.Get("Authorization")
					g.Dump(token)
					g.Log().Errorf(ctx, "Current Key is %s,  OpenAI API error: %s", token, resp.Status)
					return errors.New("OpenAI API error")
				}
				return nil
			},
		}
		// proxy:= &httputil.ReverseProxy{}
		// 获取 request 的 Authorization 中的 token
		token := r.Request.Header.Get("Authorization")
		g.Dump(token)
		g.Dump(config.Config.Token)
		if token != "Bearer "+config.Config.Token {
			r.Response.WriteJson(g.Map{
				"code":    1001,
				"message": "Key错误",
			})
			return
		}
		newreq := r.Request.Clone(r.Context())
		newreq.Header.Set("Authorization", "Bearer "+key)
		newreq.Host = u.Host
		g.Dump(newreq.Header)
		proxy.ServeHTTP(r.Response.Writer, newreq)
	})
}
