package role

import (
	"go-web/schemas"
	"go-web/services/role"
	"go-web/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	var schema schemas.Role
	if id != "" {
		if ID, err := strconv.Atoi(id); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
			return
		} else {
			schema.ID = ID
		}
	}

	if name != "" {
		schema.Name = name
	}
	if data, err := role.Query(&schema); err != nil {
		utils.Logger.Error("查询角色出错", data)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "查询出错"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": "OK"})
		return
	}
}

func Update(c *gin.Context) {
	var schema schemas.Role
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if schema.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if schema.Name == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if affected, err := role.Update(&schema); affected == 0 || err != nil {
		utils.Logger.Error("更新出错", err, affected)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "更新出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "OK"})
}

func Create(c *gin.Context) {
	var schema schemas.Role
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if schema.Name == "" {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if len(schema.ResourceIds) == 0 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if exist, err := role.IsExist(schema.Name); err != nil {
		utils.Logger.Error("创建角色出错", err)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建出错"})
		return
	} else if exist {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "角色名字已经存在"})
		return
	}
	affected, ID, err := role.Insert(&schema)
	if affected == 0 || err != nil {
		utils.Logger.Error("创建角色出错", err, affected)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建出错"})
		return
	}
	affected, err = role.CreateResources(ID, schema.ResourceIds)
	if affected == 0 || err != nil {
		utils.Logger.Error("创建角色出错", err, affected)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "创建出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "OK"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	if ID, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	} else {
		if affected, err := role.Delete(ID); affected == 0 || err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "删除出错"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "OK"})
}
