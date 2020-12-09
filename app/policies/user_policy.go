package policies

import (
	//"github.com/pangxianfei/framework/helpers/debug"
	"github.com/pangxianfei/framework/request/http/auth"
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

	return true
}
func (up *userPolicy) Restore(IUser auth.IUser, routeParamMap map[string]string) bool {
	return true
}
