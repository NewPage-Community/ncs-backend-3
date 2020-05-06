package ecode

import (
	"google.golang.org/grpc/codes"
	"testing"
)

func TestError(t *testing.T) {
	err := Errorf(codes.Internal, "OK!")
	if GetError(err).Code != codes.Internal {
		t.Errorf("Error code: v = %d, want = 13", GetError(err).Code)
	}
	if GetError(nil) != nil {
		t.Errorf("GetError(): v = %v, want = nil", GetError(err))
	}
}
