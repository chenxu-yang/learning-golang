package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	r:=gin.Default()
	r.GET("/json",func(c *gin.Context){
		//method 1,using map
		//data:=map[string]interface{}{
		//	"name":"chenxu",
		//	"message":"hello",
		//	"age":18,
		//}
		//gin.H-> map[string]interface{}
		data:=gin.H{"name":"chenxu",
			"message":"hello",
			"age":18}
		//method 2 struct
		c.JSON(http.StatusOK,data)
	})
	type msg struct{
		//灵活使用tag来做对结构体字段做定制化操作
		Name string   `json:"name"`//变小写
		Message string
		Age int
	}
	r.GET("/another", func(c *gin.Context) {
		data:=msg{
			"chenxu",
			"niubi",
			23,
		}
		c.JSON(http.StatusOK,data)
	})
	r.Run(":9090")
}
