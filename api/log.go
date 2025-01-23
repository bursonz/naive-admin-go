package api

import (
	"github.com/gin-gonic/gin"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

type srvLog struct{}

var SrvLog = &srvLog{}

func (srvLog) List(c *gin.Context) {
	var data inout.LogListRes
	// 查询参数
	var userId = c.DefaultQuery("userId", "")
	var userName = c.DefaultQuery("userName", "")
	var method = c.DefaultQuery("method", "")
	var target = c.DefaultQuery("target", "")
	var ip = c.DefaultQuery("ip", "")
	var pageNo, _ = strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// 条件查询
	var orm = db.Dao.Model(&model.SysLog{})
	// 查询总数
	if userId != "" {
		orm.Where("user_id like ?", userId)
	}
	if userName != "" {
		orm.Where("user_name like ?", userName)
	}
	if method != "" {
		orm.Where("method like ?", method)
	}
	if target != "" {
		orm.Where("target like ?", target)
	}
	if ip != "" {
		orm.Where("ip like ?", ip)
	}
	orm.Count(&data.Total)
	// 分页查询
	if pageNo < 1 { // pageNo 小于1 时，查询所有
		orm.Find(&data.PageData)
	} else {
		orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&data.PageData)
	}
	Resp.Succ(c, data)

}
