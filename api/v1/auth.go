package v1

type SendVerifyCodeRequest struct {
	Phone string `json:"phone" binding:"required" example:"13800138000"`
}

type LoginRequest struct {
	Phone      string `json:"phone" binding:"required" example:"13800138000"`
	VerifyCode string `json:"verifyCode" example:"1234"`
	OpenId     string `json:"openId" example:"123456"`
}

type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}
