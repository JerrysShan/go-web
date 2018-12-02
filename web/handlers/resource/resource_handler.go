package resource

import (
	"go-web/schemas"
	"go-web/services/resource"
	"go-web/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	id := c.Query("id")
	code := c.Query("code")
	name := c.Query("name")
	schema := schemas.Resource{}
	if id != "" {
		if ID, err := strconv.Atoi(id); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
			return
		} else {
			schema.ID = ID
		}
	}
	if code != "" {
		schema.Code = code
	}
	if name != "" {
		schema.Name = name
	}
	if data, err := resource.Query(&schema); err != nil {
		utils.Logger.Error("查询资源出错", err)
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "查询出错"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "OK", "data": data})
	}
}

func Update(c *gin.Context) {
	var schema schemas.Resource
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if affected, err := resource.Update(&schema); err != nil || affected == 0 {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "更新出错"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "OK"})
}

func Create(c *gin.Context) {
	var schema schemas.Resource
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "参数错误"})
		return
	}
	if affected, err := resource.Insert(&schema); err != nil || affected == 0 {
		utils.Logger.Error("创建出错", err, affected)
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
		if affected, err := resource.Delete(ID); err != nil || affected == 0 {
			utils.Logger.Error("删除出错", err, affected)
			c.JSON(http.StatusOK, gin.H{"code": -1, "msg": "删除出错"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "OK"})
}
