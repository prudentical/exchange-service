package api

type ValidationError struct {
	msg string
}

func (v ValidationError) Error() string {
	return v.msg
}

type QueryParamRequiredError struct {
	msg string
}

func (v QueryParamRequiredError) Error() string {
	return v.msg
}

type InvalidIDError struct {
}

func (e InvalidIDError) Error() string {
	return "Invalid id"
}
