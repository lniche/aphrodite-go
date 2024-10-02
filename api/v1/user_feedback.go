package v1

type AddUserFeedbackRequest struct {
	Feedback string `json:"feedback" binding:"required" example:"banana"`
}
