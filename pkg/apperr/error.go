package apperr

import (
	"google.golang.org/grpc/codes"
)


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


var status = map[int]codes.Code {
	0: codes.OK,
	1: codes.Canceled,
	2: codes.Unknown,
	3: codes.InvalidArgument,
	4: codes.DeadlineExceeded,
	5: codes.NotFound,
	6: codes.AlreadyExists,
	7: codes.PermissionDenied,
	8: codes.ResourceExhausted,
	9: codes.FailedPrecondition,
	10: codes.Aborted,
	11: codes.OutOfRange,
	12: codes.Unimplemented,
	13: codes.Internal,
	14: codes.Unavailable,
	15: codes.DataLoss,
	16: codes.Unauthenticated,
}