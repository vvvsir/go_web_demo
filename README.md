## go_web_demo

## 使用技术

* GIN 框架
* JWT 认证
* Casbin权限
* GORM 框架
* Swagger Api
* logrus日志

该项目仅仅包含了最简单的[用户-角色-权限]管理系统
如果有需要，请修改conf文件夹下的配置文件

## 使用方法

### api

```bash
$ go run main.go -init=true #-init=true初始sql脚本（第一次需要，之后运行无需-init参数）
```
### client-app

```bash
$ cd client-app
# 安装依赖文件, 可以使用cnpm镜像加速
# npm install -g cnpm --registry=https://registry.npm.taobao.org
# cnpm install
$ npm install
# 本地调试
$ npm run serve
# 正式发布
$ npm run build:prod
```

感谢以下框架的开源支持
* [Gin] - https://gin-gonic.com/
* [GORM] - http://gorm.io/
* [Casbin] - https://casbin.org/
* [vue-element-admin] - https://github.com/PanJiaChen/vue-element-admin/
