package v1

type SendVerifyCodeRequest struct {
	Phone string `json:"phone" binding:"required" example:"13288888888"`
}

type LoginRequest struct {
	Phone      string `json:"phone" example:"123456"`
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
