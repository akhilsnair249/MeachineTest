package main

import (
	"errors"
	"fmt"
)

type ErrorFrom struct {
	Subsystem, Operation string
}

type DBError struct {
	ActualError error
	Messsage    string
	ErrorFrom   ErrorFrom
}

func DBErrorNew(message string, actualErr error, errorFrom ErrorFrom) DBError {

	errObj := DBError{Messsage: message, ActualError: actualErr, ErrorFrom: errorFrom}

	return errObj
}

func (dbErr DBError) Error() string {

	status := dbErr.Status()
	oErr := errors.New("nil")

	if dbErr.ActualError != nil {
		oErr = dbErr.ActualError
	}
	return fmt.Sprintf("Database error occured: %d - %s - %s | Actual error: %s", status, dbErr.StackTrace(), dbErr.Message(), oErr.Error())
}

func (dbErr DBError) Status() int {
	return 500
}

func (dbErr DBError) Message() string {
	return dbErr.Messsage
}

func (dbErr DBError) StackTrace() string {
	return dbErr.ErrorFrom.Subsystem + " -> " + dbErr.ErrorFrom.Operation
}
