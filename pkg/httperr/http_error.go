package httperr

type RestErr struct {
	Msg  string `json:"msg"`
	Err  string `json:"error,omitempty"`
	Code int    `json:"code"`
	// Fields 
}

type Fields struct {
	Field string 
	Value any
	
}