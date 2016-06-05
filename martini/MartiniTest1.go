package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

//有返回值的处理器
func validate() string {
	return "pass"
}

type Person struct {
	Name string
}
type Student struct {
	Name string
}
//获取注入Person的值
func personHandler(person *Person) string {
	fmt.Println(person)
	return person.Name
}

//使用martini.Params获得参数
func GetBook(params martini.Params){
	fmt.Println("正在根据"+params["id"]+"返回golang编程")
}

//post 里面的参数要通过http.Request获取
func AddBook(req * http.Request){
	fmt.Println("正在把"+req.FormValue("name")+"存入数据库")
}

func DelBook(params martini.Params){
	fmt.Println("正在删除"+params["id"]+"编程")
}

func UpdateBook(params martini.Params){
	fmt.Println("正在修改"+params["id"]+"编程")
}

//没有返回值的处理器
//func validate(){
//	fmt.Println("验证通过")
//}

func main() {
	m := martini.Classic()//创建一个典型的martini实例。
	m.Get("/", func() string{//基本使用
		return "Hello world，hello zmy!"
	})//接收对\的GET方法请求，第二个参数是对一请求的处理方法。
	//m.Get("/", func() (int,string){//带状态码的使用
	//	return 405,"状态码测试"
	//})//接收对\的GET方法请求，第二个参数是对一请求的处理方法。
	//m.Get("/hello/:name", func(params martini.Params)string{//请求中带有参数
	//	return "welcome ==> "+params["name"]
	//})
	//4 使用多个Handler
	// 注意.当多个Handler组合处理的时候,除了最后一个Handler之外,其它Handler不能有返回.换句话说,遇到return就己经终止处理流程,后面的Handler不再执行
	//m.Get("/auth",validate, func()string{
	//	return "恭喜你,通过验证"
	//},validate)
	//person := &Person{Name:"milo"}
	//m.Map(person)
	//m.Get("/person",personHandler)
	//m.Get("/req",func(c martini.Context){
	//	s := &Student{Name:"req"}
	//	c.Map(s)
	//	fmt.Println(s.Name)
	//})
	//m.Use(martini.Static("webSource"))
	//m.Use(func() {
	//	fmt.Println("加入中间件1")
	//})
	//m.Use(func() {
	//	fmt.Println("加入中间件2")
	//})
	//m.Get("/req",func(c martini.Context){
	//	s := &Student{Name:"req"}
	//	c.Map(s)
	//	fmt.Println(s.Name)
	//})
	//m.Group("/libaray", func(r martini.Router) {
	//		r.Get("/:id",GetBook)
	//		r.Post("/add",AddBook)
	//		r.Delete("/del/:id",DelBook)//用submit
	//		r.Put("/update/:id",UpdateBook)
	//	})
	//m.Use(martini.Static("assets"))
	m.RunOnAddr(":9000")
	m.Run()//运行服务器 默认运行在3000端口
}
// 通过http://localhost:3000可以访问


