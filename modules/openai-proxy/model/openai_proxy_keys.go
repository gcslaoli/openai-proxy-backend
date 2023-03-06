package model

import (
	"time"

	"github.com/cool-team-official/cool-admin-go/cool"
)

const TableNameOpenaiProxyKeys = "openai_proxy_keys"

// OpenaiProxyKeys mapped from table <openai_proxy_keys>
type OpenaiProxyKeys struct {
	*cool.Model
	Key            string     `gorm:"column:key;not null;comment:Key" json:"key"`
	Status         bool       `gorm:"column:status;not null;comment:状态" json:"status"`
	ExpireTime     *time.Time `gorm:"column:expire_time;comment:过期时间" json:"expire_time"`
	TotalGranted   float64    `gorm:"column:total_granted;comment:总授权次数" json:"total_granted"`
	TotalUsed      float64    `gorm:"column:total_used;comment:总使用次数" json:"total_used"`
	TotalAvailable float64    `gorm:"column:total_available;comment:总可用次数" json:"total_available"`
	Remark         string     `gorm:"column:remark;comment:备注" json:"remark"`
}

// TableName OpenaiProxyKeys's table name
func (*OpenaiProxyKeys) TableName() string {
	return TableNameOpenaiProxyKeys
}

// GroupName OpenaiProxyKeys's table group
func (*OpenaiProxyKeys) GroupName() string {
	return "default"
}

// NewOpenaiProxyKeys create a new OpenaiProxyKeys
func NewOpenaiProxyKeys() *OpenaiProxyKeys {
	return &OpenaiProxyKeys{
		Model: cool.NewModel(),
	}
}

// init 创建表
func init() {
	cool.CreateTable(&OpenaiProxyKeys{})
}
