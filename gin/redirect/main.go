package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	r:=gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"OK",
		//})
		//redirect to google
		c.Redirect(http.StatusMovedPermanently,"http://www.google.com")
	})
	r.GET("/a", func(c *gin.Context) {
		//redirect to b
		c.Request.URL.Path="/b" //change the path
		r.HandleContext(c)   //continue

	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"ojbk",
		})
	})
	r.Run(":9090")
}