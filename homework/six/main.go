package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func main() {
	conn := connectDb()
	var server ServerInterface = DBConnect(conn)

	r := gin.Default()
	pprof.Register(r)
	//对接注册
	r.POST("/save", func(c *gin.Context) {
		var ps *CircleState
		if err := c.BindJSON(&ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "无法绑定数据" + err.Error(),
			})
			return
		}
		if err := server.SaveData(ps); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "添加数据失败！",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	//对接更新
	r.PUT("/update", func(c *gin.Context) {
		var cs *CircleState
		if err := c.BindJSON(&cs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "更新失败",
			})
			return
		}
		if err := server.UpdateData(cs); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "更新数据失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				"data":    cs,
			})
		}
	})
	r.GET("/info/all", func(c *gin.Context) {
		if data, err := server.GetDataList(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "查询所有数据失败！",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				"date":    data,
			})
		}
	})
	r.GET("/info/:id", func(c *gin.Context) {
		id := c.Param("id")
		idString, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("转换失败", err)
		}
		if ps, err := server.GetOneData(int64(idString)); err != nil { //通过
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "获取数据失败!",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
				"date":    ps,
			})
		}
	})
	r.DELETE("/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		idString, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("转换失败", err)
		}
		if err := server.DeleteDataByDeleted(int64(idString)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "删除失败！",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": "删除成功",
			})
		}
	})

	if err := r.Run(":9999"); err != nil {
		log.Fatal("启动错误：%v\n", err)
	}
}

func connectDb() *gorm.DB {
	//parseTime=true解決unsupported Scan, storing driver.Value type []uint8 into type *time.Time问题
	conn, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/learngo?parseTime=true")) //与数据库建立连接
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func (*CircleState) GetTableName() string {
	return "t_circle_state"
}
