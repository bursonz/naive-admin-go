package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"naive-admin-go/db"
	"naive-admin-go/model"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody)) // 重新写入
		}
		c.Next()
		uid, _ := c.Get("uid")
		uname, _ := c.Get("uname")
		newLog := &model.SysLog{
			BaseModel: model.BaseModel{},
			UserId:    uint(uid.(int)),
			UserName:  uname.(string),
			Method:    c.Request.Method,    // 请求方法
			Target:    c.FullPath(),        // 操作目标
			Content:   string(requestBody), // 请求内容
			IP:        c.ClientIP(),        // 请求IP
			Path:      c.Request.URL.Path,  // 请求路径
		}
		if err := db.Dao.Create(newLog).Error; err != nil {
			log.Println(newLog)
			log.Println(err)
		}
	}
}
