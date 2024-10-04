package v1

type UpdateProfileRequest struct {
	Nickname    string `json:"nickname" example:"banana"`
	Email       string `json:"email" binding:"email" example:"1234@gmail.com"`
	OldPassword string `json:"oldPassword" example:"banana"`
	NewPassword string `json:"newPassword" example:"banana"`
	OldPhone    string `json:"oldPhone"  example:"13288888888"`
	NewPhone    string `json:"newPhone"  example:"13288888888"`
	VerifyCode  string `json:"verifyCode"  example:"1234"`
}
type GetProfileResponseData struct {
	UserNo   string `json:"userNo"`
	UserCode string `json:"userCode"`
	Nickname string `json:"nickname" example:"banana"`
	Email    string `json:"email" binding:"email" example:"1234@gmail.com"`
	Phone    string `json:"phone" binding:"required" example:"13288888888"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}
