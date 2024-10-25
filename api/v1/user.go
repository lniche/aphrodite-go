package v1

type UpdateUserRequest struct {
	Nickname    string `json:"nickname" example:"banana"`
	Email       string `json:"email" binding:"email" example:"1234@gmail.com"`
	OldPassword string `json:"oldPassword" example:"banana"`
	NewPassword string `json:"newPassword" example:"banana"`
	OldPhone    string `json:"oldPhone"  example:"13800138000"`
	NewPhone    string `json:"newPhone"  example:"13800138000"`
	VerifyCode  string `json:"verifyCode"  example:"1234"`
}
type GetUserResponseData struct {
	UserNo   string `json:"userNo"`
	UserCode string `json:"userCode"`
	Nickname string `json:"nickname" example:"banana"`
	Email    string `json:"email" binding:"email" example:"1234@gmail.com"`
	Phone    string `json:"phone" binding:"required" example:"13800138000"`
}
type GetUserResponse struct {
	Response
	Data GetUserResponseData
}
