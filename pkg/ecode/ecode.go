package ecode

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Code    codes.Code
	Message string
}

//Errorf returns an error representing c and msg. If c is OK, returns nil.
func Errorf(c codes.Code, format string, a ...interface{}) error {
	return status.Errorf(c, format, a...)
}

func GetError(err error) *Error {
	if err != nil {
		errStatus, ok := status.FromError(err)
		if ok && errStatus != nil {
			return &Error{
				Code:    errStatus.Code(),
				Message: errStatus.Message(),
			}
		}
	}
	return nil
}
