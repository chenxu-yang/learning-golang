package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Gender string
	Age int
}
func sayHello(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err!=nil{
		fmt.Println("failed,err:%v",err)
	}
	//渲染模板
	u1:=User{
		Name: "chenxu",
		Gender: "male",   //首字母大写：公用
		Age: 22,
	}
	m1:=map[string]interface{}{
		"Name":"chen",
		"Gender":"female",
		"Age":21,
	}
	hobbyList:=[]string{
		"sing",
		"dance",
		"basketball",
	}
	t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})
}
func main(){
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000",nil)
	if err!=nil{
		fmt.Println("http server start failed,err:%v",err)
		return
	}
}