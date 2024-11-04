package v1

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrMethodNotAllow      = newError(404, "Method Not Allowed")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrData    = newError(1001, "Data Error")
	ErrService = newError(1002, "Service Error")
)
