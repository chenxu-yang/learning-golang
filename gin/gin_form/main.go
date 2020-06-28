package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:=gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})
	// get the post request from the form
	r.POST("/login", func(c *gin.Context) {
		//method 1
		//username:=c.PostForm("username")
		//password:=c.PostForm("password")
		username:=c.DefaultPostForm("username","somebody")
		password:=c.DefaultPostForm("password","somebody")
		//username,ok:=c.GetPostForm("username")
		//if !ok{
		//
		//}
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name":username,
			"Password":password,
		})
	})
	r.Run(":9090")
}
