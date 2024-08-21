package enum

type UTMConfigDataType string

const (
	TextType     UTMConfigDataType = "text"
	TelType      UTMConfigDataType = "tel"
	PasswordType UTMConfigDataType = "password"
	EmailType    UTMConfigDataType = "email"
	NumberType   UTMConfigDataType = "number"
	BooleanType  UTMConfigDataType = "bool"
	ListType     UTMConfigDataType = "list"
)
