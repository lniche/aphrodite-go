package v1

type SendVerifyCodeRequest struct {
	Phone string `json:"phone" binding:"required" example:"13288888888"`
}

type RegisterRequest struct {
	Nickname   string `json:"nickname" example:"banana"`
	Email      string `json:"email" binding:"email" example:"1234@gmail.com"`
	Password   string `json:"password" binding:"required" example:"123456"`
	Phone      string `json:"phone" binding:"required" example:"123456"`
	VerifyCode string `json:"verifyCode"  example:"1234"`
	ClientIp   string `json:"clientIp" example:"127.0.0.1"`
}

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required" example:"13288888888"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

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
