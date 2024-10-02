package model

type UserFeedback struct {
	UserCode string `gorm:"index;not null;comment:用户编码"`
	Feedback string `gorm:"not null;type:text;comment:反馈"`
	BaseModel
}

func (m *UserFeedback) TableName() string {
	return "t_user_feedback"
}
