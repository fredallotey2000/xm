package validator

type Validator interface {
	ValidateJSON(data interface{}) (bool, error)
}
