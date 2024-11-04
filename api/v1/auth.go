package v1

// SendVerifyCodeReq 发送验证码请求结构体
// @Description 发送验证码请求数据
// @Param phone body string true "用户手机号" example("13800138000")
type SendVerifyCodeReq struct {
	Phone string `json:"phone" binding:"required" example:"13800138000"`
}

// LoginReq 登录请求结构体
// @Description 用户登录请求数据
// @Param phone body string true "用户手机号" example("13800138000")
// @Param verifyCode body string false "验证码" example("1234")
type LoginReq struct {
	Phone      string `json:"phone" binding:"required" example:"13800138000"`
	VerifyCode string `json:"verifyCode" example:"1234"`
}

// LoginRespData 登录响应数据结构体
// @Description 登录响应数据
// @Property accessToken string "访问令牌"
type LoginRespData struct {
	AccessToken string `json:"accessToken"`
}

// LoginResp 登录完整响应结构体
// @Description 登录完整响应
// @Property code int "响应状态码"
// @Property message string "响应消息"
// @Property data LoginRespData "登录响应数据"
type LoginResp struct {
	Response
	Data LoginRespData `json:"data"` // 登录响应数据
}
