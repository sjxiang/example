package apperr

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

func HTTPStatus(err error) int {
	if err == nil {
		return 0
	}

	var statusErr interface {
		error 
		HTTPStatus() int
	}

	if errors.As(err, &statusErr) {
		return statusErr.HTTPStatus()
	}

	return http.StatusInternalServerError
}


func GRPCStatus(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	
	var statusErr interface {
		error 
		GRPCStatus() codes.Code
	}

	if errors.As(err, &statusErr) {
		return statusErr.GRPCStatus()
	}

	return codes.Internal
}
