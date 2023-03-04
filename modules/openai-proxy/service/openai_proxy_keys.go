package service

import (
	"openai-proxy-backend/modules/openai-proxy/config"
	"openai-proxy-backend/modules/openai-proxy/model"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type OpenaiProxyKeysService struct {
	*cool.Service
}

func NewOpenaiProxyKeysService() *OpenaiProxyKeysService {
	return &OpenaiProxyKeysService{
		&cool.Service{
			Model: model.NewOpenaiProxyKeys(),
		},
	}
}

// CheckKey 检查Key
func (s *OpenaiProxyKeysService) CheckKey(key string) (err error) {
	host := config.Config.OpenAIHost
	g.Dump(host)
	return
}

// RondomGetKey 从数据库中随机获取一个status=1的key
func (s *OpenaiProxyKeysService) RondomGetKey() (string, error) {
	m := cool.DBM(s.Model)
	record, err := m.Where("status", 1).OrderRandom().One()
	if err != nil {
		return "", err
	}
	if record == nil {
		return "", gerror.New("没有可用的key")
	}
	return record["key"].String(), nil
}

// GetKeyInfo 获取key信息
func (s *OpenaiProxyKeysService) GetKeyInfo(key string) (info g.MapStrAny, err error) {
	return
}
