package exceptions

import "fmt"

type CommonExceptionInterface interface {
	GetHTTPErrorCode() int
	GetErrorMessage() string
	Error() string
}
type CommonException struct {
	errorCode    int
	errorMessage string
}

func (ex CommonException) Error() string {
	return fmt.Sprintf("Error %d: %s", ex.errorCode, ex.errorMessage)
}

func NewCommonException(errorCode int, errorMessage string) CommonExceptionInterface {
	return CommonException{
		errorCode:    errorCode,
		errorMessage: errorMessage,
	}
}

func (ex CommonException) GetHTTPErrorCode() int {
	return ex.errorCode
}

func (ex CommonException) GetErrorMessage() string {
	return ex.errorMessage
}
func NewBadRequestException(errorMessage string) CommonException {
	return CommonException{
		errorCode:    400,
		errorMessage: errorMessage,
	}
}
func NewNotHavePermissionException() CommonException {
	return CommonException{
		errorCode:    403,
		errorMessage: "You do not have permission",
	}
}
func NewTargetNotExistException(target string) CommonException {
	return CommonException{
		errorCode:    404,
		errorMessage: target + " not exist",
	}
}
func NewTargetAlreadyExistException(target string) CommonException {
	return CommonException{
		errorCode:    409,
		errorMessage: target + " already exist",
	}
}
func NewInternalErrorException(errorMessage string) CommonException {
	return CommonException{
		errorCode:    500,
		errorMessage: errorMessage,
	}
}
