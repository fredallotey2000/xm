package validator

type Validator interface {
	ValidateJSON(data interface{}) error
}
