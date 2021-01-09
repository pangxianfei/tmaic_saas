package policies

import (
	"strconv"

	"gitee.com/pangxianfei/frame/request/http/auth"

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
	// get current user
	currentUser := IUser.Value().(*models.User)
	// debug.Dump(currentUser, routeParamMap)
	// debug.Dump("pangxianfei")

	// if use Authorize func, routeParamMap is nil
	if routeParamMap == nil {
		return true
	}

	// get param user
	userIdStr, ok := routeParamMap["userId"]
	if !ok {
		return false
	}
	userIdUint, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return false
	}

	if currentUser.ID != int64(userIdUint) {
		return false
	}

	return true
}
func (up *userPolicy) Restore(IUser auth.IUser, routeParamMap map[string]string) bool {
	return true
}
