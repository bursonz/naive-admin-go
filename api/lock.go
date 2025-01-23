package api

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"naive-admin-go/utils"
	"strconv"
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
		StationId:   params.StationId,
		AdminId:     params.AdminId,
		SN:          params.SN,
		Mac:         params.MAC,
		FactoryId:   params.FactoryId,
		CurrentKey:  params.CurrentKey,
		FactoryKey:  params.FactoryKey,
		Location:    params.Location,
		Description: params.Description,
		Enable:      params.Enable,
	}

	if err := db.Dao.Model(&model.Lock{}).Create(&newLock).Error; err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, "")
	}

}

func (lock) Update(c *gin.Context) {
	var params inout.PatchLockReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	orm := db.Dao.Model(&model.Lock{}).Where("id =?", params.Id)
	if params.StationId != nil {
		orm = orm.Update("station_id", params.StationId)
	}
	if params.AdminId != nil {
		orm = orm.Update("admin_id", params.AdminId)
	}
	if params.SN != nil {
		orm = orm.Update("sn", params.SN)
	}
	if params.MAC != nil {
		orm = orm.Update("mac", params.MAC)
	}
	if params.CurrentKey != nil {
		orm = orm.Update("current_key", params.CurrentKey)
	}
	if params.FactoryKey != nil {
		orm = orm.Update("factory_key", params.FactoryKey)
	}
	if params.Location != nil {
		orm = orm.Update("location", params.Location)
	}
	if params.Description != nil {
		orm = orm.Update("description", params.Description)
	}
	if params.Enable != nil {
		orm = orm.Update("enable", params.Enable)
	}
	if params.AlarmMode != nil {
		orm = orm.Update("alarm_mode", params.AlarmMode)
	}
	if params.Muted != nil {
		orm = orm.Update("muted", params.Muted)
	}
	if params.Hibernate != nil {
		orm = orm.Update("hibernate", params.Hibernate)
	}
	if err := orm.Error; err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, err)
	}
}

func (lock) List(c *gin.Context) {
	var data inout.LockListRes
	// 查询条件
	// TODO 其他条件
	var stationId = c.DefaultQuery("stationId", "")
	var mac = c.DefaultQuery("mac", "")
	var enable = c.DefaultQuery("enable", "")
	// 分页
	var pageNo, _ = strconv.Atoi(c.DefaultQuery("pageNo", "1"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 查询操作
	orm := db.Dao.Model(&model.Lock{})
	if stationId != "" {
		orm = orm.Where("station_id like ?", "%"+stationId+"%")
	}
	if mac != "" {
		orm = orm.Where("mac like ?", "%"+mac+"%") // 模糊查询
	}
	if enable != "" {
		orm = orm.Where("enable = ?", enable)
	}
	orm.Count(&data.Total)
	orm.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&data.PageData)

	// 返回
	Resp.Succ(c, data)
}

func (lock) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		tx.Where("id =?", id).Delete(&model.Lock{})
		// 级联删除与该锁相关的工单、锁事件
		var oldOrderSteps []model.OrderStep
		tx.Where("lock_id =?", id).Find(&oldOrderSteps)
		for _, oldOrderStep := range oldOrderSteps {
			tx.Where("order_id =?", oldOrderStep.OrderId).Delete(&model.OrderApproval{})
			tx.Where("order_id =?", oldOrderStep.OrderId).Delete(&model.Order{})
			tx.Where("order_id =?", oldOrderStep.OrderId).Delete(&model.OrderStep{})
		}
		return nil
	}); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	} else {
		Resp.Succ(c, "删除成功")
	}
}

func (lock) BatchDelete(c *gin.Context) {
	var params inout.BatchDeleteReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	if err := db.Dao.Transaction(func(tx *gorm.DB) error {
		for _, id := range params.Ids {
			tx.Where("id =?", id).Delete(&model.Lock{})
			var oldOrderSteps []model.OrderStep
			tx.Where("lock_id =?", id).Find(&oldOrderSteps)
			for _, oldOrderStep := range oldOrderSteps {
				tx.Where("order_id =?", oldOrderStep.OrderId).Delete(&model.OrderApproval{})
				tx.Where("order_id =?", oldOrderStep.OrderId).Delete(&model.Order{})
				tx.Where("order_id =?", oldOrderStep.OrderId).Delete(&model.OrderStep{})
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

func (lock) Command(c *gin.Context) {
	// 参数准备
	var params inout.LockCommandReq
	if err := c.BindJSON(&params); err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	// 获取记录
	var record model.Lock
	var mac, key []byte
	if err := db.Dao.Where("id=?", params.Id).Find(&record).Error; err != nil {
		Resp.Err(c, 20001, "查找锁记录失败")
		return
	} else if mac, err = hex.DecodeString(record.Mac); err != nil {
		Resp.Err(c, 20001, "Mac地址错误，请联系管理员重置该锁")
		return
	} else if key, err = hex.DecodeString(record.CurrentKey); err != nil {
		Resp.Err(c, 20001, "密钥错误，请联系管理员重置该锁")
		return
	}
	// 更新记录
	if params.Cmd != nil {
		// TODO 解析命令
		cmd := *params.Cmd
		switch cmd[0] {
		case 0x01:
			record.HardwareVersion = hex.EncodeToString(cmd[10:11])
			record.SoftwareVersion = hex.EncodeToString(cmd[11:13])
			record.FactoryId = hex.EncodeToString(cmd[13:17])
			record.AlarmMode = hex.EncodeToString(cmd[17:18])
			record.LockStatus = hex.EncodeToString(cmd[18:19])
			record.BackupData = hex.EncodeToString(cmd[19:23])
			record.NewLock = cmd[23] == 0x55 //TODO 旧锁二次添加，目前不需要
			record.UnlockRecord = hex.EncodeToString(cmd[24:26])
			record.Power = hex.EncodeToString(cmd[26:27])
			record.Muted = hex.EncodeToString(cmd[27:28])
			record.Hibernate = hex.EncodeToString(cmd[28:29])
			break
		case 0xE0:
			switch cmd[11] {
			case 0x01:
				// 开锁成功
				record.LockStatus = hex.EncodeToString([]byte{0x01})
				break
			case 0x05:
				Resp.Err(c, 20001, "开锁失败,MAC地址错误")
				return
			}
			break
		default:
			Resp.Err(c, 20001, "未支持的命令")
			return
		}
		// 更新数据库
		if err := db.Dao.Where("id=?", params.Id).Updates(&record).Error; err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
	}
	// 生成命令
	if params.Type != nil {
		newCommand := utils.GenerateCommand(*params.Type, params.Roll, mac, key)
		Resp.Succ(c, newCommand) // TODO 改为前端解析的格式
	} else {
		Resp.Succ(c, "更新成功")
	}
}
