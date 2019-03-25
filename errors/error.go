package errors

import "fmt"

const (
	ENV_VAR_ERR        = Error("Environment Variable NOT Set")
	MISSING_CONTACT_ID = Error("Missing Contact id in request")
)

type Error string

func (e Error) Error() string { return string(e) }

func (e Error) New(s string) error {
	return Error(s)
}

type DetailedError struct {
	Err         error
	Description string
}

func (de DetailedError) Error() string {
	return fmt.Sprintf("Description: %s | error: %s", de.Description, de.Err.Error())
}
