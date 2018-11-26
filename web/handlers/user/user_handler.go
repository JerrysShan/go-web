package user

import (
	"go-web/schemas"
	"go-web/services/user"
	"go-web/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")
	uid := c.Query("uid")
	schema := &schemas.User{}
	if name != "" {
		schema.Name = name
	}
	if id != "" {
		if t, err := strconv.Atoi(id); err == nil {
			schema.ID = t
		}
	}
	if uid != "" {
		if t, err := strconv.ParseInt(uid, 10, 64); err == nil {
			schema.UID = t
		}
	}
	if data, err := user.Query(schema); err != nil {
		utils.Logger.Error("查询账户出现错误", err)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "查询出错"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": "查询成功"})
	}
}

func Show(c *gin.Context) {

}

func Create(c *gin.Context) {
	var schema schemas.User
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误" + err.Error()})
		return
	}

	err := user.Create(schema)
	if err != nil {
		utils.Logger.Error("查询用户出现错误", err)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "创建成功"})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
