package test

import (
	"blog/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestGet(c *gin.Context) {
	tasks := []map[string]interface{}{}
	item := map[string]interface{}{"id": 1, "name": "task1", "done": true}
	item1 := map[string]interface{}{"id": 2, "name": "task2", "done": false}
	tasks = append(tasks, item)
	tasks = append(tasks, item1)
	c.JSON(http.StatusOK, gin.H{
		"code":       common.Success,
		"msg":        "ok",
		"data":       tasks,
		"pagination": nil,
	})
}