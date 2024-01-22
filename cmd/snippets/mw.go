package main

import (
	"net/http"

	"github.com/sjxiang/example/pkg/apperr"
)


func customHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), apperr.HTTPStatus(err))
		}
	}
}


// gin => ctx.JSON(apperr.HTTPStatus(err), err.Error())