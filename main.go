package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type result struct{
	gitlabPassword string
	shPath string
	shName string
	port string
}

func initcfg()  result {
	cfg, err := goconfig.LoadConfigFile("./cfg.ini")
	if err != nil {
		fmt.Printf("无法加载配置文件：%s\n", err)
	}
	gitlabPassword,_:= cfg.GetValue("", "gitlab_password")
	shPath,_:= cfg.GetValue("", "sh_path")
	shName,_:= cfg.GetValue("", "sh_name")
	port,_:= cfg.GetValue("", "port")
	if gitlabPassword == "" || shPath== "" || shName == ""{

	}
	rs := result{gitlabPassword, shPath, shName, port}
	return rs
}

func main() {
	config := initcfg()
	r := gin.Default()
	r.POST("/gitlabapi", func(context *gin.Context) {
		context.Request.ParseMultipartForm(128)//保存表单缓存的内存大小128M
		header := context.GetHeader("X-Gitee-Token:")
		if header != config.gitlabPassword {
			fmt.Println("Password Dont Match")
			return
		}
		cmd := exec.Command("sh", config.shPath+config.shName)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		context.String(http.StatusOK, fmt.Sprintf("success"))
	})

	r.Run(":" + config.port) // 监听并在 0.0.0.0:8080 上启动服务
}