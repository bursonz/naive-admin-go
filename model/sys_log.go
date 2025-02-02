package model

// TODO 补全gorm tag
type SysLog struct {
	BaseModel
	UserId   uint   `json:"userId"`   // 用户id
	UserName string `json:"userName"` // 用户名称
	Method   string `json:"method"`   // 操作 添加，修改，删除，查询
	Target   string `json:"target"`   // 目标对象 用户，工单，工单步骤，工单审批...
	Content  string `json:"content"`  // 内容
	IP       string `json:"ip"`       // ip
	Path     string `json:"path"`     // 路径
}

func (SysLog) TableName() string { return "sys_log" }
