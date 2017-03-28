package main

import (
	"fmt"
)

//START OMIT
type CustomError struct {
	Status  int
	Message string
	Type    string
}

func InternalServerError(message string) CustomError {
	return newError(500, "internal_server_error", message)
}
func (e CustomError) Error() string {
	return fmt.Sprintf("An error has occured during processing! Details %s ", e.Message)
}

func newError(status int, errType string, message string) CustomError {
	ce := CustomError{status, errType, message}
	return ce
}

func process() error {
	//assume some processing is done here
	return InternalServerError("Error while calling upstream service")
}
//END OMIT
func main() {
	err := process()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

