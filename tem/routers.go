package tem

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Start 启动http服务
func Start(host, port string) {
	temHost = host
	temPort = port
	g := gin.Default()
	g.GET("/login", temLogin)
	go g.Run(host + ":" + port)
}

func temLogin(c *gin.Context) {
	code := c.Query("code")
	u, err := teambitionAPI.Login(code)
	if err != nil {
		c.String(http.StatusOK, "login:"+err.Error())
	} else {
		c.String(http.StatusOK, "login,user name:"+u.Name)
	}
}
