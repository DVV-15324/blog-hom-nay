package common

import "context"

<<<<<<< HEAD
type requestResponse struct {
=======
type RequestResponse struct {
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	sub string
	tid string
}

<<<<<<< HEAD
func NewRequestResponse(sub string, tid string) *requestResponse {
	return &requestResponse{sub: sub, tid: tid}
}

func (r *requestResponse) GetSub() string {
	return r.sub
}

func (r *requestResponse) GetTid() string {
=======
func NewRequestResponse(sub string, tid string) *RequestResponse {
	return &RequestResponse{sub: sub, tid: tid}
}

func (r *RequestResponse) GetSub() string {
	return r.sub
}

func (r *RequestResponse) GetTid() string {
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
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
