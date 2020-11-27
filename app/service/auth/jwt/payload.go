package jwt

import (
	"backend/pkg/ecode"
	"google.golang.org/grpc/codes"
)

type Payload map[string]interface{}

func (p Payload) Get(key string) (res interface{}, err error) {
	res, ok := p[key]
	if !ok {
		err = ecode.Errorf(codes.NotFound, "missing key: %s", key)
	}
	return
}

func (p *Payload) SetUID(uid int64) {
	(*p)["uid"] = uid
}

func (p Payload) GetUID() (uid int64, err error) {
	res, err := p.Get("uid")
	if err != nil {
		return
	}
	uid, ok := res.(int64)
	if !ok {
		err = ecode.Errorf(codes.NotFound, "invalid uid of type")
	}
	return
}
