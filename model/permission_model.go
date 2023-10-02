package model

type PermissionModel struct {
	Type   string   `json:"type,omitempty" gorm:"column:p_type"`
	User   string   `json:"user,omitempty" gorm:"column:v0"`
	Source string   `json:"source,omitempty" gorm:"column:v1"`
	Value  string   `json:"value,omitempty" gorm:"column:v2"`
	Values []string `json:"values,omitempty" gorm:"-"`
}
