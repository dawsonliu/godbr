package main

import (
	"fmt"

	"github.com/dawsonliu/dbrgo/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	//LoadApis()

	LoadCsis()
	Start()
}

func serve() {
	r := gin.Default()
	r.GET("/query/:sql", func(c *gin.Context) {
		sql := c.Param("sql")
		result, ok := repository.Invoke(sql)

		if ok == nil {
			c.JSON(200, result)
		} else {
			fmt.Println("failed: ", ok)
		}
	})

	r.Run()
}
