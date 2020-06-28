package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
func indexHandler(c *gin.Context){
	fmt.Printf("index in")
	name,ok:=c.Get("name")
	if !ok{
		name="unknow"
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":name,
	})
}

//定义一个中间件
func m1(c *gin.Context){
	fmt.Printf("m1 in...")
	start:=time.Now()
	go funcXX(c.Copy())//在funcXX中只能使用c的拷贝
	c.Next()//调用后续的函数
	cost:=time.Since(start)
	fmt.Printf("cost time:%v\n",cost)

	fmt.Println("m1 out ")
}
func m2(c *gin.Context){
	fmt.Println("m2 in ...")
	c.Set("name","chenxu")  //设置值， 跨中间件存取值
	//c.Abort()//阻止调用后续的处理函数
	fmt.Println("m2 out ...")
}

func authMiddleware(turnon bool)gin.HandlerFunc{
	//connect to database
	//other prepare work
	return func(c *gin.Context) {
		if turnon{
			//是否登录
			//if 是登录用户
			//c.Next()
			//else
			//c.Abort()
		}

	}
}

func main(){
	r:=gin.Default()
	r.Use(m1,m2)

	r.GET("/index", indexHandler)
	r.GET("shop",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"shop",
		})
	})
	r.GET("user", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"user",
		})
	})
	//xxGroup:=r.Group("/xx",authMiddleware(true)){
	//	xxGroup.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK,gin.H{
	//			"msg":"xxGroup",
	//		})
	//	})
	//}
	r.Run(":8080")

}
