package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

func main() {

	config := api.DefaultConfig()
	config.Address = "123.249.32.232:3001"
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	// 创建服务实例
	registration := &api.AgentServiceRegistration{}
	registration.Name = "cowstudio/be-sample"
	registration.Address = "123.249.32.232"
	registration.Port = 3000

	check := &api.AgentServiceCheck{}
	check.HTTP = fmt.Sprintf("http://%s:%d/health", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "10s"
	check.DeregisterCriticalServiceAfter = "30s"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	// 创建gin引擎
	router := gin.Default()

	// 定义路由及处理函数
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World, Welcome to the Cow Studio!")
	})

	router.GET("/health", func(c *gin.Context) {
		fmt.Fprintln(c.Writer, "OK")
	})

	// 启动HTTP服务，默认监听地址为:8080
	router.Run("0.0.0.0:8080")
}
