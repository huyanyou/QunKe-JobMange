package client

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 根管理员
type Root struct {
	SysManger []SysManger
}

// 系统管理员结构体
type SysManger struct {
	R_position uint8
	Login      Login
}

// root创建管理员
func (r *Root) CreateSysManger(sys *SysManger) {
	r.SysManger = append(r.SysManger, *sys)
	sys.R_position = uint8(len(r.SysManger)) - 1
}

// root删除管理员
func (r *Root) DeleteSysManger(position uint8) {
	for i := position; i < uint8(len(r.SysManger)); i++ {
		if i == uint8(len(r.SysManger)-1) {
			r.SysManger = r.SysManger[:i]
			break
		}
		r.SysManger[i] = r.SysManger[i+1]
		r.SysManger[i].R_position = i
	}
}

var root = new(Root)

// 创建管理员
func CreateManger(c *gin.Context) {
	var login Login // 声明一个结构体
	// 将请求的body信息，解析到login中
	if err := c.BindJSON(&login); err != nil {
		c.JSON(400, gin.H{
			"msg": "参数错误",
		})
		return
	}
	var sys SysManger
	sys.Login = login
	root.CreateSysManger(&sys)
	fmt.Println(sys.R_position)
	c.JSON(200, gin.H{
		"msg":  "访问成功",
		"data": sys.R_position,
	})
}

// 删除管理员
func DeleteManger(c *gin.Context) {
	position, _ := strconv.Atoi(c.Param("position"))
	root.DeleteSysManger(uint8(position))
	fmt.Println(root)
	c.JSON(200, gin.H{
		"data": "删除成功",
	})
}
