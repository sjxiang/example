package apperr

import "google.golang.org/grpc/codes"


// 鸡肋，友好提示不能搞，沾点其它也不行
type statusError struct {
	// 嵌入 error，类型提升
	error 
	status int
	code   codes.Code
}

func (e statusError) Unwrap() error  {
	return e.error
}

func (e statusError) HTTPStatus() int {
	return e.status
}

func (e statusError) GRPCStatus() codes.Code {
	return e.code
}

func WithHTTPStatus(err error, status int) error {
	return statusError{
		error:  err,
		status: status,
	}
}

func WithGRPCStatus(err error, code codes.Code) error {
	return statusError{
		error: err,
		code:  code,
	}
}


// `"OK"`: OK, 0
// `"CANCELLED"`:/* [sic] */ Canceled,
// `"UNKNOWN"`:             Unknown,
// `"INVALID_ARGUMENT"`:    InvalidArgument,
// `"DEADLINE_EXCEEDED"`:   DeadlineExceeded,
// `"NOT_FOUND"`:           NotFound,
// `"ALREADY_EXISTS"`:      AlreadyExists,
// `"PERMISSION_DENIED"`:   PermissionDenied,
// `"RESOURCE_EXHAUSTED"`:  ResourceExhausted,
// `"FAILED_PRECONDITION"`: FailedPrecondition,
// `"ABORTED"`:             Aborted,
// `"OUT_OF_RANGE"`:        OutOfRange,
// `"UNIMPLEMENTED"`:       Unimplemented,
// `"INTERNAL"`:            Internal,
// `"UNAVAILABLE"`:         Unavailable,
// `"DATA_LOSS"`:           DataLoss,
// `"UNAUTHENTICATED"`:     Unauthenticated, 16