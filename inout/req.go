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
	Id       int     `json:"id"  binding:"required"`
	Enable   *bool   `json:"enable,omitempty"`
	RoleIds  *[]int  `json:"roleIds,omitempty"`
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}
type PatchProfileUserReq struct {
	Id       int     `json:"id"  binding:"required"`
	Gender   *int    `json:"gender"`
	NickName *string `json:"nickName"`
	Address  *string `json:"address"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}
type EnableRoleReq struct {
	Enable bool `json:"enable" binding:"required"`
	Id     int  `json:"id"`
}

type AddUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Enable   bool   `json:"enable" binding:"required"`
	RoleIds  []int  `json:"roleIds" binding:"required"`
}

type AddRoleReq struct {
	Code          string `json:"code" binding:"required"`
	Enable        bool   `json:"enable"`
	Name          string `json:"name" binding:"required"`
	PermissionIds []int  `json:"permissionIds"`
}
type PatchRoleReq struct {
	Id            int     `json:"id"  binding:"required"`
	Code          *string `json:"code,omitempty"`
	Enable        *bool   `json:"enable,omitempty"`
	Name          *string `json:"name,omitempty"`
	PermissionIds *[]int  `json:"permissionIds,omitempty"`
}

type PatchRoleOpeateUserReq struct {
	Id      int   `json:"id" `
	UserIds []int `json:"userIds"`
}

type AddPermissionReq struct {
	Type      string `json:"type" binding:"required"`
	ParentId  *int   `json:"parentId"`
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
	Id        int    `json:"id"  binding:"required"`
	Type      string `json:"type" binding:"required"`
	ParentId  *int   `json:"parentId"`
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
	AdminUserId int    `json:"adminUserId" binding:"required"`
	Location    string `json:"location"`
	StationType string `json:"stationType"`
}

type PatchStationReq struct {
	Id          int    `json:"id" binding:"required"` // 必须要带id，其他都可以不更新
	Code        string `json:"code"`
	Name        string `json:"name"`
	AdminUserId int    `json:"adminUserId"`
	Location    string `json:"location"`
	StationType string `json:"stationType"`
}

type AddLockReq struct {
	StationId       int    `json:"stationId" binding:"required"`
	AdminId         int    `json:"adminId" binding:"required"`
	SN              int    `json:"sn" binding:"required"`
	MAC             string `json:"mac" binding:"required"`
	FactoryId       int    `json:"factoryId" binding:"required"`
	CurrentKey      string `json:"currentKey"  binding:"required" comment:"当前密钥"`
	FactoryKey      string `json:"factoryKey" binding:"required" comment:"出厂密钥"`
	SoftwareVersion string `json:"softwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	Location        string `json:"location"`
	Power           int    `json:"power"`
	Description     string `json:"description"`
	Enable          bool   `json:"enable" binding:"required" comment:"1: 启用 0: 禁用"`
}
type PatchLockReq struct {
	ID              int    `json:"id" binding:"required"`
	StationId       int    `json:"stationId"`
	AdminId         int    `json:"adminId"`
	SN              int    `json:"sn"`
	MAC             string `json:"mac"`
	FactoryId       int    `json:"factoryId"`
	CurrentKey      string `json:"currentKey"`
	FactoryKey      string `json:"factoryKey"`
	SoftwareVersion string `json:"softwareVersion"`
	HardwareVersion string `json:"hardwareVersion"`
	Location        string `json:"location"`
	Power           int    `json:"power"`
	Description     string `json:"description"`
	Enable          bool   `json:"enable"`
}
