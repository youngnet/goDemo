package main

// baseType bool string int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64
// byte=uint8  rune=int32  float32 float64 complex64 complex128

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// User userInfo
type User struct {
	Username string
	Passwd   string
	Age      int
}

func cal(num int) (ans float64) {
	ans = float64(num) * 4 / 9
	return
}

var publicPath = "../web/dist"

func main() {
	// router := gin.Default()
	var router = gin.Default()
	// z := cmplx.Sqrt(-5 + 12i)
	// router.LoadHTMLFiles("templates/*")
	router.Use(MiddleWare())
	router.StaticFS("/static", gin.Dir(publicPath+"/static", true))
	router.StaticFile("/manifest.json", "./build/manifest.json")
	router.GET("/api", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		fmt.Println(c.Param("path"))
		c.String(http.StatusOK, "hello world")
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.File(publicPath + "/index.html")
	})

	router.POST("/api/test", func(c *gin.Context) {
		var user User
		var err error
		user.Username = "young"
		user.Passwd = "123456"
		user.Age = 18
		user.Username = "yang"
		// err := c.Bind(&user)
		contentType := c.Request.Header.Get("Content-Type")

		switch contentType {
		case "application/json":
			err = c.BindJSON(&user)
		case "application/x-www-form-urlencoded":
			err = c.BindWith(&user, binding.Form)
		}
		fmt.Println(c.Query("id"))
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + c.FullPath())
		}()

		c.JSON(http.StatusOK, gin.H{
			"user":   user.Username,
			"passwd": user.Passwd,
			"age":    user.Age,
		})
	})

	router.POST("/multi/upload", func(c *gin.Context) {
		err := c.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
		}

		formdata := c.Request.MultipartForm

		files := formdata.File["upload"]

		// index, val := range Array
		for i := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			out, err := os.Create(files[i].Filename)
			defer out.Close()
			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}
			c.String(http.StatusOK, "upload success")
		}
	})

	router.Run(":8888")
}

// MiddleWare router中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.FullPath())
	}
}
