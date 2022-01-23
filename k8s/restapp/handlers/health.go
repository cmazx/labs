package handlers

import (
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/restapi/operations/health"
	"github.com/go-openapi/runtime/middleware"
)

type hlf struct {
}

func NewHealthHandler() *hlf {
	return &hlf{}
}

func (impl *hlf) Handle(params health.HealthParams) middleware.Responder {
	return health.NewHealthOK()
}
