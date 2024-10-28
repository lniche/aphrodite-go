package v1

// UpdateUserReq 用户更新请求结构体
// @Description 用户更新信息请求
// @Param nickname body string false "用户昵称" example("banana")
// @Param email body string false "用户邮箱" example("1234@gmail.com")
// @Param oldPassword body string false "旧密码" example("old_password")
// @Param newPassword body string false "新密码" example("new_password")
// @Param oldPhone body string false "旧手机号" example("13800138000")
// @Param newPhone body string false "新手机号" example("13800138000")
// @Param verifyCode body string false "验证码" example("1234")
type UpdateUserReq struct {
	Nickname    string `json:"nickname" example:"banana"`
	Email       string `json:"email" binding:"email" example:"1234@gmail.com"`
	OldPassword string `json:"oldPassword" example:"banana"`
	NewPassword string `json:"newPassword" example:"banana"`
	OldPhone    string `json:"oldPhone"  example:"13800138000"`
	NewPhone    string `json:"newPhone"  example:"13800138000"`
	VerifyCode  string `json:"verifyCode"  example:"1234"`
}

// GetUserRespData 用户信息响应结构体
// @Description 用户信息响应数据
// @Property userNo string "用户编号"
// @Property userCode string "用户代码"
// @Property nickname string "用户昵称" example("banana")
// @Property email string "用户邮箱" example("1234@gmail.com")
// @Property phone string "用户手机号" example("13800138000")
type GetUserRespData struct {
	UserNo   string `json:"userNo"`
	UserCode string `json:"userCode"`
	Nickname string `json:"nickname" example:"banana"`
	Email    string `json:"email" binding:"email" example:"1234@gmail.com"`
	Phone    string `json:"phone" binding:"required" example:"13800138000"`
}

// GetUserResp 用户信息完整响应结构体
// @Description 用户信息完整响应
// @Property code int "响应状态码"
// @Property message string "响应消息"
// @Property data GetUserResponseData "用户信息数据"
type GetUserResp struct {
	Response
	Data GetUserRespData `json:"data"` // 用户信息数据
}
