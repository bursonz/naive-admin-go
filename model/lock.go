package model

type Lock struct {
	BaseModel
	// 填写信息
	StationId   uint   `json:"stationId" gorm:"type:bigint unsigned;notNull;comment:站点id"`
	AdminId     uint   `json:"adminId" gorm:"type:bigint unsigned;notNull;comment:管理员id"`
	Location    string `json:"location" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	// 设备信息
	SN         string `json:"sn" gorm:"type:varchar(8)"`
	Mac        string `json:"mac" gorm:"type:varchar(12)"`
	CurrentKey string `json:"currentKey" gorm:"type:varchar(32)" comment:"当前密钥"`
	FactoryKey string `json:"factoryKey" gorm:"type:varchar(32)" comment:"出厂密钥"`
	// 设备状态
	Enable bool `json:"enable" comment:"1: 启用 0: 禁用"`
	// 01 获取
	SoftwareVersion string `json:"softwareVersion" gorm:"type:varchar(4)"`
	HardwareVersion string `json:"hardwareVersion" gorm:"type:varchar(2)"`
	FactoryId       string `json:"factoryId" gorm:"type:varchar(8)"`
	AlarmMode       string `json:"alarmMode" gorm:"type:varchar(2)"`
	LockStatus      string `json:"lockStatus" gorm:"type:varchar(2)"`
	BackupData      string `json:"backupDate" gorm:"type:varchar(8)"`
	NewLock         bool   `json:"newLock" gorm:"type:varchar(2)"`
	UnlockRecord    string `json:"unlockRecord" gorm:"type:varchar(4)"`
	Power           string `json:"power" gorm:"type:varchar(2)"`
	Muted           string `json:"muted" gorm:"type:varchar(2)"`
	Hibernate       string `json:"hibernate" gorm:"type:varchar(2)"`
}

func (Lock) TableName() string { return "lock" }
