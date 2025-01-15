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

type lock struct{}

var Lock = &lock{}

func (lock) Add(c *gin.Context) {
	var params inout.AddLockReq

	if err := c.Bind(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}

	var newLock = model.Lock{
		StationId:       params.StationId,
		AdminId:         params.AdminId,
		SN:              params.SN,
		Mac:             params.MAC,
		FactoryId:       params.FactoryId,
		CurrentKey:      params.CurrentKey,
		FactoryKey:      params.FactoryKey,
		SoftwareVersion: params.SoftwareVersion,
		HardwareVersion: params.HardwareVersion,
		Location:        params.Location,
		Power:           params.Power,
		Description:     params.Description,
		Enable:          params.Enable,
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
	}

	if err := db.Dao.Model(&model.Lock{}).Create(&newLock).Error; err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}

	Resp.Succ(c, "")
}

func (lock) Update(c *gin.Context) {
	var params inout.PatchLockReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Model(&model.Lock{}).Where("id=?", params.Id).Updates(&model.Lock{
		StationId:       params.StationId,
		AdminId:         params.AdminId,
		SN:              params.SN,
		Mac:             params.MAC,
		FactoryId:       params.FactoryId,
		CurrentKey:      params.CurrentKey,
		FactoryKey:      params.FactoryKey,
		SoftwareVersion: params.SoftwareVersion,
		HardwareVersion: params.HardwareVersion,
		Location:        params.Location,
		Power:           params.Power,
		Description:     params.Description,
		Enable:          params.Enable,
		UpdateTime:      time.Now(),
	}).Error; err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, err)
	}
}

func (lock) Delete(c *gin.Context) {
	lid := c.Param("id")
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", lid).Delete(&model.Lock{})
		// TODO 级联删除与该锁相关的工单、锁事件、
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	Resp.Succ(c, "")
}

func (lock) List(c *gin.Context) {
	var data inout.LockListRes
	// 查询条件 TODO 其他条件
	var stationId = c.DefaultQuery("stationId", "")
	var mac = c.DefaultQuery("mac", "")
	var enable = c.DefaultQuery("enable", "")
	// 分页
	var pageNo, _ = strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 查询操作
	orm := db.Dao.Model(&model.Lock{})
	if stationId != "" {
		orm = orm.Where("station_id=?", stationId)
	}
	if mac != "" {
		orm = orm.Where("mac=?", mac)
	}
	if enable != "" {
		orm = orm.Where("enable = ?", enable)
	}
	orm.Count(&data.Total)
	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&data.PageData)

	// 返回
	Resp.Succ(c, data)
}

func (lock) BatchDelete(c *gin.Context) {
	ids := c.QueryArray("id")
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			tx.Where("id =?", id).Delete(&model.Lock{})
			var oldOrderSteps []model.OrderStep
			tx.Where("lock_id =?", id).Find(&oldOrderSteps)
			for _, orderStep := range oldOrderSteps {
				tx.Where("order_id =?", orderStep.OrderId).Delete(&model.OrderApproval{})
				tx.Where("order_id =?", orderStep.OrderId).Delete(&model.Order{})
				tx.Where("order_id =?", orderStep.OrderId).Delete(&model.OrderStep{})
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
