package service

import (
	"go_web_demo/models"
	"go_web_demo/models/response"
)

var (
	// 缓存权限树结构
	authTreeCache []response.Auth

	// 缓存权限map结构
	authMapCache map[int]models.Auth
)

// 返回无极限分类方式的权限
func AuthTreeCache() []response.Auth {
	if len(authTreeCache) == 0 {
		authTreeCache = authTree(0)
	}
	return authTreeCache
}

func authTree(pid int) []response.Auth {
	res := models.GetAllAuth("pid = ?", pid)
	for i, v := range res {
		res[i].Children = authTree(v.Id)
	}
	return res
}

// 缓存权限
func AuthMapCache() map[int]models.Auth {
	if len(authMapCache) == 0 {
		authMapCache = make(map[int]models.Auth)
		base := models.GetAllBaseAuth()
		for _, v := range base {
			authMapCache[v.Id] = v
		}
	}
	return authMapCache
}

// 树形菜单（修改后 必须更新对应缓存authTreeCache、authMapCache）
func TreeMenu() []response.Auth {
	return authTree(0)
}
