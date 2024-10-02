package model

type UserAddress struct {
	UserCode         string `gorm:"index;not null;comment:用户编码"`
	RecipientName    string `gorm:"size:50;not null;comment:收件人姓名"`
	RecipientPhone   string `gorm:"size:15;not null;comment:收件人电话"`
	RecipientAddress string `gorm:"size:200;not null;comment:收件地址"`
	Default          bool   `gorm:"column:is_default;default:false;comment:是否为默认地址"`
	BaseModel
}

func (m *UserAddress) TableName() string {
	return "t_user_address"
}
