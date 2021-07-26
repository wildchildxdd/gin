package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

func year_check(number int) bool {
	var resault bool
	var division400, division4, division100 int

	division400 = number % 400
	division4 = number % 4
	division100 = number % 100


	if division400 == 0 {
		resault = true
	}else {

		if division4 == 0 {
			if division100 != 0 {
				resault = true
			} else {
				resault = false
			}
		} else {
			resault = false
		}

	}
	return resault
}


func get_time() string {
	return time.Now().String()
}

func upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "/file/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}


func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")
	router.MaxMultipartMemory = 8 << 20  // 8 MiB

	router.POST("/file", upload)
	router.StaticFS("/file", http.Dir("public"))


	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/post", func(c *gin.Context) {
		year, _ := strconv.Atoi(c.PostForm("year"))
		leap := year_check(year)
		fmt.Println(leap)
		if leap {
			c.String(http.StatusOK, "YES")
		} else {
			c.String(http.StatusOK, "NO")
		}

	})
	router.GET("/time", func(c *gin.Context) {
		c.String(http.StatusOK, get_time())
	})



	router.Run("localhost:8080")
}