package handlerfunc

import (
	"net/http"

	"github.com/dacut/aws-lambda-go-api-proxy/httpadapter"
)

type HandlerFuncAdapterALB = httpadapter.HandlerAdapterALB

func NewALB(handlerFunc http.HandlerFunc) *HandlerFuncAdapterALB {
	return httpadapter.NewALB(handlerFunc)
}
