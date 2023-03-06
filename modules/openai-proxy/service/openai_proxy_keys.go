package service

import (
	"openai-proxy-backend/modules/openai-proxy/config"
	"openai-proxy-backend/modules/openai-proxy/model"
	"time"

	"github.com/cool-team-official/cool-admin-go/cool"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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

// GetKeyById 根据id获取key
func (s *OpenaiProxyKeysService) GetKeyById(id uint) (key string, err error) {
	m := cool.DBM(s.Model)
	record, err := m.Where("id", id).One()
	if err != nil {
		return
	}
	if record == nil {
		return "", gerror.New("没有可用的key")
	}
	return record["key"].String(), nil
}

// CheckKey 检查Key
func (s *OpenaiProxyKeysService) CheckKey(ctx g.Ctx, key string) (err error) {
	type successRes struct {
		Object         string   `json:"object"`
		TotalGranted   float64  `json:"total_granted"`
		TotalUsed      float64  `json:"total_used"`
		TotalAvailable float64  `json:"total_available"`
		Grants         struct { // 信用额度
			Object string `json:"object"`
			Data   []struct {
				Object      string  `json:"object"`
				Id          string  `json:"id"`
				GrantAmount float64 `json:"grant_amount"`
				UsedAmount  float64 `json:"used_amount"`
				EffectiveAt float64 `json:"effective_at"`
				ExpiresAt   float64 `json:"expires_at"`
			}
		} `json:"grants"`
	}

	host := config.Config.OpenAIHost
	g.Dump(host)
	url := host + "/dashboard/billing/credit_grants"
	c := g.Client()
	c.SetHeader("Authorization", "Bearer "+key)
	c.SetHeader("Content-Type", "application/json")
	res, err := c.Get(ctx, url)
	if err != nil {
		return
	}
	res.RawDump()
	if res.StatusCode != 200 {
		return gerror.Newf("请求失败,状态码:%d", res.StatusCode)
	}
	SuccessRes := &successRes{}
	gconv.Struct(res.ReadAllString(), SuccessRes)
	g.Dump(SuccessRes)
	m := cool.DBM(s.Model)

	_, err = m.Data(g.Map{
		"total_granted":   SuccessRes.TotalGranted,
		"total_used":      SuccessRes.TotalUsed,
		"total_available": SuccessRes.TotalAvailable,
		"expire_time":     time.Unix(int64(SuccessRes.Grants.Data[0].ExpiresAt), 0),
	}).Where("key", key).Update()
	// 如果当前时间大于过期时间,则更新key状态为0
	if time.Now().Unix() > int64(SuccessRes.Grants.Data[0].ExpiresAt) {
		_, err = m.Data(g.Map{
			"status": 0,
		}).Where("key", key).Update()
	}

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

// ModifyAfter 修改后
func (s *OpenaiProxyKeysService) ModifyAfter(ctx g.Ctx, method string, param map[string]interface{}) (err error) {
	if method == "Add" {
		// 添加后
		g.Log().Infof(ctx, "添加后")
		s.CheckKey(ctx, param["key"].(string))
	}
	return
}
