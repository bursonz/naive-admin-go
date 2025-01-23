package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"strconv"
)

type station struct{}

var Station = &station{}

// List 获取所有站点
func (station) List(c *gin.Context) {
	var data inout.StationListRes
	// 查询条件 TODO 其他条件
	var code = c.DefaultQuery("code", "")
	var name = c.DefaultQuery("name", "")
	var stationType = c.DefaultQuery("stationType", "")
	var pageNoReq = c.DefaultQuery("pageNo", "1")
	var pageSizeReq = c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(pageNoReq)
	pageSize, _ := strconv.Atoi(pageSizeReq)
	// 条件查询
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
	// 查询总数
	orm.Count(&data.Total)
	// 分页查询
	if pageNo < 1 {
		orm.Find(&data.PageData)
	} else {
		orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&data.PageData)
	}
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
		tx.Where("station_id=?", sid).Delete(&model.Lock{})
		var oldOrders []model.Order
		tx.Where("station_id=?", sid).Find(&oldOrders)
		for _, odr := range oldOrders {
			tx.Where("order_id=?", odr.ID).Delete(&model.OrderApproval{})
			tx.Where("order_id=?", odr.ID).Delete(&model.OrderStep{})
			tx.Where("order_id=?", odr.ID).Delete(&model.Order{})
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	Resp.Succ(c, "")
}

func (station) BatchDelete(c *gin.Context) {
	var params inout.BatchDeleteReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	//ids := c.QueryArray("id")
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		for _, id := range params.Ids {
			tx.Where("id =?", id).Delete(&model.Station{})
			tx.Where("station_id=?", id).Delete(&model.Lock{})
			var oldOrders []model.Order
			tx.Where("station_id=?", id).Find(&oldOrders)
			for _, odr := range oldOrders {
				tx.Where("order_id=?", odr.ID).Delete(&model.OrderApproval{})
				tx.Where("order_id=?", odr.ID).Delete(&model.OrderStep{})
				tx.Where("order_id=?", odr.ID).Delete(&model.Order{})
			}
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, "批量删除成功")
	}
}
