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
<<<<<<< HEAD
=======

package common

import "context"

type requestResponse struct {
	sub string
	tid string
}

func NewRequestResponse(sub string, tid string) *requestResponse {
	return &requestResponse{sub: sub, tid: tid}
}

func (r *requestResponse) GetSub() string {
	return r.sub
}

func (r *requestResponse) GetTid() string {
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
>>>>>>> 70a38361bb67beb662f248595a90edb388469f20
