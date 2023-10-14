package model

type UserModel struct {
	Id       int64  `json:"id,omitempty"`
	RuleId   int    `json:"rule_id"`
	Email    string `json:"email"`
	Ip       string `json:"ip"`
	Password string `json:"-"`
	Username string `json:"username,omitempty"`
	Time
}

type Time struct {
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type NewPermissionModel struct {
	Id         int64          `json:"id,omitempty"`
	Key        string         `json:"key,omitempty"`
	Value      string         `json:"value,omitempty"`
	Type       PermissionType `json:"type,omitempty"` //0 is button 1 is menu 2 is page 3 api
	Desc       string         `json:"desc,omitempty"`
	Status     int            `json:"status,omitempty"`
	OtherValue string         `json:"other_value,omitempty" gorm:"column:value1"`
	Time
}
type PermissionType int

const (
	Button PermissionType = iota
	Menu
	Page
	Api
)

type RuleModel struct {
	Id           int    `json:"id,omitempty"`
	PermissionId int    `json:"permission_id,omitempty"`
	SuperId      int    `json:"super_id,omitempty"` // 父类id
	Status       int    `json:"status,omitempty"`
	Name         string `json:"name,omitempty"`
	Desc         string `json:"desc,omitempty"`
	HasChildren  bool   `json:"has_children" gorm:"-"`
	Time
}
type RuleUserModel struct {
	Id         int64 `json:"id,omitempty"`
	UserId     int64 `json:"user_id,omitempty"`
	MerchantId int64 `json:"merchant_id,omitempty"`
	RuleId     int64 `json:"rule_id,omitempty"`
	Time
}
