package main
import (
	"github.com.gin-gonic/gin"
	"net/http"
)
func main(){
	r:=gin.Default()
 	r.GET("/index",func(c *gin.Context){
 		c.HTML(http.StatusOK,"index.tmpl",gin.H{
 			"title":"chenxu",

		})
	})
}
