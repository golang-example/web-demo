package user

import (
	"web-demo/web"
	"github.com/kataras/iris"
	"web-demo/service"
	"web-demo/model/user"
	"web-demo/exception"
)

func init() {
	web.RegisterHandler("Get", "/webdemo/user/v1/isexist", isexist)
	web.RegisterHandler("Post", "/webdemo/user/v1/update", update)
	web.RegisterHandler("Post", "/webdemo/user/v1/add", add)
}

//查询用户名是否已经存在
func isexist(ctx *iris.Context) {
	//request
	userName := ctx.URLParam("userName")
	if userName == "" {
		ctx.JSON(iris.StatusOK, exception.ParamException())
		return
	}

	//process
	user := service.GetUserInfo(userName)
	if user == nil {
		ctx.JSON(iris.StatusOK, iris.Map{"code": -1})
		return
	}

	//response
	ctx.JSON(iris.StatusOK, iris.Map{"code": 0})
}

//修改用户信息
func update(ctx *iris.Context) {
	//request
	request := new(user.UserRequest)
	if err := ctx.ReadJSON(request); err != nil {
		ctx.JSON(iris.StatusOK, exception.ParamException())
		return
	}

	if request.UserName == "" || request.Pwd == "" {
		ctx.JSON(iris.StatusOK, exception.ParamException())
		return
	}

	//process
	success := service.UpdateUserInfo(user.User{UserName: request.UserName, Pwd: request.Pwd})
	if success {
		ctx.JSON(iris.StatusOK, iris.Map{"code": -1})
		return
	}

	//response
	response := &user.UserResponse{Code: 0, Msg: "修改成功",}

	ctx.JSON(iris.StatusOK, response)
}

//添加用户
func add(ctx *iris.Context) {
	//request
	request := new(user.UserRequest)
	if err := ctx.ReadJSON(request); err != nil {
		ctx.JSON(iris.StatusOK, exception.ParamException())
		return
	}

	if request.UserName == "" || request.Pwd == "" {
		ctx.JSON(iris.StatusOK, exception.ParamException())
		return
	}

	//process
	success := service.AddUserInfo(request.UserName, request.Pwd)
	if success {
		ctx.JSON(iris.StatusOK, iris.Map{"code": -1})
		return
	}

	//response
	response := &user.UserResponse{Code: 0, Msg: "添加成功",}

	ctx.JSON(iris.StatusOK, response)
}
