package common

import "context"

type RequestResponse struct {
	sub string
	tid string
}

func NewRequestResponse(sub string, tid string) *RequestResponse {
	return &RequestResponse{sub: sub, tid: tid}
}

func (r *RequestResponse) GetSub() string {
	return r.sub
}

func (r *RequestResponse) GetTid() string {
	return r.tid
}

type Request interface {
	GetSub() string
	GetTid() string
}

type keyRequest string

var KeyReq keyRequest

func SaveRequestContext(cxt context.Context, r Request) context.Context {
	return context.WithValue(cxt, KeyReq, r)
}

func GetRequestContext(cxt context.Context) Request {
	r, _ := cxt.Value(KeyReq).(Request)
	return r
}
