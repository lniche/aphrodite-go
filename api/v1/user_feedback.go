package v1

type CreateUserFeedbackRequest struct {
	Feedback string `json:"feedback" binding:"required" example:"banana"`
}
