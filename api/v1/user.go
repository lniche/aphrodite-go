package v1

// UpdateUserReq User update request structure
// @Description Request data for updating user information
// @Param nickname body string false "User nickname" example("john")
// @Param email body string false "User email" example("john@example.com")
type UpdateUserReq struct {
	Nickname string `json:"nickname" example:"john"`
	Email    string `json:"email" binding:"email" example:"john@example.com"`
}

// GetUserRespData User information response structure
// @Description Response data for user information
// @Property userNo string "User number"
// @Property userCode string "User code"
// @Property nickname string "User nickname" example("john")
// @Property email string "User email" example("john@example.com")
// @Property phone string "User phone number" example("13800138000")
type GetUserRespData struct {
	UserNo   string `json:"userNo"`
	UserCode string `json:"userCode"`
	Nickname string `json:"nickname" example:"john"`
	Email    string `json:"email" binding:"email" example:"john@example.com"`
	Phone    string `json:"phone" binding:"required" example:"13800138000"`
}

// GetUserResp Complete structure for user information response
// @Description Complete response for user information
// @Property code int "Response status code"
// @Property message string "Response message"
// @Property data GetUserRespData "User information data"
type GetUserResp struct {
	Response
	Data GetUserRespData `json:"data"` // User information data
}
