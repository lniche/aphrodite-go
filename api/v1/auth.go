package v1

// SendVerifyCodeReq Request structure for sending verification code
// @Description Request data for sending verification code
// @Param phone body string true "User phone number" example("13800138000")
type SendVerifyCodeReq struct {
	Phone string `json:"phone" binding:"required" example:"13800138000"`
}

// LoginReq Login request structure
// @Description User login request data
// @Param phone body string true "User phone number" example("13800138000")
// @Param verifyCode body string false "Verification code" example("1234")
type LoginReq struct {
	Phone      string `json:"phone" binding:"required" example:"13800138000"`
	VerifyCode string `json:"verifyCode" example:"1234"`
}

// LoginRespData Structure for login response data
// @Description Login response data
// @Property accessToken string "Access token"
type LoginRespData struct {
	AccessToken string `json:"accessToken"`
}

// LoginResp Complete structure for login response
// @Description Complete login response
// @Property code int "Response status code"
// @Property message string "Response message"
// @Property data LoginRespData "Login response data"
type LoginResp struct {
	Response
	Data LoginRespData `json:"data"` // Login response data
}
