package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	r:=gin.Default()
	//r.GET("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"method":"GET",
	//	})
	//})
	//r.POST("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"method":"POST",
	//	})
	//})
	//r.PUT("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"method":"PUT",
	//	})
	//})
	//r.DELETE("/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{
	//		"method":"DELETE",
	//	})
	//})
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK,gin.H{"method":"GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK,gin.H{"method":"POST"})
		case http.MethodPut:
			c.JSON(http.StatusOK,gin.H{"method":"PUT"})
		case http.MethodDelete:
			c.JSON(http.StatusOK,gin.H{"method":"DELETE"})

		}
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{"msg":"not found"})
	})
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
	//})
	//r.GET("/video/xx", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
	//})
	//路由组  提取公园前缀
	videoGroup:=r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
	}
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"msg":"/shop/index"})
	})
	r.Run(":9090")
}
