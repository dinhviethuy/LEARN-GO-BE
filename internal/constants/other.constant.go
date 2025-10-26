package constants

type ValidateType string

const (
	ValidateBody  ValidateType = "body"
	ValidateQuery ValidateType = "query"
	ValidateParam ValidateType = "param"
)
