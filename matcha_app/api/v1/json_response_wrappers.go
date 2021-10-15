package api

import "github.com/savsgio/atreugo/v11"

type HttpStatus int
type AnyJsonObj interface{}
type JsonResponseErrHandler func(*atreugo.RequestCtx) (error, HttpStatus, AnyJsonObj)
type JsonResponseHandler func(*atreugo.RequestCtx) (HttpStatus, AnyJsonObj)

func WrapJsonResponseWithErr(handler JsonResponseErrHandler) atreugo.View {
	return func(r *atreugo.RequestCtx) error {
		if err, status_code, json_obj := handler(r); err != nil {
			return err
		} else {
			return r.JSONResponse(json_obj, int(status_code))
		}
	}
}

func WrapJsonResponse(handler JsonResponseHandler) atreugo.View {
	return func(r *atreugo.RequestCtx) error {
		status_code, json_obj := handler(r)
		return r.JSONResponse(json_obj, int(status_code))
	}
}
