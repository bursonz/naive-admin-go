package inout

import (
	"naive-admin-go/model"
	"time"
)

// LoginRes 登录请求的响应数据
type LoginRes struct {
	AccessToken string `json:"accessToken"`
}

type UserDetailRes struct {
	model.User
	Profile     *model.Profile `json:"profile"`
	Roles       []*model.Role  `json:"roles" `
	CurrentRole *model.Role    `json:"currentRole"`
}

type RoleListRes []*model.Role

type UserListItem struct {
	ID         int           `json:"id"`
	Username   string        `json:"username"`
	Enable     bool          `json:"enable"`
	CreateTime time.Time     `json:"createTime"`
	UpdateTime time.Time     `json:"updateTime"`
	Gender     *int          `json:"gender"`
	Avatar     string        `json:"avatar"`
	Address    *string       `json:"address"`
	Email      *string       `json:"email"`
	Phone      *string       `json:"phone"`
	Roles      []*model.Role `json:"roles"`
}
type UserListRes struct {
	PageData []UserListItem `json:"pageData"`
	Total    int64          `json:"total"`
}
type RoleListPageItem struct {
	model.Role
	PermissionIds []int64 `json:"permissionIds" gorm:"-"`
}
type RoleListPageRes struct {
	PageData []RoleListPageItem `json:"pageData"`
	Total    int64              `json:"total"`
}

type StationListRes struct {
	PageData []*model.Station `json:"pageData"`
	Total    int64            `json:"total"`
}

type LockListRes struct {
	PageData []*model.Lock `json:"pageData"`
	Total    int64         `json:"total"`
}

type OrderListRes struct {
	PageData []*model.Order `json:"pageData"`
	Total    int64          `json:"total"`
}

type OrderApprovalListRes struct {
	PageData []*model.OrderApproval `json:"pageData"`
	Total    int64                  `json:"total"`
}

type OrderStepListRes struct {
	PageData []*model.OrderStep `json:"pageData"`
	Total    int64              `json:"total"`
}

type LogListRes struct {
	PageData []*model.SysLog `json:"pageData"`
	Total    int64           `json:"total"`
}
