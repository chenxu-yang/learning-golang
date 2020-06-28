package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type UserInfor struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"age"`
}
func main(){
	r:=gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user",func(c *gin.Context){
		//username:= c.Query("username")
		//password:=c.Query("password")
		//u:=UserInfor{
		//	username,
		//	password,
		//}
		var u UserInfor //declare an userinfor
		err:=c.ShouldBind(&u)//want to change value inside a func, need pointer
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
	r.POST("/form", func(c *gin.Context) {
		var u UserInfor //declare an userinfor
		err:=c.ShouldBind(&u)//want to change value inside a func, need pointer
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})
	r.POST("/json", func(c *gin.Context) {
		var u UserInfor //declare an userinfor
		err:=c.ShouldBind(&u)//want to change value inside a func, need pointer
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})
	r.Run(":9090")
}
