package main
import(
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	r:=gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//get request query string from browser
		name:=c.Query("query")//method 1
		age:=c.Query("age")
		//name:=c.DefaultQuery("query","somebody")//method 2
		//name,ok:=c.GetQuery("query")  //method 3
		//if !ok{
		//	//
		//	name="somebody"
		//}
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})
	r.Run(":9090")
}
