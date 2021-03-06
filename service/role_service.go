package service

import (
	"go_web_demo/models"
	"go_web_demo/models/request"
	"go_web_demo/models/response"
	"go_web_demo/pkg/casbin"
	"go_web_demo/utils/gerror"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
)

// 角色列表
func RoleList(page, pageSize int) (response.RolePage, error) {
	role := models.Role{}
	return role.GetAll(page, pageSize)
}

// 角色详情
func RoleDetail(id int) (response.RoleList, error) {
	role := models.Role{Id: id}
	res, err := role.Detail()
	if err != nil {
		return res, err
	}
	auth := strings.Split(res.Auths, ",")
	for _, v := range auth {
		id, _ := strconv.Atoi(v)
		res.Auth = append(res.Auth, id)
	}
	baseAuth := models.GetAllBaseAuth("is_menu = 0 and id in (?)", res.Auth)
	for _, v := range baseAuth {
		res.BaseAuth = append(res.BaseAuth, v.Id)
	}
	return res, err
}

// 添加角色
func RoleAdd(role request.RoleAdd) error {
	defer casbin.ClearEnforcer()
	r := models.Role{
		Pid:  role.Pid,
		Name: role.Name,
	}
	if r.Pid > 0 {
		_, err := r.Detail(r.Pid)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return gerror.NewDbValidateError("所选上级不存在")
			}
			return err
		}
	}

	// 判断是否有重复的角色名
	hasRole, err := models.GetRoleByWhere("name = ?", r.Name)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if hasRole.Id > 0 {
		return gerror.NewDbValidateError("角色名已存在")
	}

	// 严谨起见，从数据库获取权限
	all := models.GetAllBaseAuth("id in (?)", role.Auth)
	allIds := make([]string, 0)
	for _, v := range all {
		allIds = append(allIds, strconv.Itoa(v.Id))
	}
	r.Auth = strings.Join(allIds, ",")

	// 创建角色
	return r.Create(all)
}

// 编辑角色
func RoleEdit(role request.RoleEdit) error {
	defer casbin.ClearEnforcer()
	r := models.Role{
		Id:   role.Id,
		Pid:  role.Pid,
		Name: role.Name,
	}
	if r.Id == r.Pid {
		return gerror.NewNormalValidateError("所属上级和自己一致，数据异常")
	}
	if role.Pid > 0 {
		_, err := r.Detail(r.Pid)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return gerror.NewDbValidateError("所选上级不存在")
			}
			return err
		}
	}

	// 判断是否有重复的角色名
	hasRole, err := models.GetRoleByWhere("name = ? and id <> ?", r.Name, r.Id)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if hasRole.Id > 0 {
		return gerror.NewDbValidateError("角色名已存在")
	}

	// 严谨起见，从数据库获取权限
	all := models.GetAllBaseAuth("id in (?)", role.Auth)
	allIds := make([]string, 0)
	for _, v := range all {
		allIds = append(allIds, strconv.Itoa(v.Id))
	}
	r.Auth = strings.Join(allIds, ",")

	// 编辑角色
	return r.Edit(all)
}

// 删除角色
func RoleDel(id int) error {
	defer casbin.ClearEnforcer()
	role := models.Role{Id: id}

	// 查看用户是否使用该角色
	key := "role:" + strconv.Itoa(id)
	res, err := models.GetCasbinByWhere("p_type = 'g' and (v0 = ? or v1 = ?)", key, key)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}
	if res.PType != "" {
		return gerror.NewDbValidateError("该角色已被使用，无法删除")
	}

	// 同步更新 casbin
	return role.Del()
}

// 获取角色树
func RoleTree(self, pid int) (res []response.Roles, err error) {
	if self > 0 {
		res, err = models.GetRoleTreeByWhere("id <> ? and pid = ?", self, pid)
	} else {
		res, err = models.GetRoleTreeByWhere("pid = ?", pid)
	}
	if err != nil {
		return
	}
	for i, v := range res {
		res[i].Children, err = RoleTree(self, v.Id)
	}
	return
}
