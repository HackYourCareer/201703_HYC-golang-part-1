package main

import (
	"fmt"
)

//START OMIT
type CustomError struct {
	Status int
	Message string
	Type string
}
func (e CustomError) InternalServerError(message string) CustomError {
	return e.newError(500, "internal_server_error", message)
}
func (e CustomError) Error()string{
	return fmt.Sprintf("An error has occured during processing! Details %s ", e.Message)
}

func (e CustomError) newError(status int, errType string, message string) CustomError {
	return CustomError{
		Status: status,
		Type: errType,
		Message: message,
	}
}
//END OMIT
//SKIP OMIT
func process()error {
	//assume some processing is done here
	return CustomError{}.InternalServerError("Error while calling upstream service")
}

func main(){
	err := process()
	if err != nil {
		fmt.Printf(err.Error())
	}
}
//SKIP OMIT

