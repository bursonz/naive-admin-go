package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
	"time"
)

type station struct{}

var Station = &station{}

// List 获取所有站点
func (station) List(c *gin.Context) {
	var data []model.Station
	// 查询条件 TODO 其他条件
	var code = c.DefaultQuery("code", "")
	var name = c.DefaultQuery("name", "")
	var stationType = c.DefaultQuery("stationType", "")
	// 分页
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	// 查询
	var orm = db.Dao.Model(&model.Station{})
	if code != "" {
		orm = orm.Where("code = ?", code)
	}
	if name != "" {
		orm = orm.Where("name = ?", name)
	}
	if stationType != "" {
		orm = orm.Where("station_type = ?", stationType)
	}
	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&data)
	//db.Dao.Model(&model.Station{}).Find(&data)
	Resp.Succ(c, data)
}

func (station) Update(c *gin.Context) {
	var params inout.PatchStationReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Model(&model.Station{}).Where("id=?", params.Id).Updates(&model.Station{
		Code:        params.Code,
		Name:        params.Name,
		AdminUserId: params.AdminUserId,
		Location:    params.Location,
		StationType: params.StationType,
		UpdateTime:  time.Now(),
	}).Error; err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, err)
	}
}

func (station) Add(c *gin.Context) {
	var params inout.AddStationReq

	if err := c.Bind(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}

	if err := db.Dao.AutoMigrate(&model.Station{}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	var newStation = model.Station{
		Code:        params.Code,
		Name:        params.Name,
		AdminUserId: params.AdminUserId,
		Location:    params.Location,
		StationType: params.StationType,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	if err := db.Dao.Create(&newStation).Error; err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}

	Resp.Succ(c, "")
}

func (station) Delete(c *gin.Context) {
	sid := c.Param("id")
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", sid).Delete(&model.Station{})
		// TODO 级联删除与站点相关的锁、工单
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	Resp.Succ(c, "")
}
