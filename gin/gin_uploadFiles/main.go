package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main(){
	r:=gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//get file from request
		f,err:=c.FormFile("f1")    //
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			//save the file in local server
			dst:=fmt.Sprintf("./%s",f.Filename)
			c.SaveUploadedFile(f,dst)
			c.JSON(http.StatusOK,gin.H{
				"status":"OK",
			})
		}


	})
	r.Run(":9090")
}