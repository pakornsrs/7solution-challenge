package constants

// status code
var (
	SuccessCode    = "0000"
	SuccessMessage = "success"

	BadRequestErrorCode    = "400"
	BadRequestErrorMessage = "bad request"

	UnauthorizedErrorCode    = "401"
	UnauthorizedErrorMessage = "Unauthorized."

	InternalProcessFailedErrorCode    = "9999"
	InternalProcessFailedErrorMessage = "internal service error"

	NoData = "no data"
)

// error constant
type CustomError struct {
	Code    string
	Message string
}

var (
	UserIdFormatIncorrectError = CustomError{Code: "Error-0001", Message: "User ID format is incorrect"}
	MongoDbError               = CustomError{Code: "Error-0002", Message: "MongoDb error"}
	AuthenticationError        = CustomError{Code: "Error-0003", Message: "Not found user or password is incorrect"}

	UnknownError = CustomError{Code: "Error-9999", Message: "Unknow error"}
)
