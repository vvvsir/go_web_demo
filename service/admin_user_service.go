package service

import (
	"go_web_demo/models"
	"go_web_demo/models/request"
	"go_web_demo/models/response"
	"go_web_demo/pkg/casbin"
	"go_web_demo/pkg/jwt"
	"go_web_demo/pkg/logger"
	"go_web_demo/utils/gerror"

	"github.com/mojocn/base64Captcha"
	"golang.org/x/crypto/bcrypt"
)

// 登录逻辑
func Login(user *request.LoginUser, id string) (string, error) {
	if !base64Captcha.DefaultMemStore.Verify(id, user.VerifyCode, true) {
		return "", gerror.NewNormalValidateError("验证码错误")
	}
	adminUser := models.GetUserByWhere("user_name = ?", user.Username)
	if adminUser.Id == 0 {
		return "", gerror.NewNormalValidateError("用户不存在")
	}
	if bcrypt.CompareHashAndPassword([]byte(adminUser.Password), []byte(user.Password)) != nil {
		return "", gerror.NewNormalValidateError("用户密码错误")
	}
	if adminUser.Status != 1 {
		return "", gerror.NewNormalValidateError("用户状态错误")
	}
	return jwt.MakeToken(adminUser)
}

// 用户列表
func UserList(name string, page, pageSize int) (res response.AdminUserPage, err error) {
	adminUser := models.AdminUser{}
	if name == "" {
		res, err = adminUser.GetAll(page, pageSize)
	} else {
		res, err = adminUser.GetAll(page, pageSize, "user_name like ? or tel like ?", "%"+name+"%", "%"+name+"%")
	}
	if err != nil {
		return
	}
	roles := models.GetAllRole()
	mappings := models.GetUserRoleMapping()
	for i, v := range res.Data {
		_, ok := mappings[v.Id]
		if !ok {
			continue
		}
		res.Data[i].Roles = []response.CasRole{}
		for _, role := range mappings[v.Id] {
			if _, ok := roles[role]; !ok {
				logger.Logger.Error("角色获取错误: user_id = ", v.Id)
				continue
			}
			res.Data[i].Roles = append(res.Data[i].Roles, roles[role])
		}
	}
	return
}

// 用户详情
func UserDetail(id int) (res response.AdminUserList, err error) {
	adminUser := models.AdminUser{Id: id}
	res, err = adminUser.Detail()
	if err != nil {
		return
	}
	res.Roles = models.GetUserRole(res.Id)
	return
}

// 创建用户
func UserAdd(user request.UserAdd) error {
	defer casbin.ClearEnforcer()
	u := models.AdminUser{
		UserName: user.UserName,
		Tel:      user.Tel,
		Password: user.Password,
		RealName: user.RealName,
		Status:   user.Status,
	}

	// 1.判断用户名和手机号是否存在
	if hasName := models.GetUserByWhere("user_name = ?", user.UserName); hasName.Id > 0 {
		return gerror.NewDbValidateError("用户名已存在")
	}
	if hasTel := models.GetUserByWhere("tel = ?", user.Tel); hasTel.Id > 0 {
		return gerror.NewDbValidateError("手机号已存在")
	}

	// 2.判断角色是否存在
	hasRoles, err := models.GetRolesByWhere("id in (?)", user.Roles)
	if err != nil {
		return err
	}
	if len(hasRoles) != len(user.Roles) {
		return gerror.NewDbValidateError("选择的角色不存在")
	}

	// 3.密码脱敏处理
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(bcryptPassword)
	return u.Create(user.Roles)
}

// 编辑用户
func UserEdit(user request.UserEdit) error {
	defer casbin.ClearEnforcer()
	u := models.AdminUser{
		Id:       user.Id,
		UserName: user.UserName,
		Tel:      user.Tel,
		Password: user.Password,
		RealName: user.RealName,
		Status:   user.Status,
	}

	// 1.判断用户名和手机号是否存在
	if hasName := models.GetUserByWhere("user_name = ? and id <> ?", user.UserName, user.Id); hasName.Id > 0 {
		return gerror.NewDbValidateError("用户名已存在")
	}
	if hasTel := models.GetUserByWhere("tel = ? and id <> ?", user.UserName, user.Id); hasTel.Id > 0 {
		return gerror.NewDbValidateError("手机号已存在")
	}

	// 2.判断角色是否存在
	hasRoles, err := models.GetRolesByWhere("id in (?)", user.Roles)
	if err != nil {
		return err
	}
	if len(hasRoles) != len(user.Roles) {
		return gerror.NewDbValidateError("选择的角色不存在")
	}

	// 3.密码脱敏处理
	if u.Password != "" {
		bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(bcryptPassword)
	}

	return u.Edit(user.Roles)
}

// 删除用户
func UserDel(id int) error {
	defer casbin.ClearEnforcer()
	u := models.AdminUser{
		Id: id,
	}
	return u.Del()
}
