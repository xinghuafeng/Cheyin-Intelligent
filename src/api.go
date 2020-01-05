package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Person struct {
	name string `json:"name" form:"name"`
	tel  string `json:"tel" form:"tel"`
}

func APIRouter() *gin.Engine {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/tisaninfo?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	router := gin.Default()
	router.GET("compuer", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "北京车音智能科技有限公司",
		})
	})
	//新增用户信息接口
	router.POST("/person", func(c *gin.Context) {
		name := c.Request.FormValue("name")
		tel := c.Request.FormValue("tel")

		rs, err := db.Exec("INSERT INTO userinfo(name, tel) VALUES (?, ?)", name, tel)
		if err != nil {
			log.Fatalln(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert userinfo id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	})
	//查询用户信息接口
	router.GET("/persons", func(c *gin.Context) {
		rows, err := db.Query("SELECT  name, tel FROM userinfo")
		defer rows.Close()

		if err != nil {
			log.Fatalln(err)
		}

		persons := make([]Person, 0)

		for rows.Next() {
			var person Person
			rows.Scan(&person.name, &person.tel)
			persons = append(persons, person)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"persons": persons,
		})

	})

	return router

}
