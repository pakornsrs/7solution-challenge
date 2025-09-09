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
	UserAlreeadyExistError     = CustomError{Code: "Error-0004", Message: "This e-main has been registed"}
	BadRequestError            = CustomError{Code: "Error-0005", Message: "Bad request"}
	ValidateRequestError       = CustomError{Code: "Error-0006", Message: "Validate request error"}
	UserNotFoundError          = CustomError{Code: "Error-0007", Message: "Not found user"}
	UnauthorizedRequestError   = CustomError{Code: "Error-0008", Message: "Unauthorized request"}

	UnknownError = CustomError{Code: "Error-9999", Message: "Unknow error"}
)
