package service

import (
	"backend/app/hello/api/grpc"
	"context"
	"reflect"
	"testing"
)

func TestHello_Say(t *testing.T) {
	type args struct {
		ctx context.Context
		req *grpc.SayReq
	}
	tests := []struct {
		name     string
		args     args
		wantResp *grpc.SayResp
		wantErr  bool
	}{
		{
			name:     "test",
			args:     args{context.Background(), &grpc.SayReq{Message: "world"}},
			wantResp: &grpc.SayResp{Message: "Hello world!"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			he := &Hello{}
			gotResp, err := he.Say(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Say() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("Say() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
