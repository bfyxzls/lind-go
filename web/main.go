package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// https://blog.csdn.net/u013302168/article/details/126903057
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})

	// 路由参数
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	//curl http://localhost:9999/user/geektutu

	// 查询参数：users?name=xxx&role=xxx，role可选
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher") // query 参数默认值
		c.String(http.StatusOK, "%s is a %s", name, role)
	})
	// curl "http://localhost:9999/users?name=Tom&role=student"

	// POST请求
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		// 返回json格式
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	//  http://localhost:9999/form  -X POST -d 'username=geektutu&password=1234'

	// GET 和 POST 混合
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})
	// $ curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'

	// Map参数
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	//curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'

	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})
	/**
	$ curl -i http://localhost:9999/redirect
	HTTP/1.1 301 Moved Permanently
	Content-Type: text/html; charset=utf-8
	Location: /
	Date: Thu, 08 Aug 2019 17:22:14 GMT
	Content-Length: 36

	<a href="/">Moved Permanently</a>.

	$ curl "http://localhost:9999/goindex"
	Who are you?
	*/

	// html模板渲染
	type student struct {
		Name string
		Age  int8
	}
	r.LoadHTMLGlob("templates/*")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
