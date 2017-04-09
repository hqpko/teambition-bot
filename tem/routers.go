package tem

import "github.com/gin-gonic/gin"

import "net/http"

//Start 启动http服务
func Start(host, port string) {
	temHost = host
	temPort = port
	g := gin.Default()
	g.GET("/login", temLogin)
	go g.Run(host + ":" + port)
}

func temLogin(c *gin.Context) {
	c.String(http.StatusOK, "")
}
