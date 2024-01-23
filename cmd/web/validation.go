package main

import (
  "fmt"
  "reflect"
  "strings"

  "github.com/go-playground/validator/v10"
 )


type invalidArgument struct {
	Field   string `json:"field"`              // 字段名
	Value   any    `json:"value,omitempty"`    // 参数传递值
	Message string `json:"message"`            // 提示
}

func ValidateHttpData(d interface{}) []invalidArgument {
	val := validator.New()

	// extract json tag name
	val.RegisterTagNameFunc(
		func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		},
	)

	if err := val.Struct(d); err != nil {
		var errorsCauses []invalidArgument

		for _, e := range err.(validator.ValidationErrors) {
			cause := invalidArgument{}
			fieldName := e.Field()

			switch e.Tag() {
			case "required":
				cause.Message = fmt.Sprintf("%s is required", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "uuid4":
				cause.Message = fmt.Sprintf("%s is not a valid uuid", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "boolean":
				cause.Message = fmt.Sprintf("%s is not a valid boolean", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "min":
				cause.Message = fmt.Sprintf("%s must be greater than %s", fieldName, e.Param())
				cause.Field = fieldName
				cause.Value = e.Value()
			case "max":
				cause.Message = fmt.Sprintf("%s must be less than %s", fieldName, e.Param())
				cause.Field = fieldName
				cause.Value = e.Value()
			case "email":
				cause.Message = fmt.Sprintf("%s is not a valid email", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			case "containsany":
				cause.Message = fmt.Sprintf("%s must contain at least one of the following characters: !@#$%%*", fieldName)
				cause.Field = fieldName
				cause.Value = e.Value()
			default:
				cause.Message = "invalid field"
				cause.Field = fieldName
				cause.Value = e.Value()
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return errorsCauses
	}

	return []invalidArgument{}
}

/*
curl --location 'http://localhost:9000/register' \
--header 'Content-Type: application/json' \
--data '{
    "name":"fo",
    "password":"12",
	"email": "1224@qq.com",
	"verify_code": "123456"
}'


	Name       string `json:"name"        validate:"required,min=3,max=20"`
	Password   string `json:"password"    validate:"required,min=3,max=20,containsany=!@#$%*"`
	Email      string `json:"email"       validate:"required,email"`
	VerifyCode string `json:"verify_code" validate:"required,min=6,max=6"`
	Sex        int     `json:"sex"        validate:"omitempty"`
*/
