package service

import (
	pb "backend/app/user/account/api/grpc"
	"backend/app/user/account/dao"
	"context"
	"reflect"
	"testing"
)

func TestService_ChangeName(t *testing.T) {
	// TODO: Add test cases.
}

func TestService_Info(t *testing.T) {
	type fields struct {
		dao dao.Dao
	}
	type args struct {
		ctx context.Context
		req *pb.InfoReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *pb.InfoResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				dao: tt.fields.dao,
			}
			gotRes, err := s.Info(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Info() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Info() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestService_Register(t *testing.T) {
	type fields struct {
		dao dao.Dao
	}
	type args struct {
		ctx context.Context
		req *pb.RegisterReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *pb.RegisterResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				dao: tt.fields.dao,
			}
			gotRes, err := s.Register(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Register() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestService_UID(t *testing.T) {
	type fields struct {
		dao dao.Dao
	}
	type args struct {
		ctx context.Context
		req *pb.UIDReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes *pb.UIDResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				dao: tt.fields.dao,
			}
			gotRes, err := s.UID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("UID() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
