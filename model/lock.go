package model

type Lock struct {
	BaseModel
	// 填写信息
	StationId   uint   `json:"stationId" gorm:"type:bigint unsigned;not null;index;comment:站点id"`
	AdminId     uint   `json:"adminId" gorm:"type:bigint unsigned;not null;index;comment:管理员id"`
	Name        string `json:"name" gorm:"type:varchar(32);notNull;uniqueIndex:idx_lock_name;comment:锁名称"`
	Location    string `json:"location" gorm:"type:text;default:null;comment:锁具位置"`
	Description string `json:"description" gorm:"type:text;default:null;comment:锁具描述"`
	// 设备信息
	SN         string `json:"sn" gorm:"type:varchar(8);uniqueIndex:idx_lock_sn;not null;comment:序列号"`
	Mac        string `json:"mac" gorm:"type:varchar(12);uniqueIndex:idx_lock_mac;not null;comment:MAC地址"`
	CurrentKey string `json:"currentKey" gorm:"type:varchar(32);not null;comment:当前密钥"`
	FactoryKey string `json:"factoryKey" gorm:"type:varchar(32);not null;comment:出厂密钥"`
	// 设备状态
	Enable bool `json:"enable" gorm:"type:tinyint(1);default:1;comment:启用状态 1: 启用 0: 禁用"`
	// 设备信息
	SoftwareVersion string `json:"softwareVersion" gorm:"type:varchar(4);default:null;comment:软件版本"`
	HardwareVersion string `json:"hardwareVersion" gorm:"type:varchar(2);default:null;comment:硬件版本"`
	FactoryId       string `json:"factoryId" gorm:"type:varchar(8);default:null;comment:工厂ID"`
	AlarmMode       string `json:"alarmMode" gorm:"type:varchar(2);default:null;comment:报警模式"`
	LockStatus      string `json:"lockStatus" gorm:"type:varchar(2);default:null;comment:锁状态"`
	BackupData      string `json:"backupDate" gorm:"type:varchar(8);default:null;comment:备份日期"`
	NewLock         string `json:"newLock" gorm:"type:varchar(2);default:null;comment:新锁标识"`
	UnlockRecord    string `json:"unlockRecord" gorm:"type:varchar(4);default:null;comment:解锁记录"`
	Power           string `json:"power" gorm:"type:varchar(2);default:null;comment:电量状态"`
	Muted           string `json:"muted" gorm:"type:varchar(2);default:null;comment:静音状态"`
	Hibernate       string `json:"hibernate" gorm:"type:varchar(2);default:null;comment:休眠状态"`
}

func (Lock) TableName() string { return "lock" }
