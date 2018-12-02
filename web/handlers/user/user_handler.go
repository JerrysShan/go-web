package user

import (
	"fmt"
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
			fmt.Println(t)
			schema.UID = t
		}
	}
	if data, err := user.Query(schema); err != nil {
		utils.Logger.Error("查询账户出现错误:", err)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "查询出错"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": "查询成功"})
	}
}

func Create(c *gin.Context) {
	var schema schemas.User
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if exist, err := user.IsExist(schema.UID); err != nil {
		utils.Logger.Error("查询账户出现错误:", err)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建出错"})
		return
	} else if exist {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "账户已存在"})
		return
	}
	affected, ID, err := user.Create(&schema)
	if err != nil || affected == 0 {
		utils.Logger.Error("创建用户出现错误", err)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建出错"})
		return
	}
	affected, err = user.CreateRoles(ID, schema.RoleIds)
	if err != nil || affected == 0 {
		utils.Logger.Error("创建用户出现错误", err)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "创建成功"})
}

func Update(c *gin.Context) {
	var schema schemas.User
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if affected, err := user.Update(&schema); err != nil || affected == 0 {
		utils.Logger.Error("更新用户出错", err, affected)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "更新出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "更新成功"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	ID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	affectd, err := user.Delete(ID)
	if err != nil || affectd == 0 {
		utils.Logger.Error("删除用户出错", err, affectd)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "删除出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "删除成功"})
}
