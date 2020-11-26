package policies

import (
	"strconv"

	//"github.com/pangxianfei/framework/helpers/debug"
	"github.com/pangxianfei/framework/request/http/auth"
	"tmaic/app/models"
)

type userPolicy struct {
}

func NewUserPolicy() *userPolicy {
	return &userPolicy{}
}

func (up *userPolicy) Before(IUser auth.IUser, routeParamMap map[string]string) *bool {
	return nil
}
func (up *userPolicy) Create(IUser auth.IUser, routeParamMap map[string]string) bool {
	return false
}
func (up *userPolicy) Update(IUser auth.IUser, routeParamMap map[string]string) bool {
	return true
}
func (up *userPolicy) Delete(IUser auth.IUser, routeParamMap map[string]string) bool {
	return true
}
func (up *userPolicy) ForceDelete(IUser auth.IUser, routeParamMap map[string]string) bool {
	return true
}
func (up *userPolicy) View(IUser auth.IUser, routeParamMap map[string]string) bool {
	//获取当前用户
	currentUser := IUser.Value().(*models.User)
	//debug.Dump(currentUser, routeParamMap)

	//如果使用Authorize func，则routeParamap为空 
	if routeParamMap == nil {
		return true
	}

	// 获取参数用户
	userIdStr, ok := routeParamMap["userId"]
	if !ok {
		return false
	}
	userIdUint, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return false
	}

	if *currentUser.ID != uint(userIdUint) {
		return false
	}

	return true
}
func (up *userPolicy) Restore(IUser auth.IUser, routeParamMap map[string]string) bool {
	return true
}
