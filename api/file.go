package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type file struct{}

var File = &file{}

func (file) Upload(c *gin.Context) {
	// 处理文件名
	filename := time.Now().Format("20060102150405") + "_" + c.Param("filename")
	// 从请求中获取上传的文件
	f, err := c.FormFile("file")
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	// 保存文件到本地
	if err = c.SaveUploadedFile(f, "./uploads/"+filename); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, filename)
	}
}

func (file) Download(c *gin.Context) {
	var filename = c.Param("filename")
	c.File("./uploads/" + filename)

}

func (file) Delete(c *gin.Context) {
	var filename = c.Param("filename")
	if err := os.Remove("./uploads/" + filename); err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, "删除成功")
	}
}
