package api

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"naive-admin-go/utils"
	"os"
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

	var oldCount int64
	db.Dao.Model(&model.Lock{}).Where("mac =?", params.MAC).Count(&oldCount)
	if oldCount > 0 {
		// 提示用户，该锁已经存在
		Resp.Err(c, 20001, "该锁已经存在,请勿重复添加！")
	}

	var newLock = model.Lock{
		StationId:       params.StationId,
		AdminId:         params.AdminId,
		SN:              params.SN,
		Mac:             params.MAC,
		Name:            params.Name,
		CurrentKey:      params.CurrentKey,
		FactoryKey:      params.FactoryKey,
		Location:        params.Location,
		Description:     params.Description,
		Enable:          params.Enable,
		SoftwareVersion: params.SoftwareVersion,
		HardwareVersion: params.HardwareVersion,
		FactoryId:       params.FactoryId,
		AlarmMode:       params.AlarmMode,
		LockStatus:      params.LockStatus,
		BackupData:      params.BackupData,
		NewLock:         params.NewLock,
		UnlockRecord:    params.UnlockRecord,
		Power:           params.Power,
		Muted:           params.Muted,
		Hibernate:       params.Hibernate,
	}
	if err := db.Dao.Model(&model.Lock{}).Create(&newLock).Error; err != nil {
		Resp.Err(c, 20001, err.Error())
	} else {
		Resp.Succ(c, newLock.ID)
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
	if params.Name != nil {
		orm = orm.Update("name", params.Name)
	}
	if params.SN != nil {
		orm = orm.Update("sn", params.SN)
	}
	if params.MAC != nil {
		orm = orm.Update("mac", params.MAC)
	}
	if params.FactoryId != nil {
		orm = orm.Update("factory_id", params.FactoryId)
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
	if params.SoftwareVersion != nil {
		orm = orm.Update("software_version", params.SoftwareVersion)
	}
	if params.HardwareVersion != nil {
		orm = orm.Update("hardware_version", params.HardwareVersion)
	}
	if params.AlarmMode != nil {
		orm = orm.Update("alarm_mode", params.AlarmMode)
	}
	if params.LockStatus != nil {
		orm = orm.Update("lock_status", params.LockStatus)
	}
	if params.BackupData != nil {
		orm = orm.Update("backup_data", params.BackupData)
	}
	if params.NewLock != nil {
		orm = orm.Update("new_lock", params.NewLock)
	}
	if params.UnlockRecord != nil {
		orm = orm.Update("unlock_record", params.UnlockRecord)
	}
	if params.Power != nil {
		orm = orm.Update("power", params.Power)
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
	var oldCmd []byte
	if params.Cmd != nil && *params.Cmd != "" {
		oldCmd, _ = hex.DecodeString(*params.Cmd)
	}
	var newKey []byte
	switch params.Type {
	case 0x10:
		if params.Key != nil && *params.Key != "" {
			newKey, _ = hex.DecodeString(*params.Key)
		} else {
			newKey, _ = hex.DecodeString(os.Getenv("LOCK_KEY"))
		}
		if len(newKey) != 16 {
			Resp.Err(c, 20001, "密钥长度错误，请输入正确的16字节密钥")
			break
		}
		fallthrough
	case 0x01, 0x02, 0x03, 0x13, 0x1F, 0xE0:
		// 生成命令
		newCmd := utils.GenerateCommand(params.Type, params.Roll, mac, key, newKey, oldCmd)
		Resp.Succ(c, inout.LockCommandRes{
			Cmd: hex.EncodeToString(newCmd),
			Key: hex.EncodeToString(newKey),
		})
	default:
		Resp.Err(c, 20001, "指令类型错误")
	}
}
