package test

import (
	"fmt"
	"gitee.com/pangxianfei/frame/saas"
	"testing"
	"tmaic/app/saasmodels"
	_ "tmaic/bootstrap"
)

// 查询某个租户信息
func TestTenantSelect(t *testing.T) {

	var TenantId int64
	TenantId = 1
	user := &saasmodels.User{}
	db := saas.SetTenantDB(TenantId)
	result := db.First(&user)
	if result.Error == nil {
		t.Errorf("单元测试失败，错误明细:%s\n", result.Error)
	}
	fmt.Printf("当前租户是：%v\n", user)
}
