package routers

import (
	"ihome/controllers"
	"net/http"
	"strings"

	"github.com/astaxie/beego/context"

	_ "ihome/models"

	"github.com/astaxie/beego"
)

func ignoreStaticPath() {

	//透明static

	beego.InsertFilter("/", beego.BeforeRouter, TransparentStatic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStatic)
}

func TransparentStatic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path
	beego.Debug("request url: ", orpath)
	//如果请求uri还有api字段,说明是指令应该取消静态资源路径重定向
	// /api/v1.0/user
	if strings.Index(orpath, "api") >= 0 {
		return
	}
	// 假设请求来的路径  如果请求的是 / 转换为 /static/html/
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}

func init() {
	ignoreStaticPath() //url重定向
	beego.Router("/", &controllers.MainController{})
	//添加营业区查询路由
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetAreas")
	///api/v1.0/session 添加session处理
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetName;delete:UserLogOut")
	//处理注册功能
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:UserReg")
	///api/v1.0/sessions 登陆功能
	beego.Router("/api/v1.0/sessions", &controllers.UserController{}, "post:UserLogin")
	///api/v1.0/user/name 更新用户名
	beego.Router("/api/v1.0/user/name", &controllers.UserController{}, "put:UpdateUserName")
	//上传头像 /api/v1.0/user/avatar
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:UploadUserPic")
	//查询用户信息
	beego.Router("/api/v1.0/user", &controllers.UserController{}, "get:GetUserInfo")
	//实名认证检查 /api/v1.0/user/auth
	beego.Router("/api/v1.0/user/auth", &controllers.UserController{}, "get:GetUserInfo;post:UpdateUserAuth")
	// /api/v1.0/user/auth

}
