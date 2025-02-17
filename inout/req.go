package inout

type LoginReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Captcha  string `form:"captcha" binding:"required"`
}
type AuthPwReq struct {
	NewPassword string `form:"newPassword" binding:"required"`
	OldPassword string `form:"oldPassword" binding:"required"`
}
type PatchUserReq struct {
	Id       uint    `json:"id"  binding:"required"`
	Enable   *bool   `json:"enable,omitempty"`
	RoleIds  *[]uint `json:"roleIds,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}
type PatchProfileUserReq struct {
	Id       uint    `json:"id"  binding:"required"`
	Gender   *int    `json:"gender"`
	NickName *string `json:"nickName"`
	Address  *string `json:"address"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}
type EnableRoleReq struct {
	Enable bool `json:"enable" binding:"required"`
	Id     uint `json:"id"`
}

type AddUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Enable   bool   `json:"enable" binding:"required"`
	RoleIds  []uint `json:"roleIds" binding:"required"`
}

type AddRoleReq struct {
	Code          string `json:"code" binding:"required"`
	Enable        bool   `json:"enable"`
	Name          string `json:"name" binding:"required"`
	PermissionIds []uint `json:"permissionIds"`
}
type PatchRoleReq struct {
	Id            uint    `json:"id"  binding:"required"`
	Code          *string `json:"code,omitempty"`
	Enable        *bool   `json:"enable,omitempty"`
	Name          *string `json:"name,omitempty"`
	PermissionIds *[]uint `json:"permissionIds,omitempty"`
}

type PatchRoleOpeateUserReq struct {
	Id      uint   `json:"id" `
	UserIds []uint `json:"userIds"`
}

type AddPermissionReq struct {
	Type      string `json:"type" binding:"required"`
	ParentId  *uint  `json:"parentId"`
	Name      string `json:"name" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Path      string `json:"path"`
	Icon      string `json:"icon"`
	Layout    string `json:"layout"`
	Component string `json:"component"`
	Show      bool   `json:"show"`
	Enable    bool   `json:"enable"`
	KeepAlive bool   `json:"keepAlive"`
	Order     int    `json:"order"`
}

type PatchPermissionReq struct {
	Id        uint   `json:"id"  binding:"required"`
	Type      string `json:"type" binding:"required"`
	ParentId  *uint  `json:"parentId"`
	Name      string `json:"name" binding:"required"`
	Code      string `json:"code" binding:"required"`
	Path      string `json:"path"`
	Icon      string `json:"icon"`
	Layout    string `json:"layout"`
	Component string `json:"component"`
	Show      int    `json:"show"`
	Enable    int    `json:"enable"`
	KeepAlive int    `json:"keepAlive"`
	Order     int    `json:"order"`
}

type AddStationReq struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	AdminUserId uint   `json:"adminUserId" binding:"required"`
	Location    string `json:"location"`
	StationType string `json:"stationType"`
}

type PatchStationReq struct {
	Id          uint   `json:"id" binding:"required"` // 必须要带id，其他都可以不更新
	Code        string `json:"code"`
	Name        string `json:"name"`
	AdminUserId uint   `json:"adminUserId"`
	Location    string `json:"location"`
	StationType string `json:"stationType"`
}

type AddLockReq struct {
	StationId       uint   `json:"stationId" binding:"required"`
	AdminId         uint   `json:"adminId" binding:"required"`
	Name            string `json:"name" comment:"锁名称"`
	Location        string `json:"location"`
	Description     string `json:"description"`
	SN              string `json:"sn" binding:"required"`
	MAC             string `json:"mac" binding:"required"`
	FactoryId       string `json:"factoryId"`
	CurrentKey      string `json:"currentKey"  binding:"required" comment:"当前密钥"`
	FactoryKey      string `json:"factoryKey" binding:"required" comment:"出厂密钥"`
	Enable          bool   `json:"enable" comment:"1: 启用 0: 禁用"`
	SoftwareVersion string `json:"softwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	AlarmMode       string `json:"alarmMode"`
	LockStatus      string `json:"lockStatus"`
	BackupData      string `json:"backupDate" `
	NewLock         string `json:"newLock" `
	UnlockRecord    string `json:"unlockRecord" `
	Power           string `json:"power" `
	Muted           string `json:"muted"`
	Hibernate       string `json:"hibernate"`
}
type PatchLockReq struct {
	Id              uint    `json:"id" binding:"required"`
	StationId       *uint   `json:"stationId,omitempty"`
	AdminId         *uint   `json:"adminId,omitempty"`
	Location        *string `json:"location"`
	Description     *string `json:"description"`
	Name            *string `json:"name"`
	SN              *string `json:"sn,omitempty"`
	MAC             *string `json:"mac,omitempty"`
	FactoryId       *string `json:"factoryId,omitempty"`
	CurrentKey      *string `json:"currentKey,omitempty"`
	FactoryKey      *string `json:"factoryKey,omitempty"`
	Enable          *bool   `json:"enable" comment:"1: 启用 0: 禁用"`
	SoftwareVersion *string `json:"softwareVersion,omitempty"`
	HardwareVersion *string `json:"hardwareVersion,omitempty"`
	AlarmMode       *string `json:"alarmMode,omitempty" comment:"0x00:不报警，0x01:迟钝模式，0x02:中等模式，0x03:敏感模式"`
	LockStatus      *string `json:"lockStatus" `
	BackupData      *string `json:"backupDate,omitempty" `
	NewLock         *string `json:"newLock"`
	UnlockRecord    *string `json:"unlockRecord,omitempty"`
	Power           *string `json:"power"`
	Muted           *string `json:"muted" comment:"0x00静音模式关闭，0x01静音模式打开，其他错误"`
	Hibernate       *string `json:"hibernate" comment:"0x00:正常休眠，0x01:Blue常开，0x02:省电模式"`
}

type LockCommandReq struct {
	Id   uint    `json:"id" binding:"required"`
	Type byte    `json:"type" binding:"required"` // 请求的数据类型
	Roll byte    `json:"roll"`                    // 0x00-0x7f
	Cmd  *string `json:"cmd,omitempty"`           // 02返回01的数据，13返回13的数据
	Key  *string `json:"key,omitempty"`           // newkey 需要变更的加密密钥
}

type AddOrderReq struct {
	StationId      uint                   `json:"stationId" binding:"required"`
	DispatcherId   uint                   `json:"userId" binding:"required"`
	OperatorId     uint                   `json:"operatorId" binding:"required"`
	Status         int                    `json:"status" binding:"required"`
	OrderApprovals []*AddOrderApprovalReq `json:"orderApprovals"`
	OrderSteps     []*AddOrderStepReq     `json:"orderSteps"`
	Content        string                 `json:"content"`
}
type PatchOrderReq struct {
	Id           int     `json:"id" binding:"required"`
	DispatcherId *uint   `json:"dispatcherId"` // 派单人id
	OperatorId   *uint   `json:"operatorId"`   // 操作员id
	StationId    *uint   `json:"stationId"`    // 站点id
	Status       *int    `json:"status"`       // 工单状态
	Content      *string `json:"content"`
}
type AddOrderApprovalReq struct {
	OrderId    uint    `json:"orderId" binding:"required"`
	ApproverId uint    `json:"approverId" binding:"required"`
	Status     int     `json:"status" binding:"required"`
	Sort       int     `json:"sort" binding:"required"`
	Comment    *string `json:"comment"`
}
type PatchOrderApprovalReq struct {
	Id         uint    `json:"id" binding:"required"`
	ApproverId *uint   `json:"approverId" binding:"required"`
	Status     *int    `json:"status"`
	Comment    *string `json:"comment"`
}

type AddOrderStepReq struct {
	OrderId    uint  `json:"orderId" binding:"required"` // 工单id
	Sort       int   `json:"sort" binding:"required"`    // 步骤 1,2,3,...10
	Task       int   `json:"task" binding:"required"`    // 执行任务 1:开锁 2:关锁 3:上传图片 4:状态量
	ReviewerId *uint `json:"reviewerId"`                 // 审核人id  TODO 要不要只保留一个userId，是否需要审核人，还是在工单中显示
	Status     int   `json:"status" binding:"required"`  // 步骤状态 0:待执行 1:已完成 2:待审核 3:重新执行
	// 工单步骤内容
	LockId     *uint   `json:"lockId"`     // 锁id
	LockStatus *int    `json:"lockStatus"` // 锁状态 1:开锁 2:关锁
	ImageUrl   *string `json:"imageUrl"`   // 图片url
	Comment    *string `json:"comment"`    // 状态量
	SwitchId   *string `json:"switchId"`   // 闸刀id
}
type PatchOrderStepReq struct {
	Id         uint  `json:"id" binding:"required"`
	OrderId    uint  `json:"orderId" binding:"required"` // 工单Id，用于确认工单状态
	Sort       *int  `json:"sort"`                       // 步骤 1,2,3,...10
	Task       *int  `json:"task"`                       // 执行任务 1:开锁 2:关锁 3:上传图片 4:状态量
	ReviewerId *uint `json:"reviewerId"`                 // 审核人id  TODO 要不要只保留一个userId，是否需要审核人，还是在工单中显示
	Status     *int  `json:"status"`                     // 步骤状态 0:待执行 1:已完成 2:待审核 3:重新执行
	// 工单步骤内容
	LockId     *uint   `json:"lockId"`     // 锁id
	LockStatus *int    `json:"lockStatus"` // 锁状态 1:开锁 2:关锁
	ImageUrl   *string `json:"imageUrl"`   // 图片url
	Comment    *string `json:"comment"`    // 状态量
	SwitchId   *string `json:"switchId"`   // 闸刀id
}

type BatchDeleteReq struct {
	Ids []uint `json:"ids"`
}
